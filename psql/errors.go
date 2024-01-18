package psql

import (
	"strings"
)

const (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound       modelError   = "models: resource not found"
	ErrUserIDRequired privateError = "models: user ID is required"

	// ErrUserNotFound is returned when a user is not found
	// in the database
	ErrUserNotFound modelError = "models: user not found"

	ErrTokenInvalid modelError = "models: token provided is not valid"

	// ErrEmailTokenInvalid is returned when the email verification
	// token is expired.
	ErrEmailTokenInvalid modelError = "models: email token provided has expired or is invalid"

	// ErrIDInvalid is returned when an invalid ID is provided
	// to a method like Delete.
	ErrIDInvalid privateError = "models: ID provided was invalid"

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrPasswordIncorrect   modelError = "models: incorrect password provided"
	ErrPasswordIncorrectEs modelError = "models: contraseña incorrecta"

	// ErrPasswordRequired is returned when a create is attempted
	// without a user password provided.
	ErrPasswordRequired modelError = "models: password is required"

	// ErrPasswordTooShort is returned when a user tries to set
	// a password that is less than 8 characters long.
	ErrPasswordTooShort modelError = "models: password must be at least 8 characters long"

	// ErrRememberTooShort is returned when a remember token is
	// not at least 32 bytes
	ErrRememberTooShort modelError = "models: remember token must be at least 32 bytes"

	// ErrRememberRequired is returned when a create or update
	// is attempted without a user remember token hash
	ErrRememberRequired modelError = "models: remember token is required"

	// ErrFirstRememberRequired is returned when a create or update
	// is attempted without a user first remember token hash
	ErrFirstRememberRequired modelError = "models: first remember token is required"

	// ErrEmailRequired is returned when an email address is
	// not provided when creating a user
	ErrEmailRequired modelError = "models: email address is required"

	// ErrEmailInvalid is returned when an email address provided
	// does not match any of our requirements
	ErrEmailInvalid modelError = "models: email address is not valid"

	// ErrEmailTaken is returned when an update or create is attempted
	// with an email address that is already in use.
	ErrEmailTaken modelError = "models: email address is already taken"

	// ErrUserNameRequired is returned when an username is
	// not provided when creating a user
	ErrUserNameRequired modelError = "models: username is required"

	// ErrUsernameInvalid is returned when an email address provided
	// does not match any of our requirements
	ErrUsernameInvalid modelError = `models: username is not valid. It must be a least 5 characters and shouldn't start with: numbers . _ or -`

	// ErrUsernameTaken is returned when an update or create is attempted
	// with an email address that is already in use.
	ErrUsernameTaken modelError = "models: username is already taken"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	// TODO: Improve this using golang.org/x/text/cases Title
	return strings.Replace(string(e), "models: ", "", 1)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
