package grpc

import (
	"context"
	"os"
	"testing"

	"github.com/egor1344/banner/rotation_banner/internal/domain/models"

	"github.com/egor1344/banner/rotation_banner/proto/slot"

	"github.com/egor1344/banner/rotation_banner/proto/server"

	"github.com/egor1344/banner/rotation_banner/internal/domain/services"

	"github.com/egor1344/banner/rotation_banner/tests"

	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/egor1344/banner/rotation_banner/proto/banner"
	"github.com/spf13/viper"
)

var grpcService *GrpcBannerServer
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
	bannerService := &services.Banner{Database: database, Log: log.Logger}
	grpcService = &GrpcBannerServer{Log: log.Logger, BannerService: bannerService}
	os.Exit(t.Run())
}

func TestGrpcBannerServer_AddBanner(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	grpcService.Log.Info("Проверка функционала")
	testAddBannerRequest := server.AddBannerRequest{Banner: &banner.Banner{Id: 1, Slot: &slot.Slot{Id: 1}, Description: "Test banner"}}
	resp, err := grpcService.AddBanner(context.Background(), &testAddBannerRequest)
	if err != nil {
		t.Error(err)
	}
	if resp.GetError() != "" {
		t.Error("grpc error ", resp.GetError())
	}
}

func TestGrpcBannerServer_DelBanner(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	_, err := pgbs.DB.Exec(`
		INSERT INTO public.banners (id) VALUES (1);
		INSERT INTO public.banners (id) VALUES (2);
		INSERT INTO public.slot (id) VALUES (1);
		INSERT INTO public.slot (id) VALUES (2);
		INSERT INTO public.rotations (id, id_slot, id_banner) VALUES (1,1,1);
		INSERT INTO public.rotations (id, id_slot, id_banner) VALUES (2,2,2);
		`)
	grpcService.Log.Info("Проверка функционала")
	testDelBannerRequest := server.DelBannerRequest{Id: 1}
	resp, err := grpcService.DelBanner(context.Background(), &testDelBannerRequest)
	if err != nil {
		t.Error(err)
	}
	if resp.GetError() != "" {
		t.Error("grpc error ", resp.GetError())
	}
	t.Log("Проверяем измение данных в БД")
	rows, err := pgbs.DB.Queryx("SELECT * FROM rotations")
	if err != nil {
		t.Error("connect database error ", err)
	}
	testCases := []struct {
		idBanner, idSlot int64
	}{
		{2, 2},
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

func TestGrpcBannerServer_CountTransition(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	_, err := pgbs.DB.Exec(`
		INSERT INTO public.banners (id) VALUES (1);
		INSERT INTO public.slot (id) VALUES (1);
`)
	grpcService.Log.Info("Проверка функционала")
	testCountTransitionRequest := server.CountTransitionRequest{IdBanner: 1, IdSocDemGroup: 1, IdSlot: 1}
	resp, err := grpcService.CountTransition(context.Background(), &testCountTransitionRequest)
	if err != nil {
		t.Error(err)
	}
	if resp.GetError() != "" {
		t.Error("grpc error ", resp.GetError())
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
		if rotation.IdSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IdSlot, " must be = ", c.idSlot)
		}
		if rotation.IdBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IdBanner, " must be = ", c.idBanner)
		}
		if rotation.IdSocDemGroup != c.idSocDemGroup {
			t.Error("rotation id soc dem group = ", rotation.IdSocDemGroup, " must be = ", c.idSocDemGroup)
		}
		if rotation.CountClick != c.countClick {
			t.Error("rotation count click = ", rotation.CountClick, " must be = ", c.countClick)
		}
		if rotation.CountViews != c.CountViews {
			t.Error("rotation count views = ", rotation.CountViews, " must be = ", c.CountViews)
		}
	}
}

func TestGrpcBannerServer_GetBanner(t *testing.T) {
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
	grpcService.Log.Info("Проверка функционала")
	testGetBannerRequest := server.GetBannerRequest{IdSocDemGroup: 1, IdSlot: 1}
	resp, err := grpcService.GetBanner(context.Background(), &testGetBannerRequest)
	if err != nil {
		t.Error(err)
	}
	if resp.GetError() != "" {
		t.Error("grpc error ", resp.GetError())
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
		if rotation.IdSlot != c.idSlot {
			t.Error("rotation id slot = ", rotation.IdSlot, " must be = ", c.idSlot)
		}
		if rotation.IdBanner != c.idBanner {
			t.Error("rotation id banner = ", rotation.IdBanner, " must be = ", c.idBanner)
		}
		if rotation.IdSocDemGroup != c.idSocDemGroup {
			t.Error("rotation id soc dem group = ", rotation.IdSocDemGroup, " must be = ", c.idSocDemGroup)
		}
		if rotation.CountClick != c.countClick {
			t.Error("rotation count click = ", rotation.CountClick, " must be = ", c.countClick)
		}
		if rotation.CountViews != c.CountViews {
			t.Error("rotation count view = ", rotation.CountViews, " must be = ", c.CountViews)
		}
	}
}
