package secretstr_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"secretstr"
	"strings"
	"testing"
)

func TestSecretString_String(t *testing.T) {
	raw := "raw_string"
	ss := secretstr.NewSecretString(&raw)

	res := ss.String()
	if res != secretstr.DummyString {
		t.Fatalf("raw=%s, expected:%s, actual=%s", raw, secretstr.DummyString, res)
	}

	for _, format := range []string{"%s", "%v", "%+v", "%#v"} {
		fmtRes := fmt.Sprintf(format, ss)
		if strings.Contains(fmtRes, raw) {
			t.Fatalf("raw=%s, actual=%s", raw, fmtRes)
		}
		t.Logf("fmt.Sprintf(\"%s\", ss) returns %s", format, fmtRes)
	}
}

func TestSecretString_RawString(t *testing.T) {
	raw := "raw_string"
	ss := secretstr.NewSecretString(&raw)

	res := ss.RawString()
	if res == nil {
		t.Fatalf("raw=%s, expected:%s, actual=nil", raw, secretstr.DummyString)
	}
	if *res != raw {
		t.Fatalf("raw=%s, expected:%s, actual=%s", raw, secretstr.DummyString, *res)
	}
}

func TestSecretString_IsEmpty(t *testing.T) {
	raw := "raw_string"
	ss := secretstr.NewSecretString(&raw)

	if ss.IsEmpty() {
		t.Fatalf("raw=%s, expected: IsEmpty()=false", raw)
	}

	raw = ""
	ss = secretstr.NewSecretString(&raw)
	if !ss.IsEmpty() {
		t.Fatalf("raw=%s, expected: IsEmpty()=true", raw)
	}

	ss = secretstr.NewSecretString(nil)
	if !ss.IsEmpty() {
		t.Fatal("raw=nil, expected: IsEmpty()=true")
	}
}

// ==================================================================
// JSON
// ==================================================================

func TestSecretString_MarshalJSON(t *testing.T) {
	type TestStruct struct {
		Str               secretstr.SecretString `json:"ss"`
		StrEmpty          secretstr.SecretString `json:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `json:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewSecretString(&raw)
	ss2 := secretstr.NewSecretString(&emptyStr)
	ss3 := secretstr.NewSecretString(nil)
	src := TestStruct{
		Str:               ss,
		StrEmpty:          ss2,
		StrEmptyOmitEmpty: ss2,
		StrNil:            ss3,
		StrNilOmitEmpty:   ss3,
	}

	b, err := json.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `{"ss":"[FILTERED]","ss_empty":"","ss_empty_omitempty":"","ss_nil":null,"ss_nil_omitempty":null}` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalJSON_marshallable(t *testing.T) {
	type TestStruct struct {
		Str               secretstr.SecretString `json:"ss"`
		StrEmpty          secretstr.SecretString `json:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `json:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewMarshallableSecureString(&raw)
	ss2 := secretstr.NewMarshallableSecureString(&emptyStr)
	ss3 := secretstr.NewMarshallableSecureString(nil)
	src := TestStruct{
		Str:               ss,
		StrEmpty:          ss2,
		StrEmptyOmitEmpty: ss2,
		StrNil:            ss3,
		StrNilOmitEmpty:   ss3,
	}

	b, err := json.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `{"ss":"raw_string","ss_empty":"","ss_empty_omitempty":"","ss_nil":null,"ss_nil_omitempty":null}` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalJSON_ptr(t *testing.T) {
	type TestStruct struct {
		StrPtr               *secretstr.SecretString `json:"ss_ptr"`
		StrPtrEmpty          *secretstr.SecretString `json:"ss_ptr_empty"`
		StrPtrEmptyOmitEmpty *secretstr.SecretString `json:"ss_ptr_empty_omitempty,omitempty"`
		StrPtrNil            *secretstr.SecretString `json:"ss_ptr_nil"`
		StrPtrNilOmitEmpty   *secretstr.SecretString `json:"ss_ptr_nil_omitempty,omitempty"`
		Nil                  *secretstr.SecretString `json:"nil"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewSecretString(&raw)
	ss2 := secretstr.NewSecretString(&emptyStr)
	ss3 := secretstr.NewSecretString(nil)
	src := TestStruct{
		StrPtr:               &ss,
		StrPtrEmpty:          &ss2,
		StrPtrEmptyOmitEmpty: &ss2,
		StrPtrNil:            &ss3,
		StrPtrNilOmitEmpty:   &ss3,
		Nil:                  nil,
	}

	b, err := json.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `{"ss_ptr":"[FILTERED]","ss_ptr_empty":"","ss_ptr_empty_omitempty":"","ss_ptr_nil":null,"ss_ptr_nil_omitempty":null,"nil":null}` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalJSON_ptr_marshallable(t *testing.T) {
	type TestStruct struct {
		StrPtr               *secretstr.SecretString `json:"ss_ptr"`
		StrPtrEmpty          *secretstr.SecretString `json:"ss_ptr_empty"`
		StrPtrEmptyOmitEmpty *secretstr.SecretString `json:"ss_ptr_empty_omitempty,omitempty"`
		StrPtrNil            *secretstr.SecretString `json:"ss_ptr_nil"`
		StrPtrNilOmitEmpty   *secretstr.SecretString `json:"ss_ptr_nil_omitempty,omitempty"`
		Nil                  *secretstr.SecretString `json:"nil"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewMarshallableSecureString(&raw)
	ss2 := secretstr.NewMarshallableSecureString(&emptyStr)
	ss3 := secretstr.NewMarshallableSecureString(nil)
	src := TestStruct{
		StrPtr:               &ss,
		StrPtrEmpty:          &ss2,
		StrPtrEmptyOmitEmpty: &ss2,
		StrPtrNil:            &ss3,
		StrPtrNilOmitEmpty:   &ss3,
		Nil:                  nil,
	}

	b, err := json.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `{"ss_ptr":"raw_string","ss_ptr_empty":"","ss_ptr_empty_omitempty":"","ss_ptr_nil":null,"ss_ptr_nil_omitempty":null,"nil":null}` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_UnmarshalJSON(t *testing.T) {
	type TestStruct struct {
		Str               secretstr.SecretString `json:"ss"`
		StrEmpty          secretstr.SecretString `json:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `json:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
	}

	var ts TestStruct
	err := json.Unmarshal([]byte(`{"ss":"raw_string","ss_empty":"","ss_empty_omitempty":"","ss_nil":null,"ss_nil_omitempty":null}`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if *ts.Str.RawString() != "raw_string" ||
		*ts.StrEmpty.RawString() != "" ||
		*ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil.RawString() != nil ||
		ts.StrNilOmitEmpty.RawString() != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}

func TestSecretString_UnmarshalJSON_ptr(t *testing.T) {
	type TestStruct struct {
		Str               *secretstr.SecretString `json:"ss"`
		StrEmpty          *secretstr.SecretString `json:"ss_empty"`
		StrEmptyOmitEmpty *secretstr.SecretString `json:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `json:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `json:"ss_nil_omitempty,omitempty"`
		Nil               *secretstr.SecretString `json:"nil,omitempty"`
	}

	var ts TestStruct
	err := json.Unmarshal([]byte(`{"ss":"raw_string","ss_empty":"","ss_empty_omitempty":"","ss_nil":null,"ss_nil_omitempty":null}`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if *ts.Str.RawString() != "raw_string" ||
		*ts.StrEmpty.RawString() != "" ||
		*ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil.RawString() != nil ||
		ts.StrNilOmitEmpty.RawString() != nil ||
		ts.Nil.RawString() != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}

// ==================================================================
// XML
// ==================================================================

func TestSecretString_MarshalXML(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name               `xml:"test"`
		Str               secretstr.SecretString `xml:"ss"`
		StrEmpty          secretstr.SecretString `xml:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `xml:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewSecretString(&raw)
	ss2 := secretstr.NewSecretString(&emptyStr)
	ss3 := secretstr.NewSecretString(nil)
	src := TestStruct{
		Str:               ss,
		StrEmpty:          ss2,
		StrEmptyOmitEmpty: ss2,
		StrNil:            ss3,
		StrNilOmitEmpty:   ss3,
	}

	b, err := xml.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `<test><ss>[FILTERED]</ss><ss_empty></ss_empty><ss_empty_omitempty></ss_empty_omitempty></test>` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalXML_marshallable(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name               `xml:"test"`
		Str               secretstr.SecretString `xml:"ss"`
		StrEmpty          secretstr.SecretString `xml:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `xml:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewMarshallableSecureString(&raw)
	ss2 := secretstr.NewMarshallableSecureString(&emptyStr)
	ss3 := secretstr.NewMarshallableSecureString(nil)
	src := TestStruct{
		Str:               ss,
		StrEmpty:          ss2,
		StrEmptyOmitEmpty: ss2,
		StrNil:            ss3,
		StrNilOmitEmpty:   ss3,
	}

	b, err := xml.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `<test><ss>raw_string</ss><ss_empty></ss_empty><ss_empty_omitempty></ss_empty_omitempty></test>` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalXML_ptr(t *testing.T) {
	type TestStruct struct {
		XMLName              xml.Name                `xml:"test"`
		StrPtr               *secretstr.SecretString `xml:"ss_ptr"`
		StrPtrEmpty          *secretstr.SecretString `xml:"ss_ptr_empty"`
		StrPtrEmptyOmitEmpty *secretstr.SecretString `xml:"ss_ptr_empty_omitempty,omitempty"`
		StrPtrNil            *secretstr.SecretString `xml:"ss_ptr_nil"`
		StrPtrNilOmitEmpty   *secretstr.SecretString `xml:"ss_ptr_nil_omitempty,omitempty"`
		Nil                  *secretstr.SecretString `xml:"nil"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewSecretString(&raw)
	ss2 := secretstr.NewSecretString(&emptyStr)
	ss3 := secretstr.NewSecretString(nil)
	src := TestStruct{
		StrPtr:               &ss,
		StrPtrEmpty:          &ss2,
		StrPtrEmptyOmitEmpty: &ss2,
		StrPtrNil:            &ss3,
		StrPtrNilOmitEmpty:   &ss3,
		Nil:                  nil,
	}

	b, err := xml.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `<test><ss_ptr>[FILTERED]</ss_ptr><ss_ptr_empty></ss_ptr_empty><ss_ptr_empty_omitempty></ss_ptr_empty_omitempty></test>` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_MarshalXML_ptr_marshallable(t *testing.T) {
	type TestStruct struct {
		XMLName              xml.Name                `xml:"test"`
		StrPtr               *secretstr.SecretString `xml:"ss_ptr"`
		StrPtrEmpty          *secretstr.SecretString `xml:"ss_ptr_empty"`
		StrPtrEmptyOmitEmpty *secretstr.SecretString `xml:"ss_ptr_empty_omitempty,omitempty"`
		StrPtrNil            *secretstr.SecretString `xml:"ss_ptr_nil"`
		StrPtrNilOmitEmpty   *secretstr.SecretString `xml:"ss_ptr_nil_omitempty,omitempty"`
		Nil                  *secretstr.SecretString `xml:"nil"`
	}

	raw := "raw_string"
	emptyStr := ""
	ss := secretstr.NewMarshallableSecureString(&raw)
	ss2 := secretstr.NewMarshallableSecureString(&emptyStr)
	ss3 := secretstr.NewMarshallableSecureString(nil)
	src := TestStruct{
		StrPtr:               &ss,
		StrPtrEmpty:          &ss2,
		StrPtrEmptyOmitEmpty: &ss2,
		StrPtrNil:            &ss3,
		StrPtrNilOmitEmpty:   &ss3,
		Nil:                  nil,
	}

	b, err := xml.Marshal(src)

	if err != nil {
		t.Fatal(err)
	}

	res := string(b)
	if res != `<test><ss_ptr>raw_string</ss_ptr><ss_ptr_empty></ss_ptr_empty><ss_ptr_empty_omitempty></ss_ptr_empty_omitempty></test>` {
		t.Fatal(res)
	}

	t.Log(res)
}

func TestSecretString_UnmarshalXML(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name               `xml:"test"`
		Str               secretstr.SecretString `xml:"ss"`
		StrEmpty          secretstr.SecretString `xml:"ss_empty"`
		StrEmptyOmitEmpty secretstr.SecretString `xml:"ss_empty_omitempty,omitempty"`
		StrNil            secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	var ts TestStruct
	err := xml.Unmarshal([]byte(`<test><ss>raw_string</ss><ss_empty></ss_empty><ss_empty_omitempty></ss_empty_omitempty></test>`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if *ts.Str.RawString() != "raw_string" ||
		*ts.StrEmpty.RawString() != "" ||
		*ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil.RawString() != nil ||
		ts.StrNilOmitEmpty.RawString() != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}

func TestSecretString_UnmarshalXML_ptr(t *testing.T) {
	type TestStruct struct {
		XMLName           xml.Name               `xml:"test"`
		Str               *secretstr.SecretString `xml:"ss"`
		StrEmpty          *secretstr.SecretString `xml:"ss_empty"`
		StrEmptyOmitEmpty *secretstr.SecretString `xml:"ss_empty_omitempty,omitempty"`
		StrNil            *secretstr.SecretString `xml:"ss_nil"`
		StrNilOmitEmpty   *secretstr.SecretString `xml:"ss_nil_omitempty,omitempty"`
	}

	var ts TestStruct
	err := xml.Unmarshal([]byte(`<test><ss>raw_string</ss><ss_empty></ss_empty><ss_empty_omitempty></ss_empty_omitempty></test>`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if *ts.Str.RawString() != "raw_string" ||
		*ts.StrEmpty.RawString() != "" ||
		*ts.StrEmptyOmitEmpty.RawString() != "" ||
		ts.StrNil.RawString() != nil ||
		ts.StrNilOmitEmpty.RawString() != nil {
		t.Fatal(ts)
	}
	t.Logf("%+v", ts)
}
