package main

import (
	"io"
	"net/http"
)

type TopCustomer struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

func main() {
	html := getHtml()
	defer html.Close()

}

func getHtml() io.ReadCloser {
	resp, _ := http.Get("https://candystore.zimpler.net")
	return resp.Body
}
