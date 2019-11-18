/*
	Запуск grpc сервиса
*/

package api

import (
	"github.com/egor1344/banner/rotation_banner/internal/api"
	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	"github.com/egor1344/banner/rotation_banner/internal/domain/services"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initGrpcServer - инициализация grpc сервера
func initGrpcServer() (*api.GrpcBannerServer, error) {
	log.Logger.Info("initGrpcServer")
	dbDsn := viper.GetString("DB_DSN")
	if dbDsn == "" {
		log.Logger.Error("dont set env variable DB_DSN")
	}
	database, err := postgres.InitPgBannerStorage(dbDsn)
	if err != nil {
		log.Logger.Error("databases error ", err)
	}
	database.Log = log.Logger
	grpcService := services.Banner{Database: database, Log: log.Logger}
	return &api.GrpcBannerServer{BannerService: &grpcService, Log: log.Logger}, nil
}

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Info("run grpc server")
		grpcServer, err := initGrpcServer()
		if err != nil {
			log.Logger.Fatal(err)
		}
		log.Logger.Info(grpcServer)
		address := viper.GetString("API_HOST") + ":" + viper.GetString("API_PORT")
		log.Logger.Info(address)
		err = grpcServer.RunServer("tcp", address)
		if err != nil {
			log.Logger.Fatal("Error run server")
		}
	},
}

func init() {
	err := viper.BindEnv("DB_DSN")
	err = viper.BindEnv("API_PORT")
	err = viper.BindEnv("API_HOST")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
