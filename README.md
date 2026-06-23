# gorm-kit

Reusable GORM connection + migration helpers for Go services. PostgreSQL today;
MySQL, SQL Server, and Oracle planned.

It exists so connection setup and migrations don't get hand-rewritten in every
new project. Scope is deliberately narrow: **connection + migration helpers
only.** It is not an ORM, a query builder, or a repository framework — those
stay in the consuming application.

## Install

```sh
go get github.com/akhakpouri/gorm-kit
```

## Usage

```go
package main

import (
	"log"

	"github.com/akhakpouri/gorm-kit/database"
	"github.com/akhakpouri/gorm-kit/pg"
)

func main() {
	cfg := database.DbConfig{
		Host:     "localhost",
		User:     "postgres",
		DbName:   "app",
		Port:     5432,
		Password: "secret",
		SSLMode:  "disable", // optional
		Schema:   "public",  // optional
	}

	db, err := pg.Connect(cfg)
	if err != nil {
		log.Fatal(err) // the app decides what to do; the library never exits
	}

	if err := database.Migrate(db, &User{}, &Order{}); err != nil {
		log.Fatal(err)
	}
}
```

The app loads its own config however it likes (file, env, flags) and hands over
a filled `DbConfig`. The `json` tags on `DbConfig` are a convenience for apps
that unmarshal from JSON.

## Public API

```go
// database (driver-agnostic)
type DbConfig struct {
	Host, User, DbName       string
	Port                     int
	Password, SSLMode, Schema string
}
func Migrate(db *gorm.DB, models ...any) error // AutoMigrate wrapper

// pg (PostgreSQL)
func Connect(cfg database.DbConfig) (*gorm.DB, error)
```

## Design principles

1. **A library never kills the host process** — functions return `error`; no
   `log.Fatal`, no `panic`.
2. **The library owns connection, not configuration** — it reads no files, env
   vars, or embeds. The app supplies a filled `DbConfig`.
3. **No domain knowledge** — no models, no hard-coded migration lists. `Migrate`
   takes the model list from the call site.
4. **No logging** — progress logging is an app concern.

## Layout

```
gorm-kit/
├── database/   # driver-agnostic core: DbConfig, Migrate
└── pg/         # PostgreSQL: Connect (DSN + gorm.Open)
```

Adding a driver later = a new sibling package (e.g. `mysql/`) that reuses
`database.Migrate` unchanged.

## License

[MIT](LICENSE).
