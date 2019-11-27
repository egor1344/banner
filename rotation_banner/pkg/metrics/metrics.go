package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ApiCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "Client user api",
	})
	AddBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_count",
		Help: "Client user api| uses method add banner",
	})
	AddBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	AddBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	DelBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_count",
		Help: "Client user api| uses method add banner",
	})
	DelBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	DelBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	CountTransitionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_count",
		Help: "Client user api| uses method add banner",
	})
	CountTransitionGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	CountTransitionRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	GetBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_count",
		Help: "Client user api| uses method add banner",
	})
	GetBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	GetBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
)
