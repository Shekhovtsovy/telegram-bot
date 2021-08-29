package main

import (
	servApi "bot/internal/api"
	"bot/internal/config"
	"bot/internal/db/postgres"
	"bot/internal/logger"
	"bot/internal/repository"
	"bot/internal/telegram"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

func main() {
	cfg := config.GetConfig()
	log := logger.NewLogger(cfg.Log.Facility)
	log.Info("bot started")
	apiServer := servApi.NewServer()
	db, err := postgres.GetDb(cfg)
	if err != nil {
		log.Error("can`t connect to database", zap.String("details", err.Error()))
		panic("can`t connect to database")
	}
	// telegram bot init
	userRep := repository.NewUser(db)
	msgRep := repository.NewMessage(db)
	bot := telegram.NewBot(cfg, msgRep, userRep)
	// run
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := apiServer.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Error("can`t start web server", zap.String("details", err.Error()))
			panic("can`t start web server")
		}
	}()
	go func() {
		bot.Listen()
	}()
	wg.Wait()

	log.Info("bot shutdown")
}
