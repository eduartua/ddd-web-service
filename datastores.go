package app

type UserStore interface {
	Authenticate(email, password string) (*User, error)

	// InitiateReset will complete all the model-related tasks
	// to start the password reset process for the user with
	// the provided email address. Once completed, it will
	// return the token, or an error if there was one.
	// InitiateReset(email string) (string, error)

	// CompleteReset will complete all the model-related tasks
	// to complete the password reset process for the user that
	// the token matches, including updating that user's pw.
	// If the token has expired, or if it is invalid for any
	// other reason the ErrTokenInvalid error will be returned.
	// CompleteReset(token, newPw string) (*User, error)
	VerifyEmail(email, token string) error
	UserDB
}

type UserDB interface {
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByUsername(username string) (*User, error)
	ByRemember(token string) (*User, error)

	// Methods for altering users
	Create(user *User) error
	Update(user *User) error
	//Delete(id uint) error
}