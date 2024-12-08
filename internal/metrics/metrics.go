package metrics

import (
	"github.com/zsais/go-gin-prometheus"
)

var (
	M *ginprometheus.Prometheus
)

func init() {
	M = ginprometheus.NewPrometheus("search_nova")
}
