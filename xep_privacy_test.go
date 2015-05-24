package xmpp

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingGetLists(t *testing.T) {
	iq := new(ClientPrivacyIQ)
	iq.From = "romeo@example.net/orchard"
	iq.Type = "get"
	iq.ID = "getlist1"

	raw, _ := xml.Marshal(iq)
	t.Log("raw:", string(raw))

	expected := "<iq xmlns=\"jabber:client\" from=\"romeo@example.net/orchard\" id=\"getlist1\" type=\"get\">" +
		"<query xmlns=\"jabber:iq:privacy\"></query>" +
		"</iq>"
	assert.EqualValues(t, expected, string(raw))
}
