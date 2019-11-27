/*
	Запуск grpc сервиса
*/

package api

import (
	"github.com/egor1344/banner/rotation_banner/internal/amqp"
	"github.com/egor1344/banner/rotation_banner/internal/api/grpc"
	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	"github.com/egor1344/banner/rotation_banner/internal/domain/services"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initGrpcServer - инициализация grpc сервера
func initGrpcServer() (*grpc.GrpcBannerServer, error) {
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
	// Инициализация очереди событий
	rabbit := &amqp.Rabbit{AMQPDSN: viper.GetString("AMQP_DSN"), QueueName: viper.GetString("QUEUE_NAME")}
	rabbit.Log = log.Logger
	err = rabbit.Init()
	if err != nil {
		log.Logger.Error("amqp error ", err)
	}
	grpcService := services.Banner{Database: database, Log: log.Logger, AMQP: rabbit}
	return &grpc.GrpcBannerServer{BannerService: &grpcService, Log: log.Logger}, nil
}

func RunGrpcServer(c chan bool) {
	log.Logger.Info("run grpc server")
	grpcServer, err := initGrpcServer()
	defer grpcServer.BannerService.CloseConnection()
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Info(grpcServer)
	address := viper.GetString("API_GRPC_HOST") + ":" + viper.GetString("API_GRPC_PORT")
	log.Logger.Info(address)
	err = grpcServer.RunServer("tcp", address)
	if err != nil {
		c <- true
		log.Logger.Fatal("Error run server")
	}
}

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan bool)
		RunGrpcServer(c)
	},
}

func init() {
	err := viper.BindEnv("DB_DSN")
	err = viper.BindEnv("API_GRPC_PORT")
	err = viper.BindEnv("API_GRPC_HOST")
	err = viper.BindEnv("AMQP_DSN")
	err = viper.BindEnv("QUEUE_NAME")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
