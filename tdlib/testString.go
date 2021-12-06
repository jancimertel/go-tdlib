// AUTOGENERATED - DO NOT EDIT

package tdlib

// TestString A simple object containing a string; for testing only
type TestString struct {
	tdCommon
	Value string `json:"value"` // String
}

// MessageType return the string telegram-type of TestString
func (testString *TestString) MessageType() string {
	return "testString"
}

// NewTestString creates a new TestString
//
// @param value String
func NewTestString(value string) *TestString {
	testStringTemp := TestString{
		tdCommon: tdCommon{Type: "testString"},
		Value:    value,
	}

	return &testStringTemp
}