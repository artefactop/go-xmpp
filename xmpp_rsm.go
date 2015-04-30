package xmpp

import (
	"encoding/xml"
)

type Rsm struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/rsm set"`
	After   string   `xml:"after,omitempty"`
	Before  string   `xml:"before,omitempty"`
	Count   string   `xml:"count,omitempty"`
	Index   string   `xml:"index,omitempty"`
	Last    string   `xml:"last,omitempty"`
	Max     string   `xml:"max,omitempty"`
	First   *RsmFirst
}

type RsmFirst struct {
	XMLName xml.Name `xml:"first"`
	Value   string   `xml:",chardata"`
	Index   string   `xml:"index,attr,omitempty"`
}
