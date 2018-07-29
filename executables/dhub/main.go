package main

import (
	"github.com/wplib/go-lib/docker/auth"
	"github.com/wplib/go-lib/docker/image"
)
const img = "wplib/nginx"
func main() {
	t,err := auth.GetAuthToken(img,"pull")
	if err != nil {
		panic(err)
	}
	tt,err := image.GetRemoteImageTagTree(t,img)
	if err != nil {
		panic(err)
	}
	tt.Print()
}
