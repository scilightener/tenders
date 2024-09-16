package organization

import "strings"

type _type string

const (
	IE  _type = "IE"
	LLC _type = "LLC"
	JSC _type = "JSC"
)

func TypeFromString(s string) (_type, error) {
	up := strings.ToUpper(s)
	switch up {
	case string(IE):
		return IE, nil
	case string(LLC):
		return LLC, nil
	case string(JSC):
		return JSC, nil
	default:
		return "", ErrInvalidOrganizationType(s)
	}
}
