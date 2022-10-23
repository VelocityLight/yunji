package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"yunji/api"
	"yunji/configs"
	"yunji/internal/app/data_fetcher"
	"yunji/internal/service/monitor"
	"yunji/utils/log"
)

var (
	config = flag.String("config", "./config.yaml", "specify the config path")
)

func main() {
	flag.Parse()
	configs.LoadConfig(*config)

	ginRouter := api.NewGinRouter()
	ginRouter = api.RouteWebsite(ginRouter, "website/build/")

	log.Log.Infof(configs.Config.DSN)
	log.Log.Infof(configs.Config.Feishu.AppId)
	log.Log.Infof(*config)

	h := api.NewHTTPHandler(ginRouter, configs.Config)
	go func() {
		err := h.Gin.Run(":8080")
		log.Log.Infof("shutting down the server, err=%v", err)
	}()

	go data_fetcher.FetchData()
	m := monitor.NewMonitor(configs.Config)
	go m.Inspect(context.TODO())

	if err := waitShutdown(h); err != nil {
		log.Log.Errorf(err, "shutdown server error")
	}
}

func waitShutdown(h *api.HTTPHandler) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	return h.Shutdown()
}
