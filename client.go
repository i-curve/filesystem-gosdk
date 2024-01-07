package filesystem_gosdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	user    string   `json:"user"`
	auth    string   `json:"auth"`
	apiHost *url.URL `json:"api_host"`
	webHost *url.URL `json:"web_host"`
}

func NewClient(user, auth, apiHost, webHost string) (*Client, error) {
	api_h, err := url.Parse(apiHost)
	if err != nil {
		return nil, fmt.Errorf("parse apihost error: %w", err)
	}
	web_h, err := url.Parse(webHost)
	if err != nil {
		return nil, fmt.Errorf("parse webhost error: %w", err)
	}
	client := &Client{
		user:    user,
		auth:    auth,
		apiHost: api_h,
		webHost: web_h,
	}
	return client, nil
}

func (c *Client) concatReq(method string, urlPath string, in interface{}) *http.Request {
	payload, _ := json.Marshal(in)
	req, _ := http.NewRequest(method, c.apiHost.String()+urlPath, bytes.NewReader(payload))
	req.Header.Set("user", c.user)
	req.Header.Set("auth", c.auth)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (c *Client) parseRes(req *http.Request) (io.Reader, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request error: %w", err)
	}
	if res.StatusCode >= 400 {
		data, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(data))
	}
	return res.Body, nil
}
