package psql

import (
	"regexp"
	"time"

	app "github.com/eduartua/ddd-web-service"
	"github.com/eduartua/ddd-web-service/hash"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserStore(db *gorm.DB, pepper, hmacKey string) app.UserStore {
	ug := &userGorm{db}
	hmac := hash.NewHMAC(hmacKey)
	uv := newUserValidator(ug, hmac, pepper)
	return &userStore{
		UserDB:    uv,
		pepper:    pepper,
		pwResetDB: newPwResetValidator(&pwResetGorm{db}, hmac),
	}
}

var _ app.UserStore = &userStore{}

type userStore struct {
	app.UserDB
	pepper    string
	pwResetDB pwResetDB
}

func (us *userStore) Authenticate(email, password string) (*app.User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+us.pepper),
	)
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrPasswordIncorrect
	default:
		return nil, err
	}
}

func (us *userStore) VerifyEmail(email, token string) error {
	user, err := us.ByEmail(email)
	if err != nil {
		return err
	}
	if token == "" || time.Now().Sub(user.UpdatedAt) > (24*time.Hour) || token != user.EmailRememberToken {
		return ErrEmailTokenInvalid
	}

	user.EmailVerification = true
	err = us.Update(user)
	if err != nil {
		return err
	}
	return nil
}

type userGorm struct {
	db *gorm.DB
}

// ByRemember looks up a user with the given remember token
// and returns that user. This method expects the remember
// token to already be hashed.
// Errors are the same as ByEmail.
func (ug *userGorm) ByRemember(rememberHash string) (*app.User, error) {
	var user app.User
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ug *userGorm) ByEmail(email string) (*app.User, error) {
	var user app.User

	db := ug.db.Where("email = ?", email)
	err := first(db, &user)

	return &user, err
}

// Create will create the provided user and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (ug *userGorm) Create(user *app.User) error {
	return ug.db.Create(user).Error
}

// CreateAnchorUser will create the provided user and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (ug *userGorm) CreateAnchorUser(user *app.User) error {
	return ug.db.Create(user).Error
}

// ByID will look up a user with the provided ID.
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
// If there is another error, we will return an error with
// more information about what went wrong. This may not be
// an error generated by the models package.
//
// As a general rule, any error but ErrNotFound should
// probably result in a 500 error.
func (ug *userGorm) ByID(id uint) (*app.User, error) {
	var user app.User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)

	return &user, err
}

func (ug *userGorm) ByUsername(username string) (*app.User, error) {
	var user app.User

	db := ug.db.Where("username = ?", username)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) Update(user *app.User) error {
	return ug.db.Save(user).Error
}

// STARTS USER VALIDATION CODE

type userValFunc func(*app.User) error

func runUserValFuncs(user *app.User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

var _ app.UserDB = &userValidator{}

func newUserValidator(udb app.UserDB, hmac hash.HMAC, pepper string) *userValidator {
	return &userValidator{
		UserDB:        udb,
		hmac:          hmac,
		emailRegex:    regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
		usernameRegex: regexp.MustCompile(`[^0-9._\-][a-z0-9._%+\-]{4,28}$`),
		pepper:        pepper,
	}
}

type userValidator struct {
	app.UserDB
	hmac          hash.HMAC
	emailRegex    *regexp.Regexp
	usernameRegex *regexp.Regexp
	pepper        string
}
