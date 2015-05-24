package xmpp

import (
	"encoding/xml"
)

type Error struct {
	XMLName xml.Name `xml:"error"`
	Code    string   `xml:"code,attr"`
	Type    string   `xml:"type,attr"`
}

type ErrorText struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas text"`
	Value   string   `xml:",chardata"`
}

type ServiceUnavailable struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas service-unavailable"`
	Text    *ErrorText
}

type ServiceUnavailableError struct {
	Error
	ServiceUnavailable
}

type BadRequest struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas bad-request"`
	Text    *ErrorText
}

type BadRequestError struct {
	Error
	BadRequest
}

type NotFound struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas not-found"`
	Text    *ErrorText
}

type NotFoundError struct {
	Error
	NotFound
}

type Confilct struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas conflict"`
	Text    *ErrorText
}

type ConfilctError struct {
	Error
	Confilct
}

type PolicyViolation struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-stanzas policy-violation"`
	Text    *ErrorText
}

type PolicyViolationError struct {
	Error
	PolicyViolation
}
