package filesystem_gosdk_test

import gosdk "github.com/i-curve/filesystem-gosdk"

var client, _ = gosdk.NewClient("system", "53fc91f6fe8ab61ca9bf5ce7c159c0c9", "http://localhost:8001", "http://localhost:8000")
