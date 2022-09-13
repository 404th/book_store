package postgres

import (
	"context"
	"fmt"

	"github.com/404th/book_store/genproto/book_service"
	bs "github.com/404th/book_store/genproto/book_service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (br *bookRepo) CreateBook(ctx context.Context, cr *book_service.CreateBookRequest) (*book_service.IDTracker, error) {
	query :=
		`
			INSERT INTO books (
				id,
				name,
				author_id,
				about,
				isbn
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5
			);
		`

	id := uuid.New().String()

	if _, err := br.db.Exec(ctx, query, id, cr.Name, cr.AuthorId, cr.About, cr.Isbn); err != nil {
		return nil, fmt.Errorf("error while creating book: %w", err)
	}

	return &bs.IDTracker{
		Id: id,
	}, nil
}
