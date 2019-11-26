/*
	Запуск rest api сервиса
*/

package api

import (
	"github.com/egor1344/banner/rotation_banner/internal/api/rest"
	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	"github.com/egor1344/banner/rotation_banner/internal/domain/services"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initRestServer - инициализация grpc сервера
func initRestServer() (*rest.RestBannerServer, error) {
	log.Logger.Info("initGrpcServer")
	dbDsn := viper.GetString("DB_DSN")
	if dbDsn == "" {
		log.Logger.Error("dont set env variable DB_DSN")
	}
	database, err := postgres.InitPgBannerStorage(dbDsn)
	if err != nil {
		log.Logger.Error("databases error", err)
	}
	database.Log = log.Logger
	grpcService := services.Banner{Database: database, Log: log.Logger}
	return &rest.RestBannerServer{BannerService: &grpcService, Log: log.Logger}, nil
}

func RunRestServer(c chan bool) {
	log.Logger.Info("run rest api server")
	restServer, err := initRestServer()
	if err != nil {
		log.Logger.Error("error run rest server ", err)
	}
	address := viper.GetString("API_REST_HOST") + ":" + viper.GetString("API_REST_PORT")
	log.Logger.Info(address)
	restServer.RunServer(address)
	c <- true
}

var RestApiServerCmd = &cobra.Command{
	Use:   "rest_api_server",
	Short: "run rest api server",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan bool)
		RunRestServer(c)
	},
}

func init() {
	err := viper.BindEnv("DB_DSN")
	err = viper.BindEnv("API_REST_PORT")
	err = viper.BindEnv("API_REST_HOST")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
