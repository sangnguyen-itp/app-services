package database

import (
	gorm_migrate "app-services/migration/gorm-migrate"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var DBContainer *Database

func newDBContainer(db *gorm.DB) *Database {
	return &Database{DB: db}
}

func (c *Database) migrateDatabase() {
	if os.Getenv("MIGRATION") == "yes" {
		gorm_migrate.NewMigrate(c.DB)
	}
}

func OpenConnection() {
	// initialize connection
	conn, err := DBContainer.connectDatabase()
	if err != nil {
		panic(err)
	}

	// init DB and migrate
	DBContainer = newDBContainer(conn)
	DBContainer.migrateDatabase()
}

func (c *Database) connectDatabase() (*gorm.DB, error) {

	// get ENV
	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	PORT := os.Getenv("PORT")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		USER,
		PASSWORD,
		HOST,
		PORT,
		DBNAME)

	return gorm.Open(postgres.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
