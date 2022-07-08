package client

type Client interface {
	SendText(text string) error
	SendImage(imageUrl string) error
}
