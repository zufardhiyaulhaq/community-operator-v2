package client

type Response struct {
	Identifier string
}

type Chattable interface {
	Value() map[string]string
	Method() string
}

type Text struct {
	Text string
}

func (t Text) Value() map[string]string {
	return map[string]string{
		"text": t.Text,
	}
}

func (t Text) Method() string {
	return "sendText"
}

type ImageUrl struct {
	ImageUrl string
}

func (i ImageUrl) Value() map[string]string {
	return map[string]string{
		"imageUrl": i.ImageUrl,
	}
}

func (i ImageUrl) Method() string {
	return "sendImageUrl"
}

type TextAndImageUrl struct {
	Text     string
	ImageUrl string
}

func (tiu TextAndImageUrl) Value() map[string]string {
	return map[string]string{
		"text":     tiu.Text,
		"imageUrl": tiu.ImageUrl,
	}
}

func (tiu TextAndImageUrl) Method() string {
	return "sendTextAndImageUrl"
}

type Client interface {
	Send(c Chattable) (Response, error)
	Reply(c Chattable, identifier string) (Response, error)
}
