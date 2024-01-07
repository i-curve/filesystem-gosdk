package filesystem_gosdk

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type File struct {
	Bucket   string `json:"bucket"`   // file bucket
	Key      string `json:"key"`      // file
	Duration int    `json:"duration"` //
}

const (
	URL_FILE_UPLOAD   = "/file/upload"
	URL_FILE_DOWNLOAD = "/file/download?"
	URL_FILE_MOVE     = "/file/move"
	URL_FILE_COPY     = "/file/copy"
	URL_FILE_DELETE   = "/file/delete"
)

// UploadFile
func (c *Client) UploadFile(file File, read io.Reader) error {
	var requestBody bytes.Buffer
	write := multipart.NewWriter(&requestBody)
	write.WriteField("bucket", file.Bucket)
	write.WriteField("key", file.Key)
	write.WriteField("duration", strconv.Itoa(file.Duration))
	writeFile, err := write.CreateFormFile("file", "file")
	if err != nil {
		return fmt.Errorf("create form file error: %w", err)
	}
	io.Copy(writeFile, read)
	write.Close()
	req, _ := http.NewRequest(http.MethodPost, c.apiHost.String()+URL_FILE_UPLOAD, &requestBody)
	req.Header.Set("user", c.user)
	req.Header.Set("auth", c.auth)
	req.Header.Set("Content-Type", write.FormDataContentType())
	_, err = c.parseRes(req)
	return err
}

// DeleteFile
// delete file in bucket
func (c *Client) DeleteFile(file File) error {
	req := c.concatReq(http.MethodDelete, URL_FILE_DELETE, file)
	_, err := c.parseRes(req)
	return err
}

// CopyFile
// copy file sour => dest. only need bucket,key in File
func (c *Client) CopyFile(sour, dest File) error {
	req := c.concatReq(http.MethodPost, URL_FILE_COPY, map[string]string{
		"s_bucket": sour.Bucket,
		"s_key":    sour.Key,
		"d_bucket": dest.Bucket,
		"d_key":    dest.Key,
	})
	_, err := c.parseRes(req)
	return err
}

// MoveFile
// move file sour => dest. only need bucket,key in File
func (c *Client) MoveFile(sour, dest File) error {
	req := c.concatReq(http.MethodPost, URL_FILE_MOVE, map[string]string{
		"s_bucket": sour.Bucket,
		"s_key":    sour.Key,
		"d_bucket": dest.Bucket,
		"d_key":    dest.Key,
	})
	_, err := c.parseRes(req)
	return err
}

// Download
// download file from bucket
func (c *Client) Download(file File) (io.Reader, error) {
	//req := c.concatReq(http.MethodGet, URL_FILE_DOWNLOAD, file)
	params := url.Values{}
	params.Add("bucket", file.Bucket)
	params.Add("key", file.Key)
	req, _ := http.NewRequest(http.MethodGet, c.apiHost.String()+URL_FILE_DOWNLOAD+params.Encode(), nil)
	req.Header.Set("user", c.user)
	req.Header.Set("auth", c.auth)
	return c.parseRes(req)
}
