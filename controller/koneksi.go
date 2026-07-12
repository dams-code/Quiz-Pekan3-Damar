package controller

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"quiz-pekan3-damar/migration"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DBSqlConn *sql.DB

func KoneksiDB() (*sql.DB, error) {

	err := godotenv.Load("config/.env")

	if err != nil {
		log.Println("File config/.env tidak ditemukan")
	}

	psqlCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDBNAME"),
	)

	typedb := "postgres"

	db, err := sql.Open(typedb, psqlCon)

	if err != nil {
		log.Fatalf("Gagal terhubung ke postgres %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke postgres %v", err)
	}

	if err = migration.MigrasiDataBuku(db); err != nil {
		log.Fatalf("Error migrasi data: %v", err)
	}

	fmt.Println("berhasil terhubung ke database postgres")

	return db, nil

}
