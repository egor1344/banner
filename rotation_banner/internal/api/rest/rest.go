package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/egor1344/banner/rotation_banner/pkg/metrics"

	"github.com/egor1344/banner/rotation_banner/internal/domain/models"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

// BannerServerRest - Реализует работу с rest сервером.
type BannerServerRest struct {
	BannerService interfaces.Service
	Log           *zap.SugaredLogger
}

// AddBannerHandler - Добавить баннер
func (rbs *BannerServerRest) AddBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest add banner")
	metrics.APICounter.Inc()
	metrics.AddBannerRestCounter.Inc()
	var rotations models.Rotation
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, err = fmt.Fprint(w, "Error enter data")
		return
	}
	rbs.Log.Info(string(reqBody))
	err = json.Unmarshal(reqBody, &rotations)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	rbs.Log.Info(rotations)
	err = rbs.BannerService.AddBanner(r.Context(), rotations.IDBanner, rotations.IDSlot)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	_, err = fmt.Fprint(w, `{"status": true}`)
}

// DelBannerHandler - Удалить баннер
func (rbs *BannerServerRest) DelBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest del banner ", mux.Vars(r))
	metrics.APICounter.Inc()
	metrics.DelBannerRestCounter.Inc()
	bannerID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		rbs.Log.Error(err)
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	rbs.Log.Info(bannerID)
	err = rbs.BannerService.DelBanner(r.Context(), bannerID)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	_, err = fmt.Fprint(w, `{"status": true}`)
}

// CountTransitionHandler - Засчитать переход
func (rbs *BannerServerRest) CountTransitionHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest count transition")
	metrics.APICounter.Inc()
	metrics.CountTransitionRestCounter.Inc()
	var statistic models.Statistic
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, err = fmt.Fprint(w, "Error enter data")
		return
	}
	rbs.Log.Info(string(reqBody))
	err = json.Unmarshal(reqBody, &statistic)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	err = rbs.BannerService.CountTransition(r.Context(), statistic.IDBanner, statistic.IDSocDemGroup, statistic.IDSlot)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	_, err = fmt.Fprint(w, `{"status": true}`)
}

// GetBannerHandler - Выбрать баннер для показа
func (rbs *BannerServerRest) GetBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest get banner")
	metrics.APICounter.Inc()
	metrics.GetBannerRestCounter.Inc()
	idSlot, err := strconv.ParseInt(mux.Vars(r)["idSlot"], 10, 64)
	idSocDemGroup, err := strconv.ParseInt(mux.Vars(r)["idSocDemGroup"], 10, 64)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	rbs.Log.Info(idSlot, idSocDemGroup)
	idBanner, err := rbs.BannerService.GetBanner(r.Context(), idSlot, idSocDemGroup)
	if err != nil {
		w.WriteHeader(500)
		_, err = fmt.Fprint(w, "Error ", err)
		return
	}
	rbs.Log.Info(idBanner)
	output := `{"status": true,"id_banner": ` + strconv.FormatInt(idBanner, 10) + `}`
	_, err = fmt.Fprint(w, output)
}

//RunServer - запуск сервера
func (rbs *BannerServerRest) RunServer(address string) {
	rbs.Log.Info("run rest server")
	router := mux.NewRouter()
	router.HandleFunc("/api/add_banner/", rbs.AddBannerHandler).Methods("POST")
	router.HandleFunc("/api/del_banner/{id}/", rbs.DelBannerHandler).Methods("DELETE")
	router.HandleFunc("/api/count_transition/", rbs.CountTransitionHandler).Methods("POST")
	router.HandleFunc("/api/get_banner/{idSlot}/{idSocDemGroup}/", rbs.GetBannerHandler).Methods("GET")
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
