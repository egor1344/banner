package api

import (
	"log"
	"net/http"
	"time"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

// RestBannerServer - Реализует работу с rest сервером.
type RestBannerServer struct {
	BannerService interfaces.Service
	Log           *zap.SugaredLogger
}

// AddBannerHandler - Добавить баннер
func (rbs *RestBannerServer) AddBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest add banner")
}

// DelBannerHandler - Удалить баннер
func (rbs *RestBannerServer) DelBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest del banner")
}

// CountTransitionHandler - Засчитать переход
func (rbs *RestBannerServer) CountTransitionHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest count transition")
}

// GetBannerHandler - Выбрать баннер для показа
func (rbs *RestBannerServer) GetBannerHandler(w http.ResponseWriter, r *http.Request) {
	rbs.Log.Info("rest get banner")
}

func (rbs *RestBannerServer) RunServer(address string) {
	rbs.Log.Info("run rest server")
	router := mux.NewRouter()
	router.HandleFunc("/api/add_banner/", rbs.AddBannerHandler).Methods("POST")
	router.HandleFunc("/api/del_banner/", rbs.DelBannerHandler).Methods("POST")
	router.HandleFunc("/api/count_transition/", rbs.CountTransitionHandler).Methods("POST")
	router.HandleFunc("/api/get_banner/", rbs.GetBannerHandler).Methods("GET")
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
