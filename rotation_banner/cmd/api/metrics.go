/*
	Запуск grpc сервиса
*/

package api

import (
	"net/http"

	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunMetricsHandler(c chan bool) {
	address := viper.GetString("METRICS_HOST") + ":" + viper.GetString("METRICS_PORT")
	log.Logger.Info(address)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(address, nil)
	if err != nil {
		c <- true
		log.Logger.Fatal("Error run server")
	}
}

var MetricsCMD = &cobra.Command{
	Use:   "metrics",
	Short: "run metrics",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan bool)
		RunMetricsHandler(c)
	},
}

func init() {
	err := viper.BindEnv("METRICS_PORT")
	err = viper.BindEnv("METRICS_HOST")
	if err != nil {
		log.Logger.Info(err)
	}
	viper.AutomaticEnv()
	log.Logger.Info(viper.AllSettings())
}
