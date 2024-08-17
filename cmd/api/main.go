package main

import (
	"fmt"
	"os"

	"github.com/HongXiangZuniga/TalanaKombat/internal/adapters/http"
	"github.com/HongXiangZuniga/TalanaKombat/internal/adapters/http/handler"
	"github.com/HongXiangZuniga/TalanaKombat/internal/adapters/repository"
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/fight"
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/game"
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/relator"
	"github.com/joho/godotenv"
)

func main() {
	//try load .envfile
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not load .env file")
	}
	//Load env
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	//Init router
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	router := http.NewRouter(env)

	//repo
	fightRepo := repository.NewFightRepository()

	//Service
	fightService := fight.NewService(fightRepo)
	relatorService := relator.NewService(1, 4)
	gameService := game.NewService(fightService, relatorService)

	//Handler
	GameHanlder := handler.NewGameHandler(gameService)
	GameHanlder.RegisterHanlder(router)

	router.Run(port)
}
