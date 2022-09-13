package service

import (
	bc "book_service/genproto/book_service"
	"context"

	"github.com/404th/book_store/config"
	"github.com/404th/book_store/genproto/book_service"
	"github.com/404th/book_store/pkg/logger"
	"github.com/404th/book_store/storage"
)

type bookService struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
	bc.UnimplementedBookServiceServer
}

func NewBookService(strg storage.StorageI, cfg config.Config, log logger.LoggerI) *bookService {
	return &bookService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (b *bookService) CreateBook(ctx context.Context, bs *book_service.CreateBookRequest) (*book_service.IDTracker, error) {
	resp, err := b.strg.Book().CreateBook(ctx, bs)
	if err != nil {
		b.log.Error("CreateBook", logger.Any("req", bs), logger.Error(err))
		return nil, err
	}
}
