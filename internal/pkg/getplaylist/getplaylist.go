package getplaylist

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// DownloadPlaylist recieves a link to download and return it as a string
func DownloadPlaylist(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %v is not %v for %v", resp.StatusCode, http.StatusOK, url)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	result = string(bodyBytes)
	return
}
