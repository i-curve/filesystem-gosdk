# filesystem-gosdk

<!-- TOC -->

* [filesystem-gosdk](#filesystem-gosdk)
    * [I. Brief](#i-brief)
    * [II. Quick tutorial](#ii-quick-tutorial)
        * [1. create a client](#1-create-a-client)
        * [2. user manage](#2-user-manage)
        * [3. bucket manage](#3-bucket-manage)
        * [4. file manage](#4-file-manage)

<!-- TOC -->

go filesystem sdk: [filesystem project](https://github.com/i-curve/filesystem)

## I. Brief

a go client for [filesystem](https://github.com/i-curve/filesystem).

support user, bucket, file manager.

## II. Quick tutorial

### 1. create a client

```go
import (
gosdk "github.com/i-curve/filesystem-gosdk"
)

var client, _ = gosdk.NewClient("user", "auth", "api_host", "web_host")
```

### 2. user manage

- create a new simple user.

```go
client.AddUser(&gosdk.User{
Name: "user1",
})
```

- delete a user

```go
client.DeleteUser("user1")
```

### 3. bucket manage

- create a new bucket

```go
client.AddBucket(&gosdk.Bucket{
Name: "bucket",
})
```

- delete a bucket

```go
client.DeleteBucket("bucket")
```

### 4. file manage

- upload file

```go
f, _ := os.Open("client.txt")
err := client.UploadFile(gosdk.File{
Bucket: "bucket",
Key:    "sdks/client.go",
}, f)
```

- downlad file

```go
f, err := os.Create("xxx.txt")
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
```

- move file

```go
client.MoveFile(gosdk.File{
Bucket: "bucket",
Key:    "sdks/client.go",
}, gosdk.File{
Bucket: "bucket",
Key:    "sdks/xxx.go",
})
```

- copy file

```go
client.CopyFile(gosdk.File{
Bucket: "bucket",
Key:    "sdks/client.go",
}, gosdk.File{
Bucket: "bucket",
Key:    "sdks/xxx.go",
})
```