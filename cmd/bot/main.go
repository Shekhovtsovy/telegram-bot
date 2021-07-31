package main

import (
	api2 "bot/internal/api"
	"bot/internal/config"
	"bot/internal/db/postgres"
	"bot/internal/logger"
	mRep "bot/internal/repository/message"
	uRep "bot/internal/repository/user"
	mSer "bot/internal/service/message"
	uSer "bot/internal/service/user"
	"bot/internal/telegram"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	cfg := config.GetConfig()
	log := logger.NewLogger(cfg.Log.Facility)
	log.Info("bot started")
	apiServer := api2.NewServer()
	db, err := postgres.GetDb(cfg)
	if err != nil {
		log.Error("can`t connect to database")
		panic("can`t connect to database")
	}
	messageRepository := mRep.NewRepository(db)
	messageService := mSer.NewService(messageRepository)
	userRepository := uRep.NewRepository(db)
	userService := uSer.NewService(userRepository)
	bot := telegram.NewBot(cfg, messageService, userService)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := apiServer.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Error("can`t start web server")
			panic("can`t start web server")
		}
	}()
	go func() {
		bot.Listen()
	}()
	wg.Wait()

	log.Info("bot shutdown")
}
