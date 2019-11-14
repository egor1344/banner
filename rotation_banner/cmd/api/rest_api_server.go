/*
	Запуск rest api сервиса
*/

package api

import (
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbDsn string

var RestApiServerCmd = &cobra.Command{
	Use:   "rest_api_server",
	Short: "run rest api server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Info("run rest api server mazafaka")
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
