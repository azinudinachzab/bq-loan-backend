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
	if _, err := s.Every(1).Day().Do(c.MonthlyMandatory); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "iuranwajib")
	}
	if _, err := s.Every(1).Day().Do(c.MonthlyBQMart); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "bq mart")
	}
	if _, err := s.Every(1).Day().Do(c.MonthlySocialFund); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "dana sosial")
	}
	if _, err := s.Every(1).Day().Do(c.DetailPaid); err != nil {
		log.Printf("error when running cron %v\n job name: %v\n", err, "change detail to paid")
	}

	return s
}

//func (c *Cron) FiveSecLogger() {
//	log.Println("5 sec log")
//}

func (c *Cron) MonthlyMandatory() {
	ctx := context.Background()
	now := time.Now()

	if now.Day() != 20 {
		return
	}

	execTime := now.Format(time.RFC3339)
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
			Amount:     200000,
			Datetime:   execTime,
			Tenor:      1,
			Status:     1,
			LoanTypeID: 1,
		}); err != nil {
			log.Printf("error when store loan %v\n to user %v", err, val.ID)
			return
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("Job done for date: ", execTime)
}

func (c *Cron) MonthlySocialFund() {
	ctx := context.Background()
	now := time.Now()

	if now.Day() != 20 {
		return
	}

	execTime := now.Format(time.RFC3339)
	log.Printf("Executing job for date: %s\n", execTime)

	flt := make(map[string]string)
	usrs, err := c.service.GetUsers(ctx, flt)
	if err != nil {
		log.Printf("error when get users %v\n", err)
		return
	}

	for _, val := range usrs {
		if val.IsLeader != 1 {
			continue
		}
		if err := c.service.CreateLoanGeneral(ctx, model.LoanGeneral{
			UserID:     val.ID,
			Title:      fmt.Sprintf("dana sosial %s", now.Month().String()),
			Amount:     25000,
			Datetime:   execTime,
			Tenor:      1,
			Status:     1,
			LoanTypeID: 5,
		}); err != nil {
			log.Printf("error when store loan %v\n to user %v", err, val.ID)
			return
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("Job done for date: ", execTime)
}

func (c *Cron) MonthlyBQMart() {
	ctx := context.Background()
	now := time.Now()

	if now.Day() != 20 {
		return
	}

	execTime := now.Format(time.RFC3339)
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
			Status:     1,
			LoanTypeID: 7,
		}); err != nil {
			log.Printf("error when store loan %v\n to user %v", err, val.ID)
			return
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("Job done for date: ", execTime)
}

func (c *Cron) DetailPaid() {
	ctx := context.Background()
	now := time.Now()

	if now.Day() != 20 {
		return
	}

	execTime := now.Format(time.RFC3339)
	log.Printf("Executing job for date: %s\n", execTime)

	details, err := c.service.GetMonthlyLoanDetails(ctx, int(now.Month()))
	if err != nil {
		log.Printf("error when get users %v\n", err)
		return
	}

	for _, val := range details {
		if err := c.service.AcceptPaymentRequest(ctx, val.ID); err != nil {
			log.Printf("error when pay loan %v\n to loan id %v", err, val.ID)
			return
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("Job done for date: ", execTime)
}

func (c *Cron) CronHealthCheck() {
	execTime := time.Now().Format(time.DateTime)
	log.Printf("health check done for date: %s\n", execTime)
}
