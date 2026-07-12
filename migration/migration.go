package migration

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var DBMigrations embed.FS

func MigrasiDataBuku(dbParam *sql.DB) error {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: DBMigrations,
		Root:       "sql_migrations",
	}

	tipedatabase := "postgres"

	migrate.SetTable("gorp_migrations_buku")

	hasilMigrasi, err := migrate.Exec(dbParam, tipedatabase, migrations, migrate.Up)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Migrasi datatable buku berhasil, total %d table termigrasi !", hasilMigrasi)
	return nil
}
