package filesystem_gosdk

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Auth  string `json:"auth"`
	UType int    `json:"u_type"`
}

const (
	UTypeSystem = iota + 1 // system
	UTypeAdmin             // admin
	UTypeUser              // user
)

const (
	URL_USER = "/user"
)

// AddUser
// only need user name to create
func (c *Client) AddUser(user *User) error {
	if user.UType < UTypeSystem || user.UType > UTypeUser {
		user.UType = UTypeUser
	}
	req := c.concatReq(http.MethodPost, URL_USER, user)
	if r, err := c.parseRes(req); err != nil {
		return err
	} else {
		data, _ := io.ReadAll(r)
		return json.Unmarshal(data, user)
	}
}

// DeleteUser
// only need user name
func (c *Client) DeleteUser(name string) error {
	req := c.concatReq(http.MethodDelete, URL_USER, map[string]string{
		"name": name,
	})
	_, err := c.parseRes(req)
	return err
}
