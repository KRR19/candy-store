package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type TopCustomer struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

func main() {
	html := getHtml()
	defer html.Close()
	topCustomers := parseHtml(html)

}

func parseHtml(html io.ReadCloser) []TopCustomer {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}

	topCustomers := []TopCustomer{}
	doc.Find("table.summary td").Each(func(index int, selection *goquery.Selection) {
		selectionAttribute, ok := selection.Attr("x-total-candy")
		if ok {
			selectionText := selection.Text()
			total, _ := strconv.Atoi(selectionAttribute)
			topCustomers = append(topCustomers, TopCustomer{Name: selectionText, TotalSnacks: total})
		} else {
			topCustomers[len(topCustomers)-1].FavouriteSnack = selection.Text()
		}
	})
	return topCustomers
}

func getHtml() io.ReadCloser {
	resp, _ := http.Get("https://candystore.zimpler.net")
	return resp.Body
}
