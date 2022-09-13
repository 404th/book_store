package grpc

import (
	"github.com/404th/book_store/config"
	bs "github.com/404th/book_store/genproto/book_service"
	"github.com/404th/book_store/grpc/service"
	"github.com/404th/book_store/pkg/logger"
	"github.com/404th/book_store/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	bs.RegisterBookServiceServer(grpcServer, service.NewBookService(strg, cfg, log))

	reflection.Register(grpcServer)
	return
}
