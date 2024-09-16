package bid

import "strings"

type status string

func (s status) canTransitTo(newStatus status) bool {
	switch s {
	case StatusCreated:
		return newStatus == StatusPublished || newStatus == StatusCancelled
	case StatusPublished:
		return newStatus == StatusCancelled
	default:
		return false
	}
}

const (
	StatusCreated   status = "CREATED"
	StatusPublished status = "PUBLISHED"
	StatusCancelled status = "CANCELLED"
)

func StatusFromString(s string) (status, error) {
	up := strings.ToUpper(s)
	switch up {
	case string(StatusCreated):
		return StatusCreated, nil
	case string(StatusPublished):
		return StatusPublished, nil
	case string(StatusCancelled):
		return StatusCancelled, nil
	default:
		return "", ErrInvalidStatus(s)
	}
}
