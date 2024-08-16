package repository

import (
	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/fight"
)

type port struct {
}

func NewFightRepository() fight.FightRepository {
	return &port{}
}

func (port *port) CheckMovePlayer1(move, hit string) *fight.Atack {
	atk := fight.Atack{
		Damage: 0,
		Name:   "",
	}
	switch hit {
	case "P":
		atk.Hit = "P"
		atk.Move = move
		if len(move) > 2 && move[len(move)-3:] == "DSD" {
			atk.Damage = 3
			atk.Name = "Taladoken"
		} else {
			atk.Damage = 1
			atk.Name = "Golpe"
		}
		return &atk
	case "K":
		atk.Hit = "K"
		atk.Move = move
		if len(move) > 1 && move[len(move)-2:] == "SD" {
			atk.Damage = 2
			atk.Name = "Remuyuken"
		} else {
			atk.Damage = 1
			atk.Name = "Patada"
		}
		return &atk
	default:
		return &atk
	}
}

func (port *port) CheckMovePlayer2(move, hit string) *fight.Atack {
	atk := fight.Atack{
		Damage: 0,
		Name:   "",
	}
	switch hit {
	case "P":
		if len(move) > 2 && move[len(move)-3:] == "ASA" {
			atk.Damage = 2
			atk.Name = "Taladoken"
		} else {
			atk.Damage = 1
			atk.Name = "Golpe"
		}
		return &atk
	case "K":
		if len(move) > 1 && move[len(move)-2:] == "SA" {
			atk.Damage = 3
			atk.Name = "Remuyuken"
		} else {
			atk.Damage = 1
			atk.Name = "Patada"
		}
		return &atk
	default:
		return &atk
	}
}
