package relator

import (
	"fmt"

	"github.com/HongXiangZuniga/TalanaKombat/internal/core/domain/fight"
)

type port struct {
	LimitEpicLife int
	LimitEasyLife int
}

type Service interface {
	Relate(atack fight.Atack, name string)
	Victory(name string, life int)
}

func NewService(limitEpicLife, limitEasyLife int) Service {
	return &port{limitEpicLife, limitEasyLife}
}

func (port *port) Relate(atack fight.Atack, name string) {
	if atack.Name == "Patada" || atack.Name == "Golpe" {
		switch name {
		case "Tonyn":
			msg := name + " "
			for _, mov := range atack.Move {
				if mov == 'D' {
					msg = msg + "Avanza y"
				} else if mov == 'S' {
					msg = msg + "Retrocede y"
				} else {
					msg = msg + "Se queda quieto y"
				}
			}
			msg = msg + "ha realizado un " + atack.Name + " con un daño de " + fmt.Sprint(atack.Damage)
			fmt.Println(msg)

		case "Arnaldor":
			msg := name + " "
			for _, mov := range atack.Move {
				if mov == 'S' {
					msg = msg + "Avanza y"
				} else if mov == 'D' {
					msg = msg + "Retrocede y"
				} else {
					msg = msg + "Se queda quieto y"
				}
			}
			msg = msg + "ha realizado un " + atack.Name + " con un daño de " + fmt.Sprint(atack.Damage)
			fmt.Println(msg)
		}
	} else if atack.Name == "" {
		fmt.Printf("%s se mueve\n", name)
	} else {
		fmt.Printf("%s ha realizado un %s con un daño de %d\n", name, atack.Name, atack.Damage)
	}
}

func (port *port) Victory(name string, life int) {
	msg := "ha tenido una"
	if life >= port.LimitEpicLife {
		msg = msg + " epica"
	} else if life <= port.LimitEasyLife {
		msg = msg + " facil"
	}
	fmt.Printf("%s %s victoria y queda con %d de energia \n", name, msg, life)
}
