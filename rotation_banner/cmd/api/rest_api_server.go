/*
	Запуск rest api сервиса
*/

package api

import (
	"github.com/egor1344/banner/rotation_banner/internal/amqp"
	"github.com/egor1344/banner/rotation_banner/internal/api/rest"
	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	"github.com/egor1344/banner/rotation_banner/internal/domain/services"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initRestServer - инициализация grpc сервера
func initRestServer() (*rest.BannerServerRest, error) {
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
	// Инициализация очереди событий
	rabbit := &amqp.Rabbit{AMQPDSN: viper.GetString("AMQP_DSN"), QueueName: viper.GetString("QUEUE_NAME")}
	rabbit.Log = log.Logger
	err = rabbit.Init()
	if err != nil {
		log.Logger.Error("amqp error ", err)
	}
	grpcService := services.Banner{Database: database, Log: log.Logger, AMQP: rabbit}
	return &rest.BannerServerRest{BannerService: &grpcService, Log: log.Logger}, nil
}

// RunRestServer - инициализация rest сервера
func RunRestServer(c chan bool) {
	log.Logger.Info("run rest api server")
	restServer, err := initRestServer()
	if restServer == nil {
		c <- true
		log.Logger.Fatal("Error run server")
		return
	}
	defer restServer.BannerService.CloseConnection()
	if err != nil {
		log.Logger.Error("error run rest server ", err)
	}
	address := viper.GetString("API_REST_HOST") + ":" + viper.GetString("API_REST_PORT")
	log.Logger.Info(address)
	restServer.RunServer(address)
	c <- true
}

// RestAPIServerCmd - Комманда для cobra
var RestAPIServerCmd = &cobra.Command{
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
	err = viper.BindEnv("AMQP_DSN")
	err = viper.BindEnv("QUEUE_NAME")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
