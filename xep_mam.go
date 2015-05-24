package xmpp

import (
	"encoding/xml"
	"fmt"
)

const (
	nsMam  = "urn:xmpp:mam:0" // not supported by mongooseim
	nsMam2 = "urn:xmpp:mam:tmp"
)

type ClientMessageMam struct {
	ClientMessage
	// XEP-0313
	Archived *MamMessageArchived
	Fin      *MamMessageFin //version3
}

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

type MamQuery struct {
	ClientIQ
	MamQueryArchive *MamQueryArchive
}

type MamQueryArchive struct {
	XMLName xml.Name `xml:"urn:xmpp:mam:tmp query"`
	QueryId string   `xml:"queryid,attr,omitempty"`
	With    string   `xml:"with,omitempty"`
	Set     *Rsm
}

func QueryArchiveV2(c *Client, id, queryid string) {
	raw := fmt.Sprintf("<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s' queryid='%s'/>"+
		"</iq>\n",
		c.Jid, XmlEscape(id), nsMam2, XmlEscape(queryid))
	c.SendRaw(raw)
}

func QueryArchive(c *Client, id, queryid string) {
	raw := fmt.Sprintf("<iq from='%s' id='%s' type='set'>"+
		"<query xmlns='%s' queryid='%s'/>"+
		"</iq>\n",
		c.Jid, XmlEscape(id), nsMam, XmlEscape(queryid))
	c.SendRaw(raw)
}
