package client

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterClient struct {
	client *twitter.Client
}

func (client TwitterClient) Send(c Chattable) (Response, error) {
	return Response{}, fmt.Errorf("chattable method not found")
}

func (client TwitterClient) Reply(c Chattable, identifier string) (Response, error) {
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
