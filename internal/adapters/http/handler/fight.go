package handler

import (
	"net/http"

	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/game"
	"github.com/gin-gonic/gin"
)

type GameHanlder interface {
	RegisterHanlder(*gin.Engine)
	Game(*gin.Context)
}

type port struct {
	GameService game.Service
}

func NewGameHandler(gameService game.Service) GameHanlder {
	return &port{gameService}
}

func (port *port) RegisterHanlder(router *gin.Engine) {
	router.POST("/game", port.Game)
}

func (port *port) Game(ctx *gin.Context) {
	var game game.Game
	err := ctx.BindJSON(&game)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid request"})
		return
	}
	port.GameService.Start(game)
	ctx.JSON(http.StatusOK, gin.H{})
}
