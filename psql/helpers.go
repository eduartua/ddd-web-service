package psql

import (
	app "github.com/eduartua/ddd-web-service"
	"gorm.io/gorm"
)

// first will query using the provided gorm.DB and it will
// get the first item returned and place it into dst. If
// nothing is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return app.ErrNotFound
	}
	return err
}
