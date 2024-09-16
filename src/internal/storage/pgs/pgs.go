package pgs

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage is a postgres database storage handler.
type Storage struct {
	DBPool *pgxpool.Pool
}

// New returns a new Storage instance.
func New(ctx context.Context, connectionString string) (*Storage, error) {
	const comp = "storage.pgs.New"

	dbPool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return &Storage{DBPool: dbPool}, nil
}

// Close closes the underlying connection to postgres database.
func (s *Storage) Close(_ context.Context) error {
	s.DBPool.Close()
	return nil
}
