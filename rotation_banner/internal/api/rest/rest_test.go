package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"

	"github.com/egor1344/banner/rotation_banner/internal/domain/models"

	"github.com/egor1344/banner/rotation_banner/internal/domain/services"

	"github.com/egor1344/banner/rotation_banner/tests"

	"github.com/egor1344/banner/rotation_banner/internal/databases/postgres"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"github.com/spf13/viper"
)

var restService *RestBannerServer
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
	restService = &RestBannerServer{Log: log.Logger, BannerService: bannerService}
	os.Exit(t.Run())
}

func TestRestBannerServer_AddBannerHandler(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	restService.Log.Info("Проверка функционала")
	body := []byte(`{ "id_banner": 1, "id_slot": 1}`)
	req, err := http.NewRequest("POST", "/api/add_banner/", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restService.AddBannerHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("wrong answer ")
	}
}

func TestRestBannerServer_DelBannerHandler(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	_, err := pgbs.DB.Exec(`
		INSERT INTO public.banners (id) VALUES (1);
		INSERT INTO public.banners (id) VALUES (2);
		INSERT INTO public.slot (id) VALUES (1);
		INSERT INTO public.slot (id) VALUES (2);
		INSERT INTO public.rotations (id, id_slot, id_banner) VALUES (1,1,1);
		INSERT INTO public.rotations (id, id_slot, id_banner) VALUES (2,2,2);
		`)
	restService.Log.Info("Проверка функционала")
	req, err := http.NewRequest("DELETE", "/api/del_banner/1/", nil)
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		t.Error(err)
	}
	t.Log(mux.Vars(req))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restService.DelBannerHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("wrong answer ")
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

func TestRestBannerServer_CountTransitionHandler(t *testing.T) {
	tests.TruncateDb(t, pgbs.DB)
	_, err := pgbs.DB.Exec(`
		INSERT INTO public.banners (id) VALUES (1);
		INSERT INTO public.slot (id) VALUES (1);
`)
	restService.Log.Info("Проверка функционала")
	body := []byte(`{"id_banner": 1,"id_slot": 1,"id_soc_dem": 1}`)
	req, err := http.NewRequest("POST", "/api/count_transition/", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restService.CountTransitionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("wrong answer ")
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

func TestRestBannerServer_GetBannerHandler(t *testing.T) {
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
	restService.Log.Info("Проверка функционала")
	req, err := http.NewRequest("GET", "/api/get_banner/1/1/", nil)
	vars := map[string]string{"idSlot": "1", "idSocDemGroup": "1"}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		t.Error(err)
	}
	t.Log(mux.Vars(req))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restService.GetBannerHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("wrong answer ")
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
