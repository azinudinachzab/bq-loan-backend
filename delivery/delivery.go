package delivery

import (
	"net/http"
	"time"

	"github.com/azinudinachzab/bq-loan-be-v2/delivery/cron"
	httpDelivery "github.com/azinudinachzab/bq-loan-be-v2/delivery/http"
	"github.com/azinudinachzab/bq-loan-be-v2/service"
	"github.com/go-co-op/gocron"
)

type (
	Dependency struct {
		Service  service.Service
		Timezone *time.Location
	}

	Delivery struct {
		HttpServer http.Handler
		Cron       *gocron.Scheduler
	}
)

func NewDelivery(dep Dependency) *Delivery {
	return &Delivery{
		HttpServer: httpDelivery.NewHttpServer(dep.Service),
		Cron:       cron.NewCron(dep.Service, dep.Timezone),
	}
}
