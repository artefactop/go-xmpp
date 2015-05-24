package xmpp

import (
	"encoding/xml"
	"fmt"
)

const (
	nsPrivacy = "jabber:iq:privacy"
)

type ClientPrivacyIQ struct {
	ClientIQ
	Privacy Privacy
}

type Privacy struct {
	XMLName xml.Name `xml:"jabber:iq:privacy query"`
	Active  *PrivacyActive
	Default *PrivacyDefault
	List    []PrivacyList `xml:"list"`
}

type PrivacyActive struct {
	XMLName xml.Name `xml:"active"`
	Name    string   `xml:"name,attr,omitempty"`
}

type PrivacyDefault struct {
	XMLName xml.Name `xml:"default"`
	Name    string   `xml:"name,attr,omitempty"`
}

type PrivacyList struct {
	XMLName xml.Name      `xml:"list"`
	Name    string        `xml:"name,attr"`
	Items   []PrivacyItem `xml:"item"`
}

type PrivacyItem struct {
	XMLName xml.Name `xml:"item"`
	Type    string   `xml:"type,attr"` // jid, group, subscription
	Value   string   `xml:"value,attr"`
	Action  string   `xml:"action,attr"` // allow, deny
	Order   string   `xml:"order,attr"`
	//TODO elems: message, presence-in, presence-out, iq
}

func GetPrivacyLists(c *Client, id string) {
	raw := fmt.Sprintf("<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s'/>"+
		"</iq>\n",
		c.Jid, XmlEscape(id), nsPrivacy)
	c.SendRaw(raw)
}

func GetPrivacyList(c *Client, id, name string) {
	raw := fmt.Sprintf("<iq from='%s' id='%s' type='get'>"+
		"<query xmlns='%s'>"+
		"<list name='%s'/>"+
		"</query></iq>\n",
		c.Jid, XmlEscape(id), nsPrivacy, name)
	c.SendRaw(raw)
}

func SetPrivayList(c *Client, id, name string, users []string) {
	raw := fmt.Sprintf("<iq from='%s' id='%s' type='set'>"+
		"<query xmlns='%s'>"+
		"<list name='%s'>"+
		generateItems(users)+
		"</list>"+
		"</query></iq>\n",
		c.Jid, XmlEscape(id), nsPrivacy, name)
	c.SendRaw(raw)
}

func generateItems(users []string) string {
	var items string
	for i, item := range users {
		items += fmt.Sprintf("<item type='jid' value='%s' action='deny' order='%d'/>", XmlEscape(item), i)
	}
	return items
}
