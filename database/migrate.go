package database

import "gorm.io/gorm"

// Migrate runs GORM's AutoMigrate for the given models.
//
// It is driver-agnostic: it operates on an already-open *gorm.DB regardless of
// the underlying database. The model list is supplied by the caller, so this
// package carries no domain knowledge of any application's schema.
//
// It returns the error from AutoMigrate unchanged; the calling application
// decides how to react (the library never logs or terminates the process).
func Migrate(db *gorm.DB, models ...any) error {
	return db.AutoMigrate(models...)
}
