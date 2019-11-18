package postgres

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/viper"
)

var dbDsn string
var pgbs *PgBannerStorage

func TestMain(t *testing.M) {
	err := viper.BindEnv("DB_DSN")
	if err != nil {
		log.Fatal(err)
	}
	dbDsn = viper.GetString("DB_DSN")
	log.Println(dbDsn)
	pgbs, err = InitPgBannerStorage(dbDsn)
	if pgbs == nil {
		log.Fatal(err)
	}
	pgbs.Log = logger.Logger
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(t.Run())
}

type BannerBD struct {
	id int64 `db:"id"`
}

func TestPgBannerStorage_AddBanner(t *testing.T) {
	// Проверяем функционал
	ctx := context.Background()
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
	}
	for i, c := range testCases {
		t.Log("add banner case #", i)
		err := pgbs.AddBanner(ctx, c.idBanner, c.idSlot)
		if err != nil {
			t.Error(err)
		}
	}
	// Проверяем наличие данных в БД
	//rows, err := pgbs.db.Queryx("SELECT * FROM banners")
	//if err != nil {
	//	t.Error("connect database error ", err)
	//}
	//for i, c := range testCases {
	//	rows.Next()
	//	var banner BannerBD
	//	err = rows.StructScan(banner)
	//}
}

func TestPgBannerStorage_DelBanner(t *testing.T) {
	ctx := context.Background()
	err := pgbs.DelBanner(ctx, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestPgBannerStorage_CountTransition(t *testing.T) {
	ctx := context.Background()
	err := pgbs.CountTransition(ctx, 1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestPgBannerStorage_GetBanner(t *testing.T) {
	ctx := context.Background()
	err := pgbs.CountTransition(ctx, 1, 1)
	if err != nil {
		t.Error(err)
	}
}
