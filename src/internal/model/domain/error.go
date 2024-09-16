package domain

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

func ValidationErr(err string) ValidationError {
	return ValidationError(err)
}
