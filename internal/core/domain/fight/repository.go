package fight

type FightRepository interface {
	CheckMovePlayer1(move, hit string) *Atack
	CheckMovePlayer2(move, hit string) *Atack
}
