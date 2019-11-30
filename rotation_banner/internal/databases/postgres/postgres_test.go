package postgres

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/egor1344/banner/rotation_banner/tests"

	"github.com/egor1344/banner/rotation_banner/internal/domain/models"
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
		return
	}
	pgbs.Log = logger.Logger
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(t.Run())
}

// createTestData - Создание и заполнение тестовых данных
func createTestData(ctx *context.Context, t *testing.T, pgbs *PgBannerStorage) {
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
	}
	for i, c := range testCases {
		t.Log("add banner case #", i)
		err := pgbs.AddBanner(*ctx, c.idBanner, c.idSlot)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPgBannerStorage_AddBanner(t *testing.T) {
	t.Log("truncateDb")
	tests.TruncateDb(t, pgbs.DB)
	t.Log("Проверяем функционал")
	ctx := context.Background()
	createTestData(&ctx, t, pgbs)
	t.Log("Проверяем наличие данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM rotations")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
	}
	for _, c := range testCases {
		rows.Next()
		var rotation models.Rotation
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		if rotation.IDSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IDSlot, " must be = ", c.idSlot)
		}
		if rotation.IDBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IDBanner, " must be = ", c.idBanner)
		}
	}
}

func TestPgBannerStorage_DelBanner(t *testing.T) {
	t.Log("truncateDb")
	tests.TruncateDb(t, pgbs.DB)
	ctx := context.Background()
	t.Log("Заполняем тестовыми данными")
	createTestData(&ctx, t, pgbs)
	t.Log("Тестирование функционала")
	err := pgbs.DelBanner(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Проверяем измение данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM rotations")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{2, 1},
		{3, 2},
	}
	for _, c := range testCases {
		rows.Next()
		var rotation models.Rotation
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		if rotation.IDSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IDSlot, " must be = ", c.idSlot)
		}
		if rotation.IDBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IDBanner, " must be = ", c.idBanner)
		}
	}
}

func TestPgBannerStorage_CountTransition(t *testing.T) {
	t.Log("truncateDb")
	tests.TruncateDb(t, pgbs.DB)
	ctx := context.Background()
	t.Log("Заполняем тестовыми данными")
	createTestData(&ctx, t, pgbs)
	t.Log("Тестирование функционала")
	err := pgbs.CountTransition(ctx, 1, 1, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Проверяем измение данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM statistic")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot, idSocDemGroup, countClick, CountViews int64
	}{
		{1, 1, 1, 1, 2},
	}
	for _, c := range testCases {
		rows.Next()
		var rotation models.Statistic
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		if rotation.IDSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IDSlot, " must be = ", c.idSlot)
		}
		if rotation.IDBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IDBanner, " must be = ", c.idBanner)
		}
		if rotation.IDSocDemGroup != c.idSocDemGroup {
			t.Error("rotation id soc dem group = ", rotation.IDSocDemGroup, " must be = ", c.idSocDemGroup)
		}
		if rotation.CountClick != c.countClick {
			t.Error("rotation count click = ", rotation.CountClick, " must be = ", c.countClick)
		}
		if rotation.CountViews != c.CountViews {
			t.Error("rotation count views = ", rotation.CountViews, " must be = ", c.CountViews)
		}
	}
}

func TestPgBannerStorage_GetBanner(t *testing.T) {
	t.Log("truncateDb")
	tests.TruncateDb(t, pgbs.DB)
	_, err := pgbs.DB.Exec(`
		INSERT INTO public.banners (id) VALUES (1);
		INSERT INTO public.banners (id) VALUES (2);
		INSERT INTO public.banners (id) VALUES (3);
		INSERT INTO public.slot (id) VALUES (1);
		INSERT INTO public.soc_dem_group (id) VALUES (1);
		INSERT INTO public.statistic (id, id_banner, id_soc_dem, count_click, count_views, id_slot) VALUES (1, 2, 1, 1, 2, 1);
		INSERT INTO public.statistic (id, id_banner, id_soc_dem, count_click, count_views, id_slot) VALUES (2, 1, 1, 2, 4, 1);
		INSERT INTO public.statistic (id, id_banner, id_soc_dem, count_click, count_views, id_slot) VALUES (3, 3, 1, 4, 6, 1);`)

	ctx := context.Background()
	_, err = pgbs.GetBanner(ctx, 1, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Проверяем измение данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM statistic")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot, idSocDemGroup, countClick, CountViews int64
	}{
		{1, 1, 1, 2, 4},
		{3, 1, 1, 4, 6},
		{2, 1, 1, 1, 3}, // здесь увеличение countViews на 1
	}
	for _, c := range testCases {
		rows.Next()
		var rotation models.Statistic
		err = rows.StructScan(&rotation)
		if err != nil {
			t.Fatal(err)
		}
		if rotation.IDSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IDSlot, " must be = ", c.idSlot)
		}
		if rotation.IDBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IDBanner, " must be = ", c.idBanner)
		}
		if rotation.IDSocDemGroup != c.idSocDemGroup {
			t.Error("rotation id soc dem group = ", rotation.IDSocDemGroup, " must be = ", c.idSocDemGroup)
		}
		if rotation.CountClick != c.countClick {
			t.Error("rotation count click = ", rotation.CountClick, " must be = ", c.countClick)
		}
		if rotation.CountViews != c.CountViews {
			t.Error("rotation count view = ", rotation.CountViews, " must be = ", c.CountViews)
		}
	}
}
