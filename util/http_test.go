package util

import (
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)


var examplePrefix = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>
`

func TestHttpGet(t *testing.T) {
	// 200.
	rc, err := HttpGet(&http.Client{}, "http://example.com", nil)
	if err != nil {
		t.Fatalf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	p, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, examplePrefix) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", examplePrefix, s)
	}
}
