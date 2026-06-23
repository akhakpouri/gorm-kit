// Package my provides MySQL-specific connection setup for gorm-kit.
package mysql

import (
	"fmt"

	"github.com/akhakpouri/gorm-kit/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect builds a MySQL DSN from cfg and opens a GORM connection.
//
// cfg.Schema is ignored: MySQL has no schema-vs-database distinction (its
// "schema" is a synonym for "database"), so cfg.DbName is the schema.
// cfg.SSLMode is likewise not part of the MySQL DSN here.
func Connect(cfg database.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
