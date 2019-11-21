package services

import (
	"context"
	"os"
	"testing"

	"github.com/egor1344/banner/rotation_banner/internal/domain/models"

	"github.com/egor1344/banner/rotation_banner/tests"

	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/viper"
)

var bannerService *Banner
var pgbs *postgres.PgBannerStorage

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
	pgbs = database
	if err != nil {
		log.Logger.Error("databases error", err)
	}
	database.Log = log.Logger
	bannerService = &Banner{Database: database, Log: log.Logger}
	os.Exit(t.Run())
}

func TestBanner_AddBanner(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	bannerService.Log.Info("Проверка функционала")
	err := bannerService.AddBanner(context.Background(), 1, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Проверяем наличие данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM rotations")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{1, 1},
	}
	for _, c := range testCases {
		rows.Next()
		var rotation models.Rotation
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		if rotation.IdSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IdSlot, " must be = ", c.idSlot)
		}
		if rotation.IdBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IdBanner, " must be = ", c.idBanner)
		}
	}
}
