package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

type Acronym struct {
	XMLName xml.Name `xml:"acronym"`
	Found   Found    `xml:"found"`
}

type Found struct {
	Acros []Acro `xml:"acro"`
}

type Acro struct {
	Expan string `xml:"expan"`
}

func usage(err error) {
	fmt.Println(err)
	fmt.Printf("Usage: %s <acronym> [<acronym>]\n", os.Args[0])
	os.Exit(1)
}

func getMeanings(word string) ([]string, error) {
	uri := fmt.Sprintf("http://acronyms.silmaril.ie/cgi-bin/xaa?%s", word)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("could not get meanings")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyStr := buf.String()

	var acr Acronym
	if err := xml.Unmarshal([]byte(bodyStr), &acr); err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	meanings := make([]string, 0)
	for _, a := range acr.Found.Acros {
		meanings = append(meanings, a.Expan)
	}
	return meanings, nil
}

func main() {
	if len(os.Args) < 2 {
		usage(fmt.Errorf("not enough arguments"))
	}
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		ms, err := getMeanings(word)
		if err != nil {
			usage(fmt.Errorf("could not get meanings: %v", err))
		}
		fmt.Printf("%s:\n", word)
		for _, m := range ms {
			fmt.Printf("\t%s\n", m)
		}
	}
}
