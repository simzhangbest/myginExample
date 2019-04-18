package main

import (
	"context"
	"fmt"
	"myginExample/models"
	"myginExample/pkg/logging"

	//"github.com/labstack/gommon/log"
	"log"
	"myginExample/routers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"myginExample/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen : %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Print("shutdown server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Print("server existing")
}
