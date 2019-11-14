package api

import (
	"github.com/spf13/cobra"
)

// RootCmd - cobra command
var RootCmd = &cobra.Command{
	Use:   "banner",
	Short: "Banner service",
}

func init() {
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(RestApiServerCmd)
}
