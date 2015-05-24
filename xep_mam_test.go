package xmpp

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingArchived(t *testing.T) {
	mamMessage := new(ClientMessageMam)
	mamMessage.Archived = new(MamMessageArchived)
	mamMessage.Archived.By = "by"
	mamMessage.Archived.ID = "id"
	mamMessage.From = "from"
	mamMessage.To = "to"

	raw, _ := xml.Marshal(mamMessage)
	t.Log("raw:", string(raw))

	expected := "<message xmlns=\"jabber:client\" from=\"from\" to=\"to\">" +
		"<archived by=\"by\" id=\"id\"></archived></message>"
	assert.EqualValues(t, expected, string(raw))
}

func TestParsingQuery(t *testing.T) {
	query := new(MamQuery)
	query.MamQueryArchive = new(MamQueryArchive)
	query.MamQueryArchive.QueryId = "queryId"
	query.MamQueryArchive.With = "with"
	query.From = "from"
	query.Type = "get"

	raw, _ := xml.Marshal(query)
	t.Log("raw:", string(raw))

	expected := "<iq xmlns=\"jabber:client\" from=\"from\" type=\"get\">" +
		"<query xmlns=\"urn:xmpp:mam:tmp\" queryid=\"queryId\"><with>with</with></query></iq>"
	assert.EqualValues(t, expected, string(raw))
}
