package main

import (
	"io"
	"net/http"
)

func main() {
	html := getHtml()
	defer html.Close()
}

func getHtml() io.ReadCloser {
	resp, _ := http.Get("https://candystore.zimpler.net")
	return resp.Body
}
