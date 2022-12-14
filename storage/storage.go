package storage

import (
	"context"

	bs "github.com/404th/book_store/genproto/book_service"
)

type StorageI interface {
	Book() BookI
}

type BookI interface {
	CreateBook(context.Context, *bs.CreateBookRequest) (*bs.IDTracker, error)
	GetAllBooks(context.Context, *bs.GetAllBooksRequest) (*bs.GetAllBooksResponse, error)
	GetBookByID(ctx context.Context, req *bs.GetBookByIDRequest) (*bs.Book, error)
	UpdateBook(ctx context.Context, req *bs.UpdateBookRequest) (*bs.IDTracker, error)
	DeleteBook(ctx context.Context, req *bs.DeleteBookRequest) (*bs.IDTracker, error)
}
