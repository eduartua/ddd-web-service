// Package app defines all of our domain types. Things like our representation
// of a user, a widget, and anything else specific to our domain.
//
// This does not include anything specific to the underlying technology. For
// instance, if we wanted to define a UserService interface that described
// functions we could use to interact with a users database that is fine, but
// we wouldn't add any database specific implementations here.
//
// See https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 for
// more info on this.
package app

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID               uuid.UUID
	Country            string
	City               string
	Email              string `gorm:"not null;unique_index"`
	EmailRememberToken string
	EmailVerification  bool
	FirstName          string
	FamilyName         string
	IsSuperAdmin       bool `gorm:"not null;default:false"`
	Lastname           string
	Password           string `gorm:"-"`
	PasswordHash       string `gorm:"not null"`
	Remember           string `gorm:"-"`
	RememberHash       string `gorm:"not null;unique_index"`
	Username           string `gorm:"not null;unique_index"`
}
