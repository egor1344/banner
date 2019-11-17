package services

import (
	"context"
	"testing"

	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/viper"
)

var bannerService *Banner

func TestMain(t *testing.M) {
	err := viper.BindEnv("DB_DSN")
	if err != nil {
		return
	}
	viper.AutomaticEnv()
	dbDsn := viper.GetString("DB_DSN")
	if dbDsn == "" {
		log.Logger.Error("dont set env variable DB_DSN")
	}
	database, err := postgres.InitPgBannerStorage(dbDsn)
	if err != nil {
		log.Logger.Error("databases error", err)
	}
	database.Log = log.Logger
	bannerService = &Banner{Database: database, Log: log.Logger}
}

func TestBanner_AddBanner(t *testing.T) {
	err := bannerService.AddBanner(context.Background(), 1, 1)
	if err != nil {
		t.Error(err)
	}
}
