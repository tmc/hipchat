package hipchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"
)

// Client is the primary struct that this package provides. It represents the
// connection to the HipChat API
type Client struct {
	authToken string

	logger *log.Logger
}

func NewClient(authToken string) (*Client, error) {
	return &Client{authToken: authToken}, nil
}

// TraceOn turns on API response tracing to the given logger.
func (c *Client) TraceOn(logger *log.Logger) {
	c.logger = logger
}

// TraceOff turns on API response tracing
func (c *Client) TraceOff() {
	c.logger = nil
}

func (c *Client) trace(args ...interface{}) {
	if c.logger != nil {
		c.logger.Println(args)
	}
}

func (c *Client) prepRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	return req, nil
}

func (c *Client) doSimple(method, endpoint string) (io.ReadCloser, error) {
	return c.do(method, endpoint, nil)
}

func (c *Client) do(method, endpoint string, body io.Reader) (io.ReadCloser, error) {
	u, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	req, err := c.prepRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	c.trace(method, u.String())
	c.trace(resp.Header)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) get(endpoint string, target interface{}) error {
	r, err := c.doSimple("GET", endpoint)
	if err != nil {
		return err
	}
	defer r.Close()
	buf := new(bytes.Buffer)
	err = json.NewDecoder(io.TeeReader(r, buf)).Decode(&target)
	c.trace("response", endpoint, string(buf.Bytes()))
	return err
}

func (c *Client) Rooms() ([]Room, error) {
	// TODO(tmc): this is ignoring pagination
	resp := new(apiResponse)
	err := c.get("room", &resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error
	}

	result := []Room{}
	err = json.Unmarshal(resp.Items, &result)
	if err == nil {
		for _, r := range result {
			r.client = c
		}
	}
	return result, err
}

func (c *Client) GetRoom(name_or_id string) (*Room, error) {
	endpoint := "room/" + name_or_id
	result := new(roomResponse)
	err := c.get(endpoint, result)
	spew.Dump(result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Room.client = c
	return result.Room, err
}

func (c *Client) Users() ([]User, error) {
	// TODO(tmc): this is ignoring pagination
	resp := new(apiResponse)
	err := c.get("user", &resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error
	}

	result := []User{}
	err = json.Unmarshal(resp.Items, &result)
	if err == nil {
		/*
			for _, r := range result {
				//r.client = c
			}
		*/
	}
	return result, err
}

func (c *Client) GetUser(id_or_email string) (*User, error) {
	endpoint := "user/" + id_or_email
	result := new(userResponse)
	err := c.get(endpoint, result)
	spew.Dump(result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	//result.User.client = c
	return result.User, err
}
