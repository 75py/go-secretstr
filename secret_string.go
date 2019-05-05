package secretstr

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// FilterMode is a configuration of SecretString.
// See FilterModeHide, FilterModeFixedString, FilterModeDisable
type FilterMode int

const (
	// FilterModeFixedString : Format SecretString to a fixed string.
	// ex) "foobar"-> "[FILTERED]"
	FilterModeFixedString FilterMode = iota
	// FilterModeHide : Format SecretString to asterisk.
	// ex) "foobar" -> "******"
	FilterModeHide
	// FilterModeDisable : Format SecretString to original string. This flag should not use on release build.
	// ex) "foobar"-> "foobar"
	FilterModeDisable
)

// SecretStringConfig is configurations of SecretString
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

// Config is the instance of SecretStringConfig.
// It's used by all SecretString instances.
var Config = SecretStringConfig{
	Marshallable:     false,
	Mode:             FilterModeFixedString,
	FixedDummyString: "[FILTERED]",
}

// SecretString is an alias for string.
// This type cannot format by normal ways.
type SecretString string

// String : Implementation of Stringer interface.
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

// GoString : Implementation of GoStringer interface.
func (ss SecretString) GoString() string {
	return ss.String()
}

// RawString is convert SecretString to basic string.
// Returned string is not safe for formatting, so be careful to use it.
func (ss SecretString) RawString() string {
	return string(ss)
}

// MarshalJSON overrides the result of json.Marshal().
// If Config.Marshallable = true, the result JSON contains raw strings.
func (ss SecretString) MarshalJSON() ([]byte, error) {
	if Config.Marshallable {
		return json.Marshal(ss.RawString())
	}
	return json.Marshal(ss.String())
}

// MarshalXML overrides the result of xml.Marshal().
// If Config.Marshallable = true, the result XML contains raw strings.
func (ss SecretString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if Config.Marshallable {
		return e.EncodeElement(ss.RawString(), start)
	}
	return e.EncodeElement(ss.String(), start)
}
