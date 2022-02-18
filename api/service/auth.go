package service

type generator interface {
	GenerateToken() (string, error)
}

type Client struct {
	token generator
}

func NewClient(token generator) Client {
	return Client{token}
}

func (c Client) GenerateToken() (string, error) {
	return c.token.GenerateToken()
}
