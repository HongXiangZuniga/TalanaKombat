package fight

import "errors"

type port struct {
	fightRepository FightRepository
}

type Service interface {
	CheckMoves(move, hit, name string) (*Atack, error)
}

func NewService(fightRepository FightRepository) Service {
	return &port{
		fightRepository: fightRepository,
	}
}

func (port *port) CheckMoves(move, hit, name string) (*Atack, error) {
	if name == "Tonyn" { // To do common const for pkg fight and game
		return port.fightRepository.CheckMovePlayer1(move, hit), nil
	} else if name == "Arnaldor" {
		return port.fightRepository.CheckMovePlayer2(move, hit), nil
	} else {
		return nil, errors.New("player not found")
	}
}
