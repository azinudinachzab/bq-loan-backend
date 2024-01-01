package cron

import (
	"context"
	"fmt"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
	"time"

	"github.com/azinudinachzab/bq-loan-be-v2/service"
	"github.com/go-co-op/gocron"
)

type Cron struct {
	service service.Service
}

func NewCron(svc service.Service, loc *time.Location) *gocron.Scheduler {
	s := gocron.NewScheduler(loc)
	c := &Cron{
		service: svc,
	}

	s.WaitForScheduleAll()
	if _, err := s.Every(1).Hour().Do(c.CronHealthCheck); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "healthcheck")
	}
	now := time.Now()
	specificTime := time.Date(now.Year(), now.Month(), 25, 0, 0, 0, 0, now.Location())
	if _, err := s.Every(1).Month(25).StartAt(specificTime).Do(c.MonthlyMandatory); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "iuranwajib")
	}

	return s
}

//func (c *Cron) FiveSecLogger() {
//	log.Println("5 sec log")
//}

func (c *Cron) MonthlyMandatory() {
	ctx := context.Background()
	now := time.Now()
	execTime := now.Format(time.DateTime)
	log.Printf("Executing job for date: %s\n", execTime)

	flt := make(map[string]string)
	usrs, err := c.service.GetUsers(ctx, flt)
	if err != nil {
		log.Printf("error when get users %v\n", err)
		return
	}

	for _, val := range usrs {
		if err := c.service.CreateLoanGeneral(ctx, model.LoanGeneral{
			UserID:     val.ID,
			Title:      fmt.Sprintf("iuran wajib %s", now.Month().String()),
			Amount:     25000,
			Datetime:   execTime,
			Tenor:      1,
			Status:     0,
			LoanTypeID: 1,
		}); err != nil {
			log.Printf("error when store loan %v\n to user %v", err, val.ID)
			return
		}
	}
	log.Println("Job done for date: ", execTime)
}

func (c *Cron) CronHealthCheck() {
	execTime := time.Now().Format(time.DateTime)
	log.Printf("health check done for date: %s\n", execTime)
}
