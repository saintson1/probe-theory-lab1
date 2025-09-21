package game_simulation

const (
	RedPlayer   = 0
	GreenPlayer = 1
)

type Player struct {
	Point    int
	Currency []int
	Pos      int
}

func (player Player) Copy() Player {
	return Player{
		Point:    player.Point,
		Currency: append([]int{}, player.Currency...),
		Pos:      player.Pos,
	}
}
