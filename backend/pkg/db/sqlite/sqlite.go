package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "social.db")
	if err != nil {
		log.Fatal("Erreur ouverture DB :", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatal("Erreur instance SQLite :", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations/sqlite",
		"sqlite", driver)
	if err != nil {
		log.Fatal("Erreur chargement migration :", err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatal("Erreur lors de l’exécution de la migration :", err)
	}

	log.Println("✅ Migration appliquée avec succès !")
	return db
}
