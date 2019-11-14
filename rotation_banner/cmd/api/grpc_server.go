/*
	Запуск grpc сервиса
*/

package api

import (
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var dbDsn string

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Info("run grpc server mazafaka")
	},
}

func init() {
	err := viper.BindEnv("DB_DSN")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
