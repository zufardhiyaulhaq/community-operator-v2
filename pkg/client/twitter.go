package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dghubble/oauth1"
	"github.com/drswork/go-twitter/twitter"
)

type TwitterClient struct {
	client *twitter.Client
}

func (t TwitterClient) Send(c Chattable) (Response, error) {
	value := c.Value()

	if c.Method() == "sendText" {
		tweet, _, err := t.client.Statuses.Update(value["text"], nil)
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: tweet.IDStr,
		}, nil
	}

	if c.Method() == "sendTextAndImageUrl" || c.Method() == "sendImageUrl" {
		response, err := http.Get(value["imageUrl"])
		if err != nil {
			return Response{}, err
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return Response{}, err
		}

		media, _, err := t.client.Media.Upload(contents, "tweet_image")
		if err != nil {
			return Response{}, err
		}

		tweet, _, err := t.client.Statuses.Update(value["text"], &twitter.StatusUpdateParams{
			MediaIds: []int64{media.MediaID},
		})
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: fmt.Sprint(tweet.ID),
		}, nil
	}

	return Response{}, fmt.Errorf("chattable method not found")
}

func (t TwitterClient) Reply(c Chattable, identifier string) (Response, error) {
	value := c.Value()

	identifiterInt64, err := strconv.ParseInt(identifier, 10, 64)
	if err != nil {
		return Response{}, err
	}

	if c.Method() == "sendText" {
		tweet, _, err := t.client.Statuses.Update(value["text"], &twitter.StatusUpdateParams{
			InReplyToStatusID: identifiterInt64,
		})
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: tweet.IDStr,
		}, nil
	}

	if c.Method() == "sendTextAndImageUrl" || c.Method() == "sendImageUrl" {
		response, err := http.Get(value["imageUrl"])
		if err != nil {
			return Response{}, err
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return Response{}, err
		}

		media, _, err := t.client.Media.Upload(contents, "tweet_image")
		if err != nil {
			return Response{}, err
		}

		tweet, _, err := t.client.Statuses.Update(value["text"], &twitter.StatusUpdateParams{
			MediaIds:          []int64{media.MediaID},
			InReplyToStatusID: identifiterInt64,
		})
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: tweet.IDStr,
		}, nil
	}

	return Response{}, fmt.Errorf("chattable method not found")
}

func NewTwitterClient(apiKey, apiKeySecret, accessToken, accessTokenSecret string) (Client, error) {
	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	_, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	})

	if err != nil {
		return TwitterClient{}, err
	}

	return TwitterClient{
		client: client,
	}, nil
}
