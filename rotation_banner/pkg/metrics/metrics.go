package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// APICounter - количество обращений к API
	APICounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "Client user api",
	})
	// AddBannerCounter - количество обращений к методу API AddBanner
	AddBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_count",
		Help: "Client user api| uses method add banner",
	})
	// AddBannerGrpcCounter - количество обращений к методу API AddBanner с помощью grpc
	AddBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	// AddBannerRestCounter - количество обращений к методу API AddBanner с помощью rest
	AddBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	// DelBannerCounter - количество обращений к методу API DelBanner
	DelBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_count",
		Help: "Client user api| uses method add banner",
	})
	// DelBannerGrpcCounter - количество обращений к методу API DelBanner с помощью grpc
	DelBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	// DelBannerRestCounter - количество обращений к методу API DelBanner с помощью rest
	DelBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "del_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	// CountTransitionCounter - количество обращений к методу API CountTransition
	CountTransitionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_count",
		Help: "Client user api| uses method add banner",
	})
	// CountTransitionGrpcCounter - количество обращений к методу API CountTransition с помощью grpc
	CountTransitionGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	// CountTransitionRestCounter - количество обращений к методу API CountTransition с помощью rest
	CountTransitionRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "count_transition_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
	// GetBannerCounter - количество обращений к методу API GetBanner
	GetBannerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_count",
		Help: "Client user api| uses method add banner",
	})
	// GetBannerGrpcCounter - количество обращений к методу API GetBanner  с помощью grpc
	GetBannerGrpcCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_grpc_count",
		Help: "Client user api| uses grpc method add banner",
	})
	// GetBannerRestCounter - количество обращений к методу API GetBanner  с помощью rest
	GetBannerRestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_banner_rest_count",
		Help: "Client user api| uses rest method add banner",
	})
)
