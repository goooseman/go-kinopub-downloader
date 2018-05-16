package parseplaylist

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/grafov/m3u8"
)

var client = &http.Client{}

func contentLink(rawurl string, u *url.URL) (uri *url.URL, err error) {

	uri, err = u.Parse(rawurl)
	if err != nil {
		return
	}

	if rawurl == u.String() {
		return
	}

	if !uri.IsAbs() { // relative URI
		if rawurl[0] == '/' { // from the root
			rawurl = fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, rawurl)
		} else { // last element
			splitted := strings.Split(u.String(), "/")
			splitted[len(splitted)-1] = rawurl

			rawurl = strings.Join(splitted, "/")
		}
	}

	rawurl, err = url.QueryUnescape(rawurl)
	if err != nil {
		return
	}

	uri, err = u.Parse(rawurl)
	if err != nil {
		return
	}

	return
}

func downloadPlaylist(u *url.URL) (io.ReadCloser, error) {

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal("cms1> " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("cms2> " + err.Error())
		time.Sleep(time.Duration(2) * time.Second)
	}

	if resp.StatusCode != 200 {
		log.Printf("Received HTTP %v for %v\n", resp.StatusCode, u.String())
	}

	return resp.Body, err

}

// ParsePlaylist recieves m3u8 playlist file and returns download links in a string devided with new line character
func ParsePlaylist(url *url.URL) (result string, err error) {
	body, err := downloadPlaylist(url)
	p, listType, err := m3u8.DecodeFrom(body, true)
	if err != nil {
		panic(err)
	}
	switch listType {
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		for _, variant := range masterpl.Variants {
			result += variant.Resolution + "\n"
			for _, alternative := range variant.Alternatives {
				if alternative.Type == "SUBTITLES" {
					altURL, err := contentLink(alternative.URI, url)
					if err != nil {
						panic(err)
					}
					result += altURL.String() + "\n"
				}
			}
			videoURL, err := contentLink(variant.URI, url)
			if err != nil {
				panic(err)
			}
			result += videoURL.String() + "\n"
			result += "\n\n"
		}
	}
	return
}
