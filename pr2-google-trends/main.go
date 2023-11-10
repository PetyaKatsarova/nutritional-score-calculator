package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)
/*
Field Tags: Field tags in Go provide meta-information about the struct fields. They are strings defined in backticks immediately after the field declaration. These tags can be used by various packages in Go for
 different purposes, like validation, formatting, or, as in your case, serialization.XML Serialization/Deserialization: The xml:"title" and xml:"item" tags are used by the Go's XML package to map struct
  fields to XML elements when the struct is serialized to XML, or when XML is deserialized into the struct.
*/

type RSS struct {
	XMLName	xml.Name 		`xml:"rss"` // unmarshal: converting, encoding
	Channel *Channel 		`xml:"channel"` // pick the info from "ht" parts of the xml code
}
type Channel struct {
	Title		string 		`xml:"title"`
	ItemList	[]Item 		`xml:"item"`
}

type Item struct {
	Title		string 		`xml:"title"`
	Link		string 		`xml:"link"`
	Traffic		string 		`xml:"approx_traffic"`
	NewsItem	[]News 		`xml:"news_item"`
}

type News struct {
	Headline		string `xml:"news_item_title"`
	HeadlineLink	string `xml:"news_item_url"`
}

func main() {
	fmt.Println("hello world")
	var r RSS
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("\n", r.Channel.Title)
	fmt.Println("---------------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the trend:", r.Channel.ItemList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemList[i].NewsItem[0].Headline)
		fmt.Println("Link to article:", r.Channel.ItemList[i].NewsItem[0].HeadlineLink)
		fmt.Println("Approximate traffic:", r.Channel.ItemList[i].Traffic)
	}
}

func getGoogleTrends() *http.Response{
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=NL")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readGoogleTrends() []byte { // returns slice of bytes
	resp := getGoogleTrends()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}