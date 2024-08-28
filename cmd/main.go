package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawpawchat/core/config"
	"github.com/pawpawchat/core/internal/app"
)

func main() {
	flag.Parse()

	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = config.ConfigureLogger(cfg); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	app.Run(ctx, cfg.Env())
}
