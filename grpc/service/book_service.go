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

func (b *bookService) CreateBook(ctx context.Context, req *bc.CreateBookRequest) (*bc.IDTracker, error) {
	resp, err := b.strg.Book().CreateBook(ctx, req)
	if err != nil {
		b.log.Error("CreateBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *bookService) GetAllBooks(ctx context.Context, req *bc.GetAllBooksRequest) (*bc.GetAllBooksResponse, error) {
	resp, err := b.strg.Book().GetAllBooks(ctx, req)
	if err != nil {
		b.log.Error("GetAllBooks", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *bookService) GetBookByID(ctx context.Context, req *bc.GetBookByIDRequest) (*bc.Book, error) {
	resp, err := b.strg.Book().GetBookByID(ctx, req)
	if err != nil {
		b.log.Error("GetBookByID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *bookService) UpdateBook(ctx context.Context, req *bc.UpdateBookRequest) (*bc.IDTracker, error) {
	resp, err := b.strg.Book().UpdateBook(ctx, req)
	if err != nil {
		b.log.Error("UpdateBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
