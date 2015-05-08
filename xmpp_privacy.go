package xmpp

import (
	"encoding/xml"
	"fmt"
)

const (
	nsPrivacy = "jabber:iq:privacy"
)

type ClientPrivacy struct {
	XMLName xml.Name `xml:"jabber:iq:privacy query"`
	Active  *ClientPrivacyActive
	Default *ClientPrivacyDefault
	List    []ClientPrivacyList `xml:"list"`
}

type ClientPrivacyActive struct {
	XMLName xml.Name `xml:"active"`
	Name    string   `xml:"name,attr,omitempty"`
}

type ClientPrivacyDefault struct {
	XMLName xml.Name `xml:"default"`
	Name    string   `xml:"name,attr,omitempty"`
}

type ClientPrivacyList struct {
	XMLName xml.Name            `xml:"list"`
	Name    string              `xml:"name,attr"`
	Items   []ClientPrivacyItem `xml:"item"`
}

type ClientPrivacyItem struct {
	XMLName xml.Name `xml:"item"`
	Type    string   `xml:"type,attr"` // jid, group, subscription
	Value   string   `xml:"value,attr"`
	Action  string   `xml:"action,attr"` // allow, deny
	Order   string   `xml:"order,attr"`
	//TODO elems: message, presence-in, presence-out, iq
}

func (c *Client) GetPrivacyLists(id string) {
	fmt.Fprintf(c.conn, "<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s'/>"+
		"</iq>\n",
		c.Jid, xmlEscape(id), nsPrivacy)
}

func (c *Client) GetPrivacyList(id, name string) {
	fmt.Fprintf(c.conn, "<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s'>"+
		"<list name='%s'/>"+
		"</query></iq>\n",
		c.Jid, xmlEscape(id), nsPrivacy, name)
}

func (c *Client) SetPrivayList(id, name string, users []string) {
	fmt.Fprintf(c.conn, "<iq from='%s' id='%s' type='set'>"+
		"<query xmlns='%s'>"+
		"<list name='%s'>"+
		generateItems(users)+
		"</list>"+
		"</query></iq>\n",
		c.Jid, xmlEscape(id), nsPrivacy, name)
}

func generateItems(users []string) string {
	var items string
	for i, item := range users {
		items += fmt.Sprintf("<item type='jid' value='%s' action='deny' order='%d'/>", xmlEscape(item), i)
	}
	return items
}
