package main

import (
	"github.com/wplib/go-lib/docker/auth"
	"github.com/wplib/go-lib/docker/image"
	"github.com/wplib/go-lib/docker"
	"fmt"
)
const img = "wplib/nginx"
func main() {
	t,err := auth.GetAuthToken(img,"pull")
	if err != nil {
		panic(err)
	}
	tags,err := image.GetImageTags(t,img)
	if err != nil {
		panic(err)
	}
	for _,t := range tags {
		if ( t=="latest") {
			continue
		}
		fmt.Printf("%v\n",t)
	}

	tt,err:= docker.NewTagTree(tags)
	if err!=nil {
		panic(err)
	}
	docker.PrintTree(tt)
}
