package game

import (
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/fight"
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/relator"
)

const (
	PLAYER_1_NAME = "Tonyn"
	PLAYER_2_NAME = "Arnaldor"
	PLAYER_1_LIFE = 6
	PLAYER_2_LIFE = 6
)

type port struct {
	FightService   fight.Service
	RelatorService relator.Service
}

type Service interface {
	Start(game Game) error
}

func NewService(fightService fight.Service, relator relator.Service) Service {
	return &port{fightService, relator}
}

func (port *port) Start(game Game) error {
	on := true
	turn := 0
	game.Player1.Name = PLAYER_1_NAME
	game.Player2.Name = PLAYER_2_NAME
	game.Player1.Life = PLAYER_1_LIFE
	game.Player2.Life = PLAYER_2_LIFE
	FirstPlayer, SecondPlayer := port.determinatePlayer(game.Player1, game.Player2)
	for on {
		if turn < len(FirstPlayer.Moves) {
			atack, err := port.FightService.CheckMoves(FirstPlayer.Moves[turn], FirstPlayer.Hits[turn], FirstPlayer.Name)
			if err != nil {
				return err
			}
			port.RelatorService.Relate(*atack, FirstPlayer.Name)
			SecondPlayer.Life = SecondPlayer.Life - atack.Damage
			if SecondPlayer.Life <= 0 {
				on = false
				port.RelatorService.Victory(FirstPlayer.Name, FirstPlayer.Life)
				break
			}
		}
		if turn < len(SecondPlayer.Moves) {
			atack, err := port.FightService.CheckMoves(SecondPlayer.Moves[turn], SecondPlayer.Hits[turn], SecondPlayer.Name)
			if err != nil {
				return err
			}
			port.RelatorService.Relate(*atack, SecondPlayer.Name)
			FirstPlayer.Life = FirstPlayer.Life - atack.Damage
			if FirstPlayer.Life <= 0 {
				on = false
				port.RelatorService.Victory(SecondPlayer.Name, SecondPlayer.Life)
				break
			}
		}
		turn++
	}
	return nil
}

func (port *port) determinatePlayer(player1, player2 Player) (*Player, *Player) {
	totalPlayer1 := len(player1.Moves) + len(player1.Hits)
	totalPlayer2 := len(player2.Moves) + len(player2.Hits)
	if totalPlayer1 > totalPlayer2 {
		return &player2, &player1
	} else if totalPlayer2 > totalPlayer1 {
		return &player1, &player2
	} else {
		if len(player1.Moves) > len(player2.Moves) {
			return &player2, &player1
		} else if len(player2.Moves) > len(player1.Moves) {
			return &player1, &player2
		} else {
			if len(player1.Hits) > len(player2.Hits) {
				return &player2, &player1
			} else if len(player2.Hits) > len(player1.Hits) {
				return &player1, &player2
			} else {
				return &player1, &player2
			}
		}
	}
}
