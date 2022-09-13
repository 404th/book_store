package postgres

import (
	"context"

	"github.com/404th/book_store/config"
	"github.com/404th/book_store/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresStore struct {
	db *pgxpool.Pool

	book storage.BookI
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	// getting pool
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: pool,
	}, nil
}

func (s *PostgresStore) Book() storage.BookI {
	if s.book == nil {
		s.book = NewBookRepo(s.db)
	}

	return s.book
}
