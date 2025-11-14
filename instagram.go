package main

import "fmt"

type cmn struct {
	comments string
	likes    bool
	shares   int
}

type dm struct {
	layout map[string]map[string]map[string]map[string]cmn
	//user-profile-post-media-image/video-comments-likes-shares
}

func insta() {
	instagram := make(map[string]dm)
	instagram["user"] = dm{layout: make(map[string]map[string]map[string]map[string]cmn)}
	instagram["user"].layout["profile"] = make(map[string]map[string]map[string]cmn)
	instagram["user"].layout["profile"]["post"] = make(map[string]map[string]cmn)
	instagram["user"].layout["profile"]["post"]["Media"] = make(map[string]cmn)
	instagram["user"].layout["profile"]["post"]["Media"]["image/video"] = cmn{
		comments: "good",
		likes:    true,
		shares:   25,
	}

	fmt.Println(instagram)
}

func main() {
	insta()
}
