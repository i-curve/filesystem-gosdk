package filesystem_gosdk_test

import (
	gosdk "github.com/i-curve/filesystem-gosdk"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("add user", func(t *testing.T) {
		user := gosdk.User{
			Name: "test1",
		}
		if err := client.AddUser(&user); err != nil {
			t.Errorf("add errors: %s", err.Error())
		}
	})
	t.Run("delete user", func(t *testing.T) {
		if err := client.DeleteUser("test1"); err != nil {
			t.Errorf("delete user error: %s", err.Error())
		}
	})
}
