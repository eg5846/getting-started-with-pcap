package fritzbox

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getUriAndReadBody(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return readBody(res.Body)
}

func postUriAndReadBody(uri string, response string, username string) ([]byte, error) {
	values := url.Values{"response": {response}, "username": {username}}
	res, err := http.PostForm(uri, values)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return readBody(res.Body)
}

func readBody(r io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return body, nil
}
