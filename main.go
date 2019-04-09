package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

var xmlString = `<?xml version="1.0"?>
<?xml-stylesheet href="/xaa.css" type="text/css"?>
<acronym api="xaa" date="2009-03-21T17:10:00-0700">
  <sought ip="89.101.70.140">api</sought>
  <found n="5">
    <acro nym="API" dewey="040"
          added="Wed Jan 01 00:00:00 IST 1992">
      <expan>American Petroleum Institute</expan>
      <comment></comment>
    </acro>
    <acro nym="API" dewey="387" added="2006-02-01">
      <expan>Apiay airport (code)</expan>
      <comment>Colombia</comment>
    </acro>
    <acro nym="API" dewey="040"
          added="Wed Jan 01 00:00:00 IST 1992">
      <expan>Application Programming Interface</expan>
      <comment></comment>
    </acro>
    <acro nym="API" dewey="040"
          added="Wed Jan 01 00:00:00 IST 1992">
      <expan>Applied Precision, Inc.</expan>
      <comment></comment>
    </acro>
    <acro nym="API" dewey="530"
          added="Mon Jul 22 14:24:35 IST 2002">
      <expan>Atmospheric Pressure Ionization</expan>
      <comment></comment>
    </acro>
  </found>
</acronym>`

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

func main() {
	word := os.Args[1]
	uri := fmt.Sprintf("http://acronyms.silmaril.ie/cgi-bin/xaa?%s", word)
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("could not get list")
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	fmt.Println(newStr)
	var acr Acronym
	if err := xml.Unmarshal([]byte(newStr), &acr); err != nil {
		fmt.Println("Did not work.")
		fmt.Println(err)
		return
	}

	fmt.Println("Must have worked.")
	//fmt.Println(acr.XMLName)

	for _, a := range acr.Found.Acros {
		fmt.Println(a.Expan)
	}
}
