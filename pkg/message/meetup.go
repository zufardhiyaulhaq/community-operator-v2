package message

import (
	"bytes"
	"html/template"
	"net/url"
)

type Meetup struct {
	Name            string
	Date            string
	Time            string
	Place           string
	RegistrationUrl string
	ImageUrl        string
	Tags            []string
	Sponsors        []Meetup_Sponsor
	Speakers        []Meetup_Speaker
}

type Meetup_Sponsor struct {
	Name     string
	ImageUrl *string
	Website  *string
}

type Meetup_Speaker struct {
	Name     string
	Position string
	Company  string
	Title    string
	ImageUrl *string
}

func (Meetup Meetup) RenderText(handler HandlerType) (bytes.Buffer, error) {
	var templateFile string
	var output bytes.Buffer

	switch handler {
	case Telegram:
		templateFile = MEETUP_TELEGRAM_TEMPLATE
	case Twitter:
		templateFile = MEETUP_TWITTER_TEMPLATE
	}

	templateEngine, err := template.New("Meetup").Parse(templateFile)
	if err != nil {
		return output, err
	}

	err = templateEngine.Execute(&output, Meetup)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (Meetup Meetup) RenderImageUrl() (string, error) {
	return Meetup.ImageUrl, nil
}

func (Meetup Meetup) SendText() bool {
	return true
}

func (Meetup Meetup) SendImage() bool {
	_, err := url.ParseRequestURI(Meetup.ImageUrl)
	return err == nil
}

func NewMeetup() Message {
	return Meetup{}
}
