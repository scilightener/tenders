package api

import (
	"strconv"

	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/lib/api/msg"
)

// ParseInt64 parses string s into an *int64 num.
// pName is the name of the parameter that is being parsed.
// If something is wrong, its name appears in the parsing error message.
func ParseInt64(s, pName string, num *int64) error {
	return parse(s, pName, num, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})
}

// ParseInt parses string s into an *int num.
// pName is the name of the parameter that is being parsed.
// If something is wrong, its name appears in the parsing error message.
func ParseInt(s, pName string, num *int) error {
	return parse(s, pName, num, strconv.Atoi)
}

// ParseBool parses string s into a *bool b.
// pName is the name of the parameter that is being parsed.
// If something is wrong, its name appears in the parsing error message.
func ParseBool(s, pName string, b *bool) error {
	return parse(s, pName, b, strconv.ParseBool)
}

// parse parses string s into a *T val.
// pName is the name of the parameter that is being parsed.
// If something is wrong, its name appears in the parsing error message.
// It panics if val is nil.
func parse[T any](s, pName string, val *T, parser func(s string) (T, error)) error {
	if val == nil {
		panic("internal.lib.parse: provided val value is nil")
	}
	if len(s) == 0 {
		return jsn.DecodingError(msg.APIEmptyParameter(pName))
	}
	v, err := parser(s)
	if err != nil {
		return jsn.DecodingError(msg.APIUnacceptableFormat(pName))
	}

	*val = v
	return nil
}
