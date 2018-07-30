package main

import (
"github.com/wplib/go-lib/docker/auth"
"github.com/wplib/go-lib/docker/image"
"fmt"
)
const img = "wplib/nginx"
func main() {
	t,err := auth.GetAuthToken(img,"pull")
	if err != nil {
		panic(err)
	}

	tags,err := image.GetRemoteImageTags(t,img)
	if err != nil {
		panic(err)
	}
	for _,t := range tags {
		if ( t=="latest") {
			continue
		}
		fmt.Printf("%v\n",t)
	}

	tt,err := image.GetRemoteImageTagTree(t,img)
	if err != nil {
		panic(err)
	}
	tt.Print()

}

