package xmpp

import (
	"encoding/xml"
)

type clientText struct {
	Lang string `xml:",attr"`
	Body string `xml:"chardata"`
}

type ClientPresence struct {
	XMLName xml.Name `xml:"jabber:client presence"`
	From    string   `xml:"from,attr"`
	ID      string   `xml:"id,attr,omitempty"`
	To      string   `xml:"to,attr,omitempty"`
	Type    string   `xml:"type,attr"` // error, probe, subscribe, subscribed, unavailable, unsubscribe, unsubscribed
	Lang    string   `xml:"lang,attr"`

	Show     string `xml:"show,omitempty"`        // away, chat, dnd, xa
	Status   string `xml:"status,attr,omitempty"` // sb []clientText
	Priority string `xml:"priority,attr"`
	Error    *Error `xml:",omitempty"`
}

type ClientIQ struct { // info/query
	XMLName xml.Name  `xml:"jabber:client iq"`
	From    string    `xml:"from,attr"`
	ID      string    `xml:"id,attr,omitempty"`
	To      string    `xml:"to,attr,omitempty"`
	Type    string    `xml:"type,attr"` // error, get, result, set
	Error   *Error    `xml:",omitempty"`
	Bind    *bindBind `xml:",omitempty"`

	Other *xml.Name `xml:",any"`
}

// RFC 3921  B.1  jabber:client
type ClientMessage struct {
	XMLName xml.Name `xml:"jabber:client message"`
	From    string   `xml:"from,attr"`
	ID      string   `xml:"id,attr,omitempty"`
	To      string   `xml:"to,attr"`
	Type    string   `xml:"type,attr,omitempty"` // chat, error, groupchat, headline or normal

	// These should technically be []clientText, but string is much more convenient.
	Subject string `xml:"subject,omitempty"`
	Body    string `xml:"body,omitempty"`
	Thread  string `xml:"thread,omitempty"`

	Error *Error

	// Any hasn't matched element
	Other []string `xml:",any"`
}

type clientQuery struct {
	Item []rosterItem
}
