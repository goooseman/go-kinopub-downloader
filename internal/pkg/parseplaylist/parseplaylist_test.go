package parseplaylist

import (
	"net/url"
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
			Name: "mother!",
			Link: "https://cdn.streambox.in/hls4/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f22369.m3u8?uid=151123",
			Result: `1920x800
https://cdn.streambox.in/hls/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/d/df/29306.srt/index.m3u8
https://cdn.streambox.in/hls/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/4/11/31231.srt/index.m3u8
https://cdn.streambox.in/hls/subtitle/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f/f8/29307.srt/index.m3u8
https://cdn.streambox.in/hls/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/c/20/K1eBMVK56bCzjL5jp.mp4/index-v1.m3u8?uid=0


1280x534
https://cdn.streambox.in/hls/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/f/c8/GJoGznysrOj184K9h.mp4/index-v1.m3u8?uid=0


720x300
https://cdn.streambox.in/hls/kinopub/aWQ9MTUxMTIzOzEzMDA5ODM2ODQ7MDs2NDk4OTMmaD1oUDBoVW9Sb2VMdXNOdURjTU1uMWRnJmU9MTUyNjQ5ODk1NQ/6/45/CoWVNDEcjdzSG2GZ2.mp4/index-v1.m3u8?uid=0


`,
		},
	}
	for _, playlistForTest := range playlistsForTest {
		url, err := url.Parse(playlistForTest.Link)
		if err != nil {
			t.Fatal(err, playlistForTest.Name)
		}
		result, err := ParsePlaylist(url)
		if err != nil {
			t.Fatal(err, playlistForTest.Name)
		}
		if result != playlistForTest.Result {
			t.Errorf("Wrong result for playlist %v", playlistForTest.Name)
		}
	}
}