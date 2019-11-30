package grpc

import (
	"context"
	"net"

	"github.com/egor1344/banner/rotation_banner/pkg/metrics"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"github.com/egor1344/banner/rotation_banner/proto/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// BannerServerGrpc - Реализует работу с grpc сервером.
type BannerServerGrpc struct {
	BannerService interfaces.Service
	Log           *zap.SugaredLogger
}

// AddBanner - Добавить баннер
func (gbs *BannerServerGrpc) AddBanner(ctx context.Context, in *server.AddBannerRequest) (*server.AddBannerResponse, error) {
	gbs.Log.Info("grpc add banner")
	metrics.APICounter.Inc()
	metrics.AddBannerGrpcCounter.Inc()
	banner := in.GetBanner()
	err := gbs.BannerService.AddBanner(ctx, banner.Id, banner.Slot.Id)
	if err != nil {
		return &server.AddBannerResponse{Result: &server.AddBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.AddBannerResponse{Result: &server.AddBannerResponse_Status{Status: true}}, nil
}

// DelBanner - Удалить баннер
func (gbs *BannerServerGrpc) DelBanner(ctx context.Context, in *server.DelBannerRequest) (*server.DelBannerResponse, error) {
	gbs.Log.Info("grpc del banner")
	metrics.APICounter.Inc()
	metrics.DelBannerGrpcCounter.Inc()
	bannerID := in.GetId()
	err := gbs.BannerService.DelBanner(ctx, bannerID)
	if err != nil {
		return &server.DelBannerResponse{Result: &server.DelBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.DelBannerResponse{Result: &server.DelBannerResponse_Status{Status: true}}, nil
}

// CountTransition - Засчитать переход
func (gbs *BannerServerGrpc) CountTransition(ctx context.Context, in *server.CountTransitionRequest) (*server.CountTransitionResponse, error) {
	gbs.Log.Info("grpc count transition")
	metrics.APICounter.Inc()
	metrics.CountTransitionGrpcCounter.Inc()
	bannerID := in.GetIdBanner()
	socDemGroupID := in.GetIdSocDemGroup()
	slotID := in.GetIdSlot()
	err := gbs.BannerService.CountTransition(ctx, bannerID, socDemGroupID, slotID)
	if err != nil {
		return &server.CountTransitionResponse{Result: &server.CountTransitionResponse_Error{Error: "Error"}}, err
	}
	return &server.CountTransitionResponse{Result: &server.CountTransitionResponse_Status{Status: true}}, nil

}

// GetBanner - Выбрать баннер для показа
func (gbs *BannerServerGrpc) GetBanner(ctx context.Context, in *server.GetBannerRequest) (*server.GetBannerResponse, error) {
	gbs.Log.Info("grpc get banner")
	metrics.APICounter.Inc()
	metrics.GetBannerGrpcCounter.Inc()
	socDemGroupID := in.GetIdSocDemGroup()
	slotID := in.GetIdSlot()
	bannerID, err := gbs.BannerService.GetBanner(ctx, slotID, socDemGroupID)
	if err != nil {
		return &server.GetBannerResponse{Result: &server.GetBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.GetBannerResponse{Result: &server.GetBannerResponse_IdBanner{IdBanner: bannerID}}, nil

}

// RunServer - запуск сервера
func (gbs *BannerServerGrpc) RunServer(network, address string) error {
	gbs.Log.Info("run grpc server")
	conn, err := net.Listen(network, address)
	if err != nil {
		gbs.Log.Error("error net listen", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterRotationBannerServer(grpcServer, gbs)
	err = grpcServer.Serve(conn)
	if err != nil {
		gbs.Log.Error("error serve connection", err)
	}
	return nil
}
