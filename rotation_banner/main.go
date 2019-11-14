package main

import (
	"github.com/egor1344/banner/rotation_banner/cmd/api"
	"log"
)

func main() {
	if err := api.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
