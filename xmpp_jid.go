package xmpp

import (
	"strings"
)

func (c *Client) BareJid() string {
	split := strings.Split(c.Jid, "/")
	return split[0]
}
