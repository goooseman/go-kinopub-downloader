package getplaylist

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDownloadPlaylist(t *testing.T) {
	type playlistForTestsType struct {
		Link       string
		ResultFile string
	}
	playlistsForTest := []playlistForTestsType{
		playlistForTestsType{
			Link:       "https://cdn.streambox.in/hls4/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f22369.m3u8?uid=151123",
			ResultFile: "./fixtures/f22369.m3u8",
		},
	}
	for _, playlistForTest := range playlistsForTest {
		result, err := DownloadPlaylist(playlistForTest.Link)
		if err != nil {
			t.Fatal(err, playlistForTest.Link)
		}
		resultFileBuf, err := ioutil.ReadFile(playlistForTest.ResultFile)
		if err != nil {
			t.Fatal(err, playlistForTest.Link)
		}
		fmt.Println(string(resultFileBuf))
		if result != string(resultFileBuf) {
			t.Errorf("Wrong result for playlist %v", playlistForTest.Link)
		}
	}
}
