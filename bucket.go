package filesystem_gosdk

import (
	"net/http"
)

type Bucket struct {
	Name   string `json:"name"`
	BType  int    `json:"b_type"`
	IsTemp bool   `json:"is_temp"`
}

const (
	BTypeRead = iota + 1
	BTypeWrite
	BTypeReadWrite
)

const (
	URL_BUCKET = "/bucket"
)

// AddBucket
func (c *Client) AddBucket(bucket *Bucket) error {
	if bucket.BType == 0 {
		bucket.BType = BTypeReadWrite
	}
	req := c.concatReq(http.MethodPost, URL_BUCKET, bucket)
	_, err := c.parseRes(req)
	return err
}

// DeleteBucket
func (c *Client) DeleteBucket(name string) error {
	req := c.concatReq(http.MethodDelete, URL_BUCKET, map[string]string{
		"name": name,
	})
	_, err := c.parseRes(req)
	return err
}
