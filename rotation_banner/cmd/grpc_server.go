package cmd

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
)

// var dbDsn string

var GrpcServerCmd = &cobra.Command{
	Use: "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string){
		log.Logger.Info("run grpc server mazafaka")
	},
}

func init () {
	err := viper.BindEnv("DB_DSN")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}