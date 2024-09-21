package tender

import "strings"

type status string

func (s status) canTransitTo(newStatus status) bool {
	switch s {
	case StatusCreated:
		return newStatus == StatusPublished || newStatus == StatusClosed
	case StatusPublished:
		return newStatus == StatusClosed
	case StatusClosed:
		return false
	default:
		return false
	}
}

const (
	StatusCreated   status = "CREATED"
	StatusPublished status = "PUBLISHED"
	StatusClosed    status = "CLOSED"
)

func StatusFromString(s string) (status, error) {
	up := strings.ToUpper(s)
	switch up {
	case string(StatusCreated):
		return StatusCreated, nil
	case string(StatusPublished):
		return StatusPublished, nil
	case string(StatusClosed):
		return StatusClosed, nil
	default:
		return "", ErrInvalidStatus(s)
	}
}
