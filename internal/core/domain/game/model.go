package game

type Game struct {
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
}

type Player struct {
	Name  string
	Life  int
	Moves []string `json:"movimientos"`
	Hits  []string `json:"golpes"`
}
