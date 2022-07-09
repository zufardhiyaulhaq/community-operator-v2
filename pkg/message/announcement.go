package message

import (
	"bytes"
	"html/template"
	"net/url"
)

type Announcement struct {
	Subject  string
	Body     string
	ImageUrl string
	Tags     []string
}

func (Announcement Announcement) RenderText(handler HandlerType) (bytes.Buffer, error) {
	var templateFile string
	var output bytes.Buffer

	switch handler {
	case Telegram:
		templateFile = ANNOUNCEMENT_TELEGRAM_TEMPLATE
	}

	templateEngine, err := template.New("Announcement").Parse(templateFile)
	if err != nil {
		return output, err
	}

	err = templateEngine.Execute(&output, Announcement)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (Announcement Announcement) RenderImageUrl() (string, error) {
	return Announcement.ImageUrl, nil
}

func (Announcement Announcement) SendText() bool {
	return true
}

func (Announcement Announcement) SendImage() bool {
	_, err := url.ParseRequestURI(Announcement.ImageUrl)
	return err == nil
}

func NewAnnouncement() Message {
	return Announcement{}
}
