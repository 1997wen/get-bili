package main

import (
	"get-bili/downloader"
	myfmt "get-bili/fmt"
)

func main() {
	request := downloader.InfoRequest{
		Bvids: []string{"BV1Ff4y187q9", "BV1hB4y1B7k5"},
	}

	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, info := range response.Infos {
		myfmt.Logger.Printf("title: %s \n desc: %s", info.Data.Title, info.Data.Desc)
	}

}
