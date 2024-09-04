package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

type server struct{}

func (s *server) StartSchedulerJob() error {
	c := cron.New()

	_, err := c.AddFunc("* * * * *", func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		fmt.Println("Rest job executed")
	})
	if err != nil {
		log.Printf("Error adding cron job: %v", err)
	}
	c.Start()

	if err != nil {
		return err
	}

	return nil
}

func (s *server) startScheduledJobHandler(c echo.Context) error {
	if err := s.StartSchedulerJob(); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("failed to start scheduled job: %v", err))
	}
	return c.String(http.StatusOK, "Scheduled job started successfully")
}

func main() {
	c := cron.New()

	_, err := c.AddFunc("@every 1m", func() {
		fmt.Println("Cron job executed!")
	})
	if err != nil {
		log.Fatalf("Error adding cron job 1: %v", err)
	}

	_, err = c.AddFunc("* * * * *", func() {
		fmt.Println("Daily job executed at midnight!")
	})
	if err != nil {
		log.Fatalf("Error adding cron job 2: %v", err)
	}

	c.Start()
	s := &server{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "qweqwe")
	})

	e.GET("/start-job", s.startScheduledJobHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
