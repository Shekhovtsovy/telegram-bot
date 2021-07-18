package main

import (
	api2 "bot/internal/api"
	"bot/internal/config"
	"bot/internal/telegram"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Started")

	cfg := config.GetConfig()

	apiServer := api2.NewServer()
	bot := telegram.NewBot(cfg)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := apiServer.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	go func() {
		bot.Listen()
	}()

	wg.Wait()

	fmt.Println("Shutdown")
}
