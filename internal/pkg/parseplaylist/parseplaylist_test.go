package parseplaylist

import (
	"fmt"
	"testing"
)

func TestParsePlaylist(t *testing.T) {
	type playlistForTestsType struct {
		Name   string
		Link   string
		Result string
	}
	playlistsForTest := []playlistForTestsType{
		playlistForTestsType{
			Name: "mother",
			Link: "https://cdn.streambox.in/hls4/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f22369.m3u8?uid=151123",
			Result: `1920x800
https://cdn.streambox.in/pd/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/d/df/29306.srt --out="mother.ru.srt"
https://cdn.streambox.in/pd/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/4/11/31231.srt --out="mother.fr.srt"
https://cdn.streambox.in/pd/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f/f8/29307.srt --out="mother.en.srt"
https://cdn.streambox.in/pd/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/c/20/K1eBMVK56bCzjL5jp.mp4 --out="mother.mp4"


1280x534
https://cdn.streambox.in/pd/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f/c8/GJoGznysrOj184K9h.mp4 --out="mother.mp4"


720x300
https://cdn.streambox.in/pd/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/6/45/CoWVNDEcjdzSG2GZ2.mp4 --out="mother.mp4"


`,
		},
	}
	for _, playlistForTest := range playlistsForTest {
		result, err := ParsePlaylist(playlistForTest.Link, playlistForTest.Name)
		if err != nil {
			t.Fatal(err, playlistForTest.Name)
		}
		if result != playlistForTest.Result {
			fmt.Println("Result should be", result)
			t.Errorf("Wrong result for playlist %v", playlistForTest.Name)
		}
	}
}
