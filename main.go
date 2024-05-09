package main

import (
	"restapi-altera/config"
	"restapi-altera/src"
	"sync"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(".env failed to load")
	}

}

func main() {

	config := config.NewConfig()
	server := src.InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()

}
