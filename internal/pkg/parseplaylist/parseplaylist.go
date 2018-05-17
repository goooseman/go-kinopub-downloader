package parseplaylist

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
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
	re := regexp.MustCompile(`\/[^\/]+\.m3u8`)
	rawurl = re.ReplaceAllString(rawurl, ``)

	rawurl = strings.Replace(rawurl, "/hls/", "/pd/", 1)

	uri, err = u.Parse(rawurl)
	if err != nil {
		return
	}
	uri.RawQuery = ""

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

func parsePlaylistFromURI(url *url.URL, movieName string) (result string, err error) {
	body, err := downloadPlaylist(url)
	p, listType, err := m3u8.DecodeFrom(body, true)
	if err != nil {
		return
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
						return "", err
					}
					result += altURL.String() + " --out=\"" + movieName + "." + alternative.Language + ".srt\"" + "\n"
				}
			}
			videoURL, err := contentLink(variant.URI, url)
			if err != nil {
				return "", err
			}
			result += videoURL.String() + " --out=\"" + movieName + ".mp4\"" + "\n"
			result += "\n\n"
		}
	}
	return
}

// ParsePlaylist recieves m3u8 playlist file link as string and returns download links in a string devided with new line character
func ParsePlaylist(link string, movieName string) (result string, err error) {
	uri, err := url.Parse(link)
	if err != nil {
		return
	}
	result, err = parsePlaylistFromURI(uri, movieName)
	return
}
