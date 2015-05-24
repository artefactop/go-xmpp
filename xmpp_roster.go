package xmpp

import (
	"encoding/xml"
	"fmt"
)

type Roster []Contact

type Contact struct {
	Remote string
	Name   string
	Group  []string
}

type rosterItem struct {
	XMLName      xml.Name `xml:"jabber:iq:roster item"`
	Jid          string   `xml:",attr"`
	Name         string   `xml:",attr"`
	Subscription string   `xml:",attr"`
	Group        []string
}

// Roster asks for the chat roster.
func (c *Client) Roster() error {
	fmt.Fprintf(c.conn, "<iq from='%s' type='get' id='roster1'><query xmlns='jabber:iq:roster'/></iq>\n", xmlEscape(c.Jid))
	return nil
}
