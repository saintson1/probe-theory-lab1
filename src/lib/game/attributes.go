package game_simulation

const (
	sellEmpty = iota
	sell1
	sell2
	sell3
	sell4
)

const (
	ActionPut = iota
	ActionTake
	ActionHold
)

const StepsCount = 15

var desisonMap = map[string][]int{
	"EmptyCase": {
		ActionPut,
		ActionHold,
	},

	"NotEmptyCase": {
		ActionTake,
		ActionHold,
	},

	"Cube": {1, 2, 3, 1, 2, 3},
}

func InitStartGame(PlayerInd int) GameState {
	return GameState{
		Field: []int{
			sell1,
			sellEmpty,
			sell1,
			sell1,
			sell1,
			sellEmpty,
		},
		Oxygen:    9,
		PlayerInd: PlayerInd,
		Players: []Player{
			Player{
				Point: 25,
				Currency: []int{
					sell4,
					sell2,
				},
				Pos: 4,
			},
			Player{
				Point: 27,
				Currency: []int{
					sell3,
					sell1,
				},
				Pos: 1,
			},
		},
	}
}
