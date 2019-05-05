package secretstr

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

type FilterMode int

const (
	// Format SecretString to a fixed string.
	// ex) "foobar"-> "[FILTERED]"
	FilterModeFixedString FilterMode = iota
	// Format SecretString to asterisk.
	// ex) "foobar" -> "******"
	FilterModeHide
	// Format SecretString to original string. This flag should not use on release build.
	// ex) "foobar"-> "foobar"
	FilterModeDisable
)

type SecretStringConfig struct {
	// If true, The members of structs can marshal by json.Marshal() or xml.Marshal().
	// Default value: false
	Marshallable bool
	// See FilterModeHide, FilterModeFixedString, FilterModeDisable
	// Default value: FilterModeFixedString
	Mode FilterMode
	// if Mode == FilterModeFixedString, this string is used for formatting.
	// Default value: "[FILTERED]"
	FixedDummyString string
}

var Config = SecretStringConfig{
	Marshallable:     false,
	Mode:             FilterModeFixedString,
	FixedDummyString: "[FILTERED]",
}

// SecretString is an alias for string.
// This type cannot format by normal ways.
type SecretString string

// Implementation of Stringer interface.
func (ss SecretString) String() string {
	switch Config.Mode {
	case FilterModeHide:
		return strings.Repeat("*", len(ss))
	case FilterModeFixedString:
		return Config.FixedDummyString
	case FilterModeDisable:
		return ss.RawString()
	default:
		panic(fmt.Errorf("unexpected FilterMode. config=%#v", Config))
	}
}

// Implementation of GoStringer interface.
func (ss SecretString) GoString() string {
	return ss.String()
}

// Convert to basic string.
// Returned string is not safe for formatting, so be careful to use it.
func (ss SecretString) RawString() string {
	return string(ss)
}

func (ss SecretString) MarshalJSON() ([]byte, error) {
	if Config.Marshallable {
		return json.Marshal(ss.RawString())
	}
	return json.Marshal(ss.String())
}

func (ss SecretString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if Config.Marshallable {
		return e.EncodeElement(ss.RawString(), start)
	}
	return e.EncodeElement(ss.String(), start)
}
