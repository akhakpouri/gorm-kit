// Package pg provides PostgreSQL-specific connection setup for gorm-kit.
package pg

import (
	"fmt"
	"strings"

	"github.com/akhakpouri/gorm-kit/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg database.DbConfig) (*gorm.DB, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "host=%s user=%s dbname=%s port=%d password=%s",
		cfg.Host, cfg.User, cfg.DbName, cfg.Port, cfg.Password)
	if cfg.SSLMode != "" {
		fmt.Fprintf(&b, " sslmode=%s", cfg.SSLMode)
	}
	if cfg.Schema != "" {
		fmt.Fprintf(&b, " search_path=%s", cfg.Schema)
	}
	return gorm.Open(postgres.Open(b.String()), &gorm.Config{})
}
