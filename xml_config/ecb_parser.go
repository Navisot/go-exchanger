package xml_config

import "encoding/xml"

// Envelope keeps the structure of the XML file in order to parse it properly
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Cube    Cube     `xml:"Cube"`
}

// Cube is part of the XML file
type Cube struct {
	XMLName  xml.Name   `xml:"Cube"`
	TimeCube []TimeCube `xml:"Cube"`
}

// TimeCube is part of the XML file that keeps the time attribute
type TimeCube struct {
	XMLName          xml.Name           `xml:"Cube"`
	Time             string             `xml:"time,attr"`
	CurrencyRateCube []CurrencyRateCube `xml:"Cube"`
}

// CurrencyRateCube is part of the XML file that keeps the currency and rate attributes
type CurrencyRateCube struct {
	XMLName  xml.Name `xml:"Cube"`
	Currency string   `xml:"currency,attr"`
	Rate     string   `xml:"rate,attr"`
}
