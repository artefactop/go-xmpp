package xmpp

import (
	"encoding/xml"
	"fmt"
)

type MamMessageArchived struct {
	XMLName xml.Name `xml:"archived"`
	By      string   `xml:"by,attr"`
	ID      string   `xml:"id,attr"`
}

//VERSION 3
type MamMessageFin struct {
	XMLName xml.Name `xml:"urn:xmpp:mam:0 fin"`
	ID      string   `xml:"id,attr"`
}

type MamQueryArchive struct {
	XMLName xml.Name `xml:"urn:xmpp:mam:tmp query"`
	QueryId string   `xml:"queryid,attr,omitempty"`
	With    string   `xml:"with,omitempty"`
	Set     *Rsm
}

const (
	nsMam  = "urn:xmpp:mam:0" // not supported by mongooseim
	nsMam2 = "urn:xmpp:mam:tmp"
)

func (c *Client) QueryArchiveV2(id, queryid string) {
	fmt.Fprintf(c.conn, "<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s' queryid='%s'/>"+
		"</iq>\n",
		c.Jid, xmlEscape(id), nsMam2, xmlEscape(queryid))
}

func (c *Client) QueryArchive(id, queryid string) {
	fmt.Fprintf(c.conn, "<iq from='%s' id='%s' type='set'>"+
		"<query xmlns='%s' queryid='%s'/>"+
		"</iq>\n",
		c.Jid, xmlEscape(id), nsMam, xmlEscape(queryid))
}
