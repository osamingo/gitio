package gitio

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GenerateShortURL generates short url by git.io.
func GenerateShortURL(target *url.URL, code string) (string, error) {

	resp, err := http.PostForm("https://git.io", url.Values{
		"url":  []string{target.String()},
		"code": []string{code},
	})
	if err != nil {
		return "", err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("invalid http status code\nstatusCode: %d\nmessage: %s", resp.StatusCode, msg)
	}

	return resp.Header.Get("location"), nil
}
