package main

import (
	"github.com/egor1344/banner/rotation_banner/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}