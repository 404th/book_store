package postgres

import (
	"context"
	"fmt"

	"github.com/404th/book_store/genproto/book_service"
	bs "github.com/404th/book_store/genproto/book_service"
	"github.com/404th/book_store/pkg/helper"
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

func (br *bookRepo) GetAllBooks(ctx context.Context, req *bs.GetAllBooksRequest) (*bs.GetAllBooksResponse, error) {
	var (
		resp   bs.GetAllBooksResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.GetLimit() == 0 {
		req.Limit = 10
	}

	if req.GetSearch() != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM books WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = br.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id, name, author_id, isbn
			FROM books
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := br.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var bk bs.Book

		err = rows.Scan(
			&bk.Id,
			&bk.About,
			&bk.Isbn,
			&bk.AuthorId,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning book err: %w", err)
		}

		resp.Books = append(resp.Books, &bk)
	}

	return &resp, nil
}

func (br *bookRepo) GetBookByID(ctx context.Context, req *bs.GetBookByIDRequest) (*bs.Book, error) {
	var (
		book bs.Book
	)
	query := `
		SELECT (
			id,
			name,
			author_id,
			about,
			isbn
		)
			FROM
				books
			WHERE
				id=$1;
	`

	row := br.db.QueryRow(ctx, query, req.GetId())
	if err := row.Scan(&book.Id, &book.Name, &book.AuthorId, &book.About, &book.Isbn); err != nil {
		return nil, fmt.Errorf("while querying row %w", err)
	}

	return &book, nil
}
