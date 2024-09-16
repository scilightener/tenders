package employee

import (
	"time"

	"github.com/google/uuid"
)

// Employee represents an individual entity.
type Employee struct {
	id        uuid.UUID
	username  string
	firstName *string
	lastName  *string
	createdAt time.Time
	updatedAt time.Time
}

// New creates a new instance of the Employee.
func New(
	id uuid.UUID, username string, firstName, lastName *string,
	createdAt, updatedAt time.Time,
) (*Employee, error) {
	if len(username) > 50 {
		return nil, ErrUsernameTooLong
	} else if len(username) == 0 {
		return nil, ErrUsernameEmpty
	}
	return &Employee{
		id:        id,
		username:  username,
		firstName: firstName,
		lastName:  lastName,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

// ID returns the employee's ID.
func (u *Employee) ID() uuid.UUID {
	return u.id
}

// Username returns the employee's username.
func (u *Employee) Username() string {
	return u.username
}

// FirstName returns the employee's first name.
func (u *Employee) FirstName() *string {
	return u.firstName
}

// LastName returns the employee's last name.
func (u *Employee) LastName() *string {
	return u.lastName
}

// CreatedAt returns the employee's creation date.
func (u *Employee) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt returns the employee's last update date.
func (u *Employee) UpdatedAt() time.Time {
	return u.updatedAt
}
