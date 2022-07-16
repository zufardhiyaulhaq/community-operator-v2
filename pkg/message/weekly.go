package message

import (
	"bytes"
	"html/template"
	"net/url"
)

type Weekly_Article struct {
	Title string
	Url   string
	Type  string
}

type Weekly struct {
	Name     string
	Date     string
	ImageUrl string
	Tags     []string
	Articles []Weekly_Article
}

func (weekly Weekly) RenderText(handler HandlerType) (bytes.Buffer, error) {
	var templateFile string
	var output bytes.Buffer

	switch handler {
	case Telegram:
		templateFile = WEEKLY_TELEGRAM_TEMPLATE
	case Twitter:
		templateFile = WEEKLY_TWITTER_TEMPLATE
	}

	templateEngine, err := template.New("weekly").Parse(templateFile)
	if err != nil {
		return output, err
	}

	err = templateEngine.Execute(&output, weekly)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (weekly Weekly) RenderImageUrl() (string, error) {
	return weekly.ImageUrl, nil
}

func (weekly Weekly) SendText() bool {
	return true
}

func (weekly Weekly) SendImage() bool {
	_, err := url.ParseRequestURI(weekly.ImageUrl)
	return err == nil
}

func NewWeekly() Message {
	return Weekly{}
}
