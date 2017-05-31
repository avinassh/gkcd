package main

import (
	"fmt"
	"os"

	xkcd "github.com/avinassh/gkcd/api"
)

func main() {
	DownloadDir := "/Users/avi/xkcd"
	err := os.MkdirAll(DownloadDir, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	comic, err := xkcd.GetLatest()
	if err != nil {
		fmt.Println(err)
	}
	xkcd.SaveComicWithMeta(comic, DownloadDir)
}
