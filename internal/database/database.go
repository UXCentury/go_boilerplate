package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-backend/internal/config"
	"go-backend/internal/util"
)

// PGXPool is a wrapper around pgxpool.Pool
var PGXPool *pgxpool.Pool

func NewPostgresDBPool(cfg *config.Database) error {
	sugar := util.Logger.Sugar()
	sugar.Info("Connecting to PostgresSQL...")
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	postgresConfig, configErr := pgxpool.ParseConfig(connString)
	if configErr != nil {
		return configErr
	}

	// Create a PostgresSQL connection pool
	getConnPool, poolErr := pgxpool.NewWithConfig(context.Background(), postgresConfig)
	if poolErr != nil {
		return poolErr
	}
	defer func(getConnPool *pgxpool.Pool, ctx context.Context) {
		//err := getConnPool.Close(ctx)
		//if err != nil {
		//	log.Printf("Query error: %v", err)
		//}
	}(getConnPool, context.Background())

	// Check if the connection is successful
	if err := getConnPool.Ping(context.Background()); err != nil {
		return err
	}

	// Connection successful, perform other operations if needed

	PGXPool = getConnPool

	sugar.Info("Connected to PostgresSQL successfully!")
	return nil
}
