package api

import (
	"context"
	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"github.com/egor1344/banner/rotation_banner/proto/server"
	"go.uber.org/zap"
)

// GrpcBannerServer - Реализует работу с grpc сервером. Реализует интерфейс RotationBannerServer
type GrpcBannerServer struct {
	BannerService interfaces.Service
	Log           *zap.SugaredLogger
}

// AddBanner - Добавить баннер
func (pgbs *GrpcBannerServer) AddBanner(ctx context.Context, in *server.AddBannerRequest) (*server.AddBannerResponse, error) {
	pgbs.Log.Info("grpc add banner")
	return nil, nil
}

// DelBanner - Удалить баннер
func (pgbs *GrpcBannerServer) DelBanner(ctx context.Context, in *server.DelBannerRequest) (*server.DelBannerResponse, error) {
	pgbs.Log.Info("grpc del banner")
	return nil, nil
}

// CountTransition - Засчитать переход
func (pgbs *GrpcBannerServer) CountTransition(ctx context.Context, in *server.CountTransitionRequest) (*server.CountTransitionResponse, error) {
	pgbs.Log.Info("grpc count transition")
	return nil, nil
}

// GetBanner - Выбрать баннер для показа
func (pgbs *GrpcBannerServer) GetBanner(ctx context.Context, in *server.GetBannerRequest) (*server.GetBannerResponse, error) {
	pgbs.Log.Info("grpc get banner")
	return nil, nil
}
