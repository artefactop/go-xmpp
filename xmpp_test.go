package xmpp

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsingPrivacyXEP(t *testing.T) {
	expected := new(ClientPrivacyIQ)
	expected.From = "romeo@example.net/orchard"
	expected.Type = "get"
	expected.ID = "getlist1"
	expected.Privacy = Privacy{XMLName: xml.Name{nsPrivacy, "query"}}

	stream := "<iq xmlns=\"jabber:client\" from=\"romeo@example.net/orchard\" id=\"getlist1\" type=\"get\">" +
		"<query xmlns=\"jabber:iq:privacy\"></query>" +
		"</iq>"
	reader := strings.NewReader(stream)
	p := xml.NewDecoder(reader)
	name, i, err := next(p)
	t.Log("Name", name)
	t.Log("Interface", i)
	t.Log("err", err)

	assert.EqualValues(t, expected, i)

}
