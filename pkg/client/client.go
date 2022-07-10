package client

type Response struct {
	Identifier string
}

type Chattable interface {
	Value() string
	Method() string
}

type Text struct {
	Text string
}

func (t Text) Value() string {
	return t.Text
}

func (t Text) Method() string {
	return "sendText"
}

type ImageUrl struct {
	ImageUrl string
}

func (i ImageUrl) Value() string {
	return i.ImageUrl
}

func (i ImageUrl) Method() string {
	return "sendImageUrl"
}

type Client interface {
	Send(c Chattable) (Response, error)
	Reply(c Chattable, identifier string) (Response, error)
}
