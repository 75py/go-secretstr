package secretstr

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

var (
	DummyString = "[FILTERED]"
)

type SecretString struct {
	raw          *string
	marshallable bool
}

func NewSecretString(raw *string) SecretString {
	return SecretString{raw: raw, marshallable: false}
}

func NewMarshallableSecureString(raw *string) SecretString {
	return SecretString{raw: raw, marshallable: true}
}

func (ss SecretString) String() string {
	return DummyString
}

// GoString() always returns DummyString.
//
// This function is the implementation of GoStringer interface.
func (ss SecretString) GoString() string {
	return DummyString
}

func (ss *SecretString) RawString() *string {
	if ss == nil {
		return nil
	}
	return ss.raw
}

func (ss *SecretString) IsEmpty() bool {
	return ss == nil || ss.raw == nil || len(*ss.raw) == 0
}

// ==================================================================
// JSON
// ==================================================================

func (ss SecretString) MarshalJSON() ([]byte, error) {
	if ss.raw == nil {
		return json.Marshal(nil)
	}
	if len(*ss.raw) == 0 {
		return json.Marshal("")
	}

	if ss.marshallable {
		return json.Marshal(ss.raw)
	}

	return json.Marshal(DummyString)
}

func (ss *SecretString) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && strings.EqualFold(string(data), "null") {
		ss.raw = nil
		return nil
	}

	var raw string
	e := json.Unmarshal(data, &raw)
	if e != nil {
		return e
	}

	ss.raw = &raw
	return nil
}

// ==================================================================
// XML
// ==================================================================

func (ss SecretString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ss.raw == nil {
		return e.EncodeElement(nil, start)
	}
	if len(*ss.raw) == 0 {
		// <tag></tag>
		return e.EncodeElement("", start)
	}

	if ss.marshallable {
		// <tag>raw</tag>
		return e.EncodeElement(ss.raw, start)
	}

	// <tag>[FILTERED]</tag>
	return e.EncodeElement(DummyString, start)
}

func (ss *SecretString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	raw := ""
	e := d.DecodeElement(&raw, &start)
	if e != nil {
		return e
	}

	ss.raw = &raw
	return nil
}
