package secretstr_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/75py/go-secretstr"
	"testing"
)

func TestSecretString_String(t *testing.T) {
	raw := "raw_string"
	ss := secretstr.SecretString(raw)

	testCases := []struct {
		Mode           secretstr.FilterMode
		ExpectedString string
	}{
		{
			Mode:           secretstr.FilterModeFixedString,
			ExpectedString: secretstr.Config.FixedDummyString,
		},
		{
			Mode:           secretstr.FilterModeHide,
			ExpectedString: "**********",
		},
		{
			Mode:           secretstr.FilterModeDisable,
			ExpectedString: raw,
		},
	}

	for _, tc := range testCases {
		secretstr.Config.Mode = tc.Mode
		res := ss.String()
		if res != tc.ExpectedString {
			t.Fatalf("raw=%s, expected:%+v, actual=%s", raw, tc, res)
		}

		for _, format := range []string{"%s", "%v", "%+v", "%#v"} {
			fmtRes := fmt.Sprintf(format, ss)
			if fmtRes != tc.ExpectedString {
				t.Fatalf("raw=%s, actual=%s", raw, fmtRes)
			}
			t.Logf("fmt.Sprintf(\"%s\", ss) returns %s", format, fmtRes)
		}
	}

	secretstr.Config.Mode = 10 // not FilterMode
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("panic doesn't occurred")
		}
		t.Log(err)
	}()
	_ = ss.String()
}

func TestSecretString_MarshalJSON(t *testing.T) {
	type TestStruct struct {
		Str               secretstr.SecretString  `json:"ss"`
		StrEmpty          secretstr.SecretString  `json:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString  `json:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
	}

	src := TestStruct{
		Str:               secretstr.SecretString("raw_string"),
		StrEmpty:          secretstr.SecretString(""),
		StrEmptyOmitEmpty: secretstr.SecretString(""),
		StrNil:            nil,
		StrNilOmitEmpty:   nil,
	}

	testCases := []struct {
		Config   secretstr.SecretStringConfig
		Expected string
	}{
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeFixedString,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `{"ss":"[FILTERED]","ss_empty":"[FILTERED]","ss_nil":null}`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeHide,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `{"ss":"**********","ss_empty":"","ss_nil":null}`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeDisable,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `{"ss":"raw_string","ss_empty":"","ss_nil":null}`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     true,
				Mode:             secretstr.FilterModeFixedString,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `{"ss":"raw_string","ss_empty":"","ss_nil":null}`,
		},
	}

	for _, tc := range testCases {
		secretstr.Config = tc.Config

		b, err := json.Marshal(src)
		if err != nil {
			t.Fatal(err)
		}

		res := string(b)
		if res != tc.Expected {
			t.Fatalf("expected=%s, actual=%s", tc.Expected, res)
		}
		t.Logf("config=%+v, result=%s", tc.Config, res)
	}
}

func TestSecretString_UnmarshalJSON(t *testing.T) {
	type TestStruct struct {
		Str               secretstr.SecretString  `json:"ss"`
		StrEmpty          secretstr.SecretString  `json:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString  `json:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
		Nil               *secretstr.SecretString `json:"nil,omitempty"`
	}

	var ts TestStruct
	err := json.Unmarshal([]byte(`{"ss":"raw_string","ss_empty":"","ss_empty_omitempty":"","ss_nil":null,"ss_nil_omitempty":null}`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if ts.Str != "raw_string" ||
		ts.StrEmpty.RawString() != "" ||
		ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil != nil ||
		ts.StrNilOmitEmpty != nil ||
		ts.Nil != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}

func TestSecretString_MarshalXML(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name                `xml:"test"`
		Str               secretstr.SecretString  `xml:"ss"`
		StrEmpty          secretstr.SecretString  `xml:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString  `xml:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	src := TestStruct{
		Str:               secretstr.SecretString("raw_string"),
		StrEmpty:          secretstr.SecretString(""),
		StrEmptyOmitEmpty: secretstr.SecretString(""),
		StrNil:            nil,
		StrNilOmitEmpty:   nil,
	}

	testCases := []struct {
		Config   secretstr.SecretStringConfig
		Expected string
	}{
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeFixedString,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `<test><ss>[FILTERED]</ss><ss_empty>[FILTERED]</ss_empty></test>`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeHide,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `<test><ss>**********</ss><ss_empty></ss_empty></test>`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     false,
				Mode:             secretstr.FilterModeDisable,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `<test><ss>raw_string</ss><ss_empty></ss_empty></test>`,
		},
		{
			Config: secretstr.SecretStringConfig{
				Marshallable:     true,
				Mode:             secretstr.FilterModeFixedString,
				FixedDummyString: "[FILTERED]",
			},
			Expected: `<test><ss>raw_string</ss><ss_empty></ss_empty></test>`,
		},
	}

	for _, tc := range testCases {
		secretstr.Config = tc.Config

		b, err := xml.Marshal(src)
		if err != nil {
			t.Fatal(err)
		}

		res := string(b)
		if res != tc.Expected {
			t.Fatalf("expected=%s, actual=%s", tc.Expected, res)
		}
		t.Logf("config=%+v, result=%s", tc.Config, res)
	}
}

func TestSecretString_UnmarshalXML(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name                `xml:"test"`
		Str               secretstr.SecretString  `xml:"ss"`
		StrEmpty          secretstr.SecretString  `xml:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString  `xml:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	var ts TestStruct
	err := xml.Unmarshal([]byte(`<test><ss>raw_string</ss><ss_empty></ss_empty><ss_empty_omitempty></ss_empty_omitempty></test>`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if ts.Str != "raw_string" ||
		ts.StrEmpty.RawString() != "" ||
		ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil != nil ||
		ts.StrNilOmitEmpty != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}
