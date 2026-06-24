/*
Package mysql provides a MySQL driver for GORM, the Go ORM library.
It allows developers to interact with MySQL databases using GORM's ORM capabilities!
Enabling easy database operations and management in Go applications.

	func main() {
		cfg := database.DbConfig{
			Host:     "localhost",
			User:     "root",
			DbName:   "app",
			Port:     3306,
			Password: "secret",
			SSLMode:  "disable", // optional
			Schema:   "public",  // optional
		}

		db, err := mysql.Connect(cfg)
		if err != nil {
			log.Fatal(err) // the app decides what to do; the library never exits
		}

		if err := database.Migrate(db, &User{}, &Order{}); err != nil {
			log.Fatal(err)
		}
	}
*/
package mysql
