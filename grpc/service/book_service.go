package service

import (
	"context"

	"github.com/404th/book_store/config"
	bc "github.com/404th/book_store/genproto/book_service"
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
		strg: strg,
		cfg:  cfg,
		log:  log,
	}
}

func (b *bookService) CreateBook(ctx context.Context, bs *bc.CreateBookRequest) (*bc.IDTracker, error) {
	resp, err := b.strg.Book().CreateBook(ctx, bs)
	if err != nil {
		b.log.Error("CreateBook", logger.Any("req", bs), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *bookService) GetAllBooks(ctx context.Context, bs *bc.GetAllBooksRequest) (*bc.GetAllBooksResponse, error) {
	resp, err := b.strg.Book().GetAllBooks(ctx, bs)
	if err != nil {
		b.log.Error("GetAllBooks", logger.Any("req", bs), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
