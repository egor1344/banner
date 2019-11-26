package api

import (
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
		go RunGrpcServer(grpc)
		go RunRestServer(rest)
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
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(RestApiServerCmd)
	RootCmd.AddCommand(MetricsCMD)
}
