package bid

import "strings"

type authorType string

const (
	AuthorTypeOrganization authorType = "ORGANIZATION"
	AuthorTypeUser         authorType = "USER"
)

func AuthorTypeFromString(s string) (authorType, error) {
	up := strings.ToUpper(s)
	switch up {
	case string(AuthorTypeOrganization):
		return AuthorTypeOrganization, nil
	case string(AuthorTypeUser):
		return AuthorTypeUser, nil
	default:
		return "", ErrInvalidAuthorType(s)
	}
}
