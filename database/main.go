package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"strconv"
	"strings"

	"github.com/davesavic/procat/database/seeders"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
	"github.com/spf13/viper"
)

//go:embed migrations
var migrationFiles embed.FS

//go:embed seeders
var seedFiles embed.FS

type Seeder interface {
	Name() string
	Run() error
}

func main() {
	mustLoadConfig()

	migrationVersion := flag.String("migrate", "", "Migration version")
	seedFile := flag.String("seed", "", "Seed file")

	flag.Parse()

	if migrationVersion != nil && *migrationVersion != "" {
		migrateDB(migrationVersion)
	}

	if seedFile != nil && *seedFile != "" {
		seedDB(seedFile)
	}
}

func seedDB(seedFile *string) {
	seeders := []Seeder{
		seeders.AccessControlSeeder{EmbededFiles: seedFiles},
	}

	for _, seeder := range seeders {
		if *seedFile != "all" && seeder.Name() != *seedFile {
			continue
		}

		err := seeder.Run()
		if err != nil {
			log.Fatalf("Error seeding %s: %v", seeder.Name(), err)
		}
	}
}

func migrateDB(migrationVersion *string) {
	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_DATABASE"),
		),
	)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to database: %w", err))
	}
	defer conn.Close(context.Background())

	migrator, err := migrate.NewMigrator(context.Background(), conn, "migration_version")
	if err != nil {
		panic(fmt.Errorf("Unable to create migrator: %w", err))
	}

	folder, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		panic(fmt.Errorf("Unable to read migrations folder: %w", err))
	}

	err = migrator.LoadMigrations(folder)
	if err != nil {
		panic(fmt.Errorf("Unable to load migrations: %w", err))
	}

	switch true {
	case *migrationVersion == "latest":
		err = migrator.Migrate(context.Background())
		if err != nil {
			panic(fmt.Errorf("Unable to migrate: %w", err))
		}
	case strings.HasPrefix(*migrationVersion, "-"):
		currentVersion, err := migrator.GetCurrentVersion(context.Background())
		if err != nil {
			panic(fmt.Errorf("Unable to get current version: %w", err))
		}

		n, err := strconv.Atoi(*migrationVersion)
		if err != nil {
			panic(fmt.Errorf("Unable to parse version: %w", err))
		}

		err = migrator.MigrateTo(context.Background(), int32(currentVersion)+int32(n))
		if err != nil {
			panic(fmt.Errorf("Unable to migrate: %w", err))
		}

		log.Printf("Rolled back from version %d to version %d", currentVersion, int32(currentVersion)+int32(n))
	default:
		n, err := strconv.Atoi(*migrationVersion)
		if err != nil {
			panic(fmt.Errorf("Unable to parse version: %w", err))
		}

		err = migrator.MigrateTo(context.Background(), int32(n))
		if err != nil {
			panic(fmt.Errorf("Unable to migrate: %w", err))
		}
	}
}

func mustLoadConfig() {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}
