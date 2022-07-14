package handler

import (
	"strings"

	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/client"
	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/message"
)

type TwitterHandler struct {
	client      client.Client
	handlerType message.HandlerType
}

func (handler TwitterHandler) SendMessage(message message.Message) error {
	renderedText, err := message.RenderText(handler.handlerType)
	if err != nil {
		return err
	}

	texts := splitText(renderedText.String())
	var response client.Response

	for index, text := range texts {
		if index == 1 {
			var chattable client.Chattable

			if message.SendImage() {
				imageUrl, err := message.RenderImageUrl()
				if err != nil {
					return err
				}

				chattable = client.TextAndImageUrl{
					Text:     text,
					ImageUrl: imageUrl,
				}
			} else {
				chattable = client.Text{
					Text: text,
				}
			}

			response, err = handler.client.Send(chattable)
			if err != nil {
				return err
			}

			continue
		}

		response, err = handler.client.Reply(client.Text{Text: text}, response.Identifier)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewTwitterHandler(client client.Client) Handler {
	return TwitterHandler{
		client:      client,
		handlerType: message.Twitter,
	}
}

const TWITTER_CHARACTER_LIMIT = 250

func splitText(text string) []string {
	texts := strings.Split(text, "\n")

	var filteredTexts []string
	for _, str := range texts {
		if str != "" {
			filteredTexts = append(filteredTexts, str)
		}
	}

	var mergedTexts []string
	for i := 0; i < len(filteredTexts); i++ {
		if i == len(filteredTexts)-1 {
			mergedTexts = append(mergedTexts, filteredTexts[i])
			continue
		}

		if len(filteredTexts[i]) > TWITTER_CHARACTER_LIMIT {
			mergedTexts = append(mergedTexts, filteredTexts[i][:TWITTER_CHARACTER_LIMIT])
			mergedTexts = append(mergedTexts, filteredTexts[i][TWITTER_CHARACTER_LIMIT:len(filteredTexts[i])])
			continue
		}

		mergedText := filteredTexts[i]
		var count int

		for j := i; i < len(filteredTexts); j++ {
			if j == len(filteredTexts)-1 {
				break
			}

			if len(mergedText)+len(filteredTexts[j+1]) < TWITTER_CHARACTER_LIMIT {
				mergedText = strings.Join([]string{mergedText, filteredTexts[j+1]}, "\n\n")
				count++
				continue
			} else {
				break
			}
		}

		mergedTexts = append(mergedTexts, mergedText)
		i += count
	}

	return mergedTexts
}
