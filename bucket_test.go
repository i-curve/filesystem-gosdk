package filesystem_gosdk_test

import (
	gosdk "github.com/i-curve/filesystem-gosdk"
	"testing"
)

func TestBucket(t *testing.T) {
	t.Run("create bucket1", func(t *testing.T) {
		bucket := gosdk.Bucket{
			Name: "bucket1",
		}
		if err := client.AddBucket(&bucket); err != nil {
			t.Fatalf("create bucket1 error: %s", err.Error())
		}
	})
	t.Run("delete bucket1", func(t *testing.T) {
		if err := client.DeleteBucket("bucket1"); err != nil {
			t.Fatalf("delete bucket1 error: %s", err.Error())
		}
	})
}

//func TestBucket_Add(t *testing.T) {
//	if err := client.AddBucket(&gosdk.Bucket{
//		Name: "bucket",
//	}); err != nil {
//		t.Fatalf("create bucket1 error: %s", err.Error())
//	}
//}
