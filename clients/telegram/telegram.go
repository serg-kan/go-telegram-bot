package telegram

import (
	"net/http"	
	"net/url"
	"strconv"
	"path"
	"fmt"
)

type Client struct {
	host	  string
	basePath  string
	client 	  http.Client
}

func New(host string, token string) {
	return Client {
		host: host,
		basePath: newBasePath(token),
		client: http.Client{},
	}
	
}

func newBasePath(token string) string {
	return "bot" + token 
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	//adding query parameters
	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	url := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

}

func (c *Client) SendMessage() {

}