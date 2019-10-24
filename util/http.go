//copy from https://github.com/unknwon/com/blob/master/http.go
package util

import (
	"net/http"
	"io"
	"bytes"
	"fmt"
)

var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"

// HttpCall makes HTTP method call.
func HttpCall(client *http.Client, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	for k, vs := range header {
		req.Header[k] = vs
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return resp.Body, nil
	}
	resp.Body.Close()
	if resp.StatusCode == 404 { // 403 can be rate limit error.  || resp.StatusCode == 403 {
		err = fmt.Errorf("resource not found: %s", url)
	} else {
		err = fmt.Errorf("%s %s -> %d", method, url, resp.StatusCode)
	}
	return nil, err
}

// HttpGet gets the specified resource.
// ErrNotFound is returned if the server responds with status 404.
func HttpGet(client *http.Client, url string, header http.Header) (io.ReadCloser, error) {
	return HttpCall(client, "GET", url, header, nil)
}

// HttpPost posts the specified resource.
// ErrNotFound is returned if the server responds with status 404.
func HttpPost(client *http.Client, url string, header http.Header, body []byte) (io.ReadCloser, error) {
	return HttpCall(client, "POST", url, header, bytes.NewBuffer(body))
}