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

type rotationBD struct {
	Id       int64 `db:"id"`
	IdBanner int64 `db:"id_banner"`
	IdSlot   int64 `db:"id_slot"`
}

func truncate_db(t *testing.T, pgbs *PgBannerStorage) {
	tables := []string{"banners", "rotations", "slot", "soc_dem_group", "statistic"}
	for _, table := range tables {
		_, err := pgbs.db.Query("truncate table " + table + " restart identity cascade;")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestPgBannerStorage_AddBanner(t *testing.T) {
	// truncate_db
	truncate_db(t, pgbs)
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
	//Проверяем наличие данных в БД
	rows, err := pgbs.db.Queryx("SELECT * FROM rotations")
	if err != nil {
		t.Error("connect database error ", err)
	}
	for _, c := range testCases {
		rows.Next()
		var rotation rotationBD
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rotation)
		if rotation.IdSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IdSlot, " must be = ", c.idSlot)
		}
		if rotation.IdBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IdBanner, " must be = ", c.idBanner)
		}
	}
}

//
//func TestPgBannerStorage_DelBanner(t *testing.T) {
//	ctx := context.Background()
//	err := pgbs.DelBanner(ctx, 1)
//	if err != nil {
//		t.Error(err)
//	}
//}
//
//func TestPgBannerStorage_CountTransition(t *testing.T) {
//	ctx := context.Background()
//	err := pgbs.CountTransition(ctx, 1, 1)
//	if err != nil {
//		t.Error(err)
//	}
//}
//
//func TestPgBannerStorage_GetBanner(t *testing.T) {
//	ctx := context.Background()
//	err := pgbs.CountTransition(ctx, 1, 1)
//	if err != nil {
//		t.Error(err)
//	}
//}
