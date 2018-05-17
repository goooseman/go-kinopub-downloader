package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goooseman/go-kinopub-downloader/internal/pkg/parseplaylist"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("You need to provide playlist location link as a first parameter and movie name as second")
	}
	playlistLocation := os.Args[1]
	movieName := os.Args[2]
	result, err := parseplaylist.ParsePlaylist(playlistLocation, movieName)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
