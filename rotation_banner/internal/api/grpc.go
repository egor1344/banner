package api

import (
	"context"
	"net"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"github.com/egor1344/banner/rotation_banner/proto/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GrpcBannerServer - Реализует работу с grpc сервером.
type GrpcBannerServer struct {
	BannerService interfaces.Service
	Log           *zap.SugaredLogger
}

// AddBanner - Добавить баннер
func (gbs *GrpcBannerServer) AddBanner(ctx context.Context, in *server.AddBannerRequest) (*server.AddBannerResponse, error) {
	gbs.Log.Info("grpc add banner")
	banner := in.GetBanner()
	err := gbs.BannerService.AddBanner(ctx, banner.Id, banner.Slot.Id)
	if err != nil {
		return &server.AddBannerResponse{Result: &server.AddBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.AddBannerResponse{Result: &server.AddBannerResponse_Status{Status: true}}, nil
}

// DelBanner - Удалить баннер
func (gbs *GrpcBannerServer) DelBanner(ctx context.Context, in *server.DelBannerRequest) (*server.DelBannerResponse, error) {
	gbs.Log.Info("grpc del banner")
	bannerID := in.GetId()
	err := gbs.BannerService.DelBanner(ctx, bannerID)
	if err != nil {
		return &server.DelBannerResponse{Result: &server.DelBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.DelBannerResponse{Result: &server.DelBannerResponse_Status{Status: true}}, nil
}

// CountTransition - Засчитать переход
func (gbs *GrpcBannerServer) CountTransition(ctx context.Context, in *server.CountTransitionRequest) (*server.CountTransitionResponse, error) {
	gbs.Log.Info("grpc count transition")
	bannerId := in.GetIdBanner()
	socDemGroupId := in.GetIdSocDemGroup()
	slotId := in.GetIdSlot()
	err := gbs.BannerService.CountTransition(ctx, bannerId, socDemGroupId, slotId)
	if err != nil {
		return &server.CountTransitionResponse{Result: &server.CountTransitionResponse_Error{Error: "Error"}}, err
	}
	return &server.CountTransitionResponse{Result: &server.CountTransitionResponse_Status{Status: true}}, nil

}

// GetBanner - Выбрать баннер для показа
func (gbs *GrpcBannerServer) GetBanner(ctx context.Context, in *server.GetBannerRequest) (*server.GetBannerResponse, error) {
	gbs.Log.Info("grpc get banner")
	socDemGroupId := in.GetIdSocDemGroup()
	slotId := in.GetIdSlot()
	bannerId, err := gbs.BannerService.GetBanner(ctx, slotId, socDemGroupId)
	if err != nil {
		return &server.GetBannerResponse{Result: &server.GetBannerResponse_Error{Error: "Error"}}, err
	}
	return &server.GetBannerResponse{Result: &server.GetBannerResponse_IdBanner{IdBanner: bannerId}}, nil

}

func (gbs *GrpcBannerServer) RunServer(network, address string) error {
	gbs.Log.Info("run grpc server")
	conn, err := net.Listen(network, address)
	if err != nil {
		gbs.Log.Error("error net listen", err)
	}
	grpc_server := grpc.NewServer()
	server.RegisterRotationBannerServer(grpc_server, gbs)
	err = grpc_server.Serve(conn)
	if err != nil {
		gbs.Log.Error("error serve connection", err)
	}
	return nil
}
