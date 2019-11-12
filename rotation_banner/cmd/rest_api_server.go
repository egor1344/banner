package cmd

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
)

var dbDsn string

var RestApiServerCmd = &cobra.Command{
	Use: "rest_api_server",
	Short: "run rest api server",
	Run: func(cmd *cobra.Command, args []string){
		log.Logger.Info("run rest api server mazafaka")
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