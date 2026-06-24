/*
Package pg provides a PostgreSQL driver for GORM, the Go ORM library.
It allows developers to interact with PostgreSQL databases using GORM's ORM capabilities!
Enabling easy database operations and management in Go applications.

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
*/
package pg
