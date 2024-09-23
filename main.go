package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/davesavic/procat/repository/query"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	path := ".env"

	flag.StringVar(&path, "env", path, "path to the environment file")
	flag.Parse()

	mustLoadConfig(path)

	db, err := connectToDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	byteUUID := uuid.MustParse("333e4567-e89b-12d3-a456-426614174012")

	q := query.New(db)
	result, err := q.RetrieveProductHierarchy(context.Background(), pgtype.UUID{
		Bytes: byteUUID,
		Valid: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", result)
}

func mustLoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func connectToDB() (*pgxpool.Pool, error) {
	conCfg, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?pool_max_conns=%d",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_DATABASE"),
			viper.GetInt("DB_POOL_MAX_CONNECTIONS"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return pgxpool.NewWithConfig(context.Background(), conCfg)
}
