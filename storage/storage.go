package storage

import (
	"context"

	"github.com/404th/book_store/genproto/book_service"
)

type StorageI interface {
	Book() BookI
}

type BookI interface {
	CreateBook(context.Context, *book_service.CreateBookRequest) (*book_service.IDTracker, error)
}
