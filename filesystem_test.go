package filesystem_gosdk_test

import (
	gosdk "github.com/i-curve/filesystem-gosdk"
	"io"
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	f, _ := os.Open("client.go")
	err := client.UploadFile(gosdk.File{
		Bucket: "bucket",
		Key:    "sdks/client.go",
	}, f)
	if err != nil {
		t.Fatalf("file upload error: %s", err.Error())
	}
}

func TestFileCopy(t *testing.T) {
	err := client.CopyFile(gosdk.File{
		Bucket: "bucket",
		Key:    "sdks/client.go",
	}, gosdk.File{
		Bucket: "",
		Key:    "client.go",
	})
	if err != nil {
		t.Fatalf("file copy error: %s", err.Error())
	}
}

func TestFileMove(t *testing.T) {
	if err := client.MoveFile(gosdk.File{
		Bucket: "bucket",
		Key:    "sdks/client.go",
	}, gosdk.File{
		Bucket: "bucket",
		Key:    "sdks/xxx.go",
	}); err != nil {
		t.Fatalf("file move error: %s", err.Error())
	}
}

func TestFileDelete(t *testing.T) {
	if err := client.DeleteFile(gosdk.File{
		Bucket: "bucket",
		Key:    "sdks/client.go",
	}); err != nil {
		t.Fatalf("file copy error: %s", err.Error())
	}
}

func TestFileDownload(t *testing.T) {
	f, err := os.Create("xxx.go")
	if err != nil {
		t.Fatalf("create file error: %s", err.Error())
	}
	r, err := client.Download(gosdk.File{
		Bucket: "bucket",
		Key:    "client.go",
	})
	if err != nil {
		t.Fatalf("file download error: %s", err.Error())
	}
	io.Copy(f, r)
}
