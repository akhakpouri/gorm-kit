# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

`gorm-kit` is a small, reusable GORM toolkit for Go services. It exists so that
connection setup and migrations don't get hand-rewritten in every new project.
It was extracted from the `commerce-api` repo, where the same `Connect`/`Migrate`
boilerplate kept reappearing.

Scope is deliberately narrow: **connection + migration helpers only.** It is not
an ORM, a query builder, or a repository framework — those stay in the consuming
application.

## Module Layout

Single Go module (`github.com/akhakpouri/gorm-kit`), one `go.mod`, one version tag.

```
gorm-kit/
├── database/        # driver-agnostic core
│   ├── config.go    # DbConfig struct (json tags for app-side unmarshalling)
│   └── migrate.go   # Migrate(db, models...) — AutoMigrate wrapper, driver-neutral
└── pg/              # PostgreSQL-specific
    └── connect.go   # Connect(cfg) (*gorm.DB, error) — builds DSN, opens GORM
```

**Why this split:** only `Connect`/DSN construction is driver-specific. `Migrate`
(`AutoMigrate`) and `DbConfig` are driver-agnostic, so they live in `database/`.
Adding a new driver later = a new sibling package (e.g. `mysql/`) that reuses
`database.Migrate` unchanged. No restructuring required.

## Design Principles

These are the rules that keep this a *library*, not an application. Hold the line on them.

1. **A library never kills the host process.** No `log.Fatal`, no `panic`.
   Functions return `error`; the calling app decides what to do.
2. **The library owns connection, not configuration.** It does not read files,
   env vars, or embeds. Apps load their own config however they like and hand
   over a filled `DbConfig`. (The json tags exist only as a convenience for apps
   that unmarshal from JSON.)
3. **No domain knowledge.** No models, no hard-coded migration lists. `Migrate`
   takes the model list as a variadic from the call site — that's precisely what
   makes it reusable.
4. **No logging.** Progress logging ("running migration…") is an app concern.

## Public API

```go
// database
type DbConfig struct { Host, User, DbName string; Port int; Password, SSLMode, Schema string }
func Migrate(db *gorm.DB, models ...any) error   // AutoMigrate wrapper

// pg
func Connect(cfg database.DbConfig) (*gorm.DB, error)
```

Typical usage in a consuming app:

```go
db, err := pg.Connect(cfg)
if err != nil { log.Fatal(err) }          // the *app* may choose to die
err = database.Migrate(db, &User{}, &Order{} /* this app's models */)
if err != nil { log.Fatal(err) }
```

## Versioning

Single module → single tag (`v0.1.0`, `v1.0.0`, …). Consumers pin via `go get`.
Keep the public API stable; breaking changes mean a major bump.

## Future: multiple drivers

Adding MySQL/SQLite is a **single-module addition** (new `mysql/` package, reuses
`database.Migrate`), not a restructure. Only graduate to module-per-driver
(separate `go.mod` each, tag-prefix versioning, a shared `core` module) if BOTH
hold: (a) ≥2 drivers are actually in use, AND (b) consumers are hurt by carrying
unused driver deps in their module graph. Until then, one module is simpler to
version and consume. Do not split speculatively.

## License

MIT — matches upstream GORM and its Postgres driver, and keeps the library
frictionless to embed anywhere (including private projects).
