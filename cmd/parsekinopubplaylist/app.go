package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goooseman/go-kinopub-downloader/internal/pkg/parseplaylist"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to provide playlist location link as a second parameter")
	}
	playlistLocation := os.Args[1]
	result, err := parseplaylist.ParsePlaylist(playlistLocation)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
