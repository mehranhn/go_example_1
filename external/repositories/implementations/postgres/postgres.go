package repositoryimppostgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mehranhn/go_example_1/db/migrations"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	db sqlx.DB
}

func NewPostgres(databaseSource string) (Postgres, error) {
	db, err := sqlx.Connect("postgres", databaseSource)
	if err != nil {
		return Postgres{}, err
	}

	err = runMigrations(db)
	if err != nil {
		return Postgres{}, err
	}

	postgres := Postgres{
		db: *db,
	}

	return postgres, nil
}

func runMigrations(db *sqlx.DB) error {
	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "."); err != nil {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
