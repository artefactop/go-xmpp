// Copyright 2013 Flo Lauber <dev@qatfy.at>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(flo):
//   - support password protected MUC rooms
//   - cleanup signatures of join/leave functions
package xmpp

import (
	"fmt"
)

const (
	nsMUC     = "http://jabber.org/protocol/muc"
	nsMUCUser = "http://jabber.org/protocol/muc#user"
)

// xep-0045 7.2
func JoinMUC(c *Client, jid, nick string) {
	if nick == "" {
		nick = c.Jid
	}
	raw := fmt.Sprintf("<presence to='%s/%s'>\n"+
		"<x xmlns='%s' />\n"+
		"</presence>",
		XmlEscape(jid), XmlEscape(nick), nsMUC)
	c.SendRaw(raw)
}

// xep-0045 7.2.6
func JoinProtectedMUC(c *Client, jid, nick string, password string) {
	if nick == "" {
		nick = c.Jid
	}
	raw := fmt.Sprintf("<presence to='%s/%s'>\n"+
		"<x xmlns='%s'>\n"+
		"<password>%s</password>\n"+
		"</x>\n"+
		"</presence>",
		XmlEscape(jid), XmlEscape(nick), nsMUC, XmlEscape(password))
	c.SendRaw(raw)
}

// xep-0045 7.14
func LeaveMUC(c *Client, jid string) {
	raw := fmt.Sprintf("<presence from='%s' to='%s' type='unavailable' />",
		c.Jid, XmlEscape(jid))
	c.SendRaw(raw)
}
