package cmd

import (
	"github.com/egor1344/banner/rotation_banner/cmd/api"
	"github.com/spf13/cobra"
)

// RootCmd - cobra command
var RootCmd = &cobra.Command{
	Use:   "banner",
	Short: "Banner service",
	Run: func(cmd *cobra.Command, args []string) {
		rest := make(chan bool)
		grpc := make(chan bool)
		metrics := make(chan bool)
		go api.RunGrpcServer(grpc)
		go api.RunRestServer(rest)
		go RunMetricsHandler(rest)
		for {
			_, ok1 := <-rest
			_, ok2 := <-grpc
			_, ok3 := <-metrics
			if ok1 == false && ok2 == false && ok3 == false {
				break
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(api.GrpcServerCmd)
	RootCmd.AddCommand(api.RestApiServerCmd)
	RootCmd.AddCommand(MetricsCMD)
}
