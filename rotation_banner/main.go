package main

import (
	"log"

	"github.com/egor1344/banner/rotation_banner/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
