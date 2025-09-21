package game_simulation

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
)

type GameState struct {
	IsEndGame            bool
	EmptyCellDesision    string
	NotEmptyCellDesision string
	CubeVals             []int
	Field                []int
	Oxygen               int
	PlayerInd            int
	Players              []Player
}

func Copy(src []Player) (dst []Player) {
	dst = make([]Player, len(src))
	for ind, p := range src {
		dst[ind] = p.Copy()
	}
	return
}

func (state GameState) Copy() GameState {
	return GameState{
		IsEndGame:            state.IsEndGame,
		Field:                append([]int{}, state.Field...),
		Oxygen:               state.Oxygen,
		PlayerInd:            state.PlayerInd,
		Players:              append([]Player{}, state.Players...),
		CubeVals:             append(state.CubeVals),
		EmptyCellDesision:    state.EmptyCellDesision,
		NotEmptyCellDesision: state.NotEmptyCellDesision,
	}
}

func (game *GameState) getCurrPlayerPtr() (ptr *Player) {
	ptr = &game.Players[game.PlayerInd]
	game.PlayerInd = (game.PlayerInd + 1) % len(game.Players)
	return
}

func (game *GameState) IsAllOnBoat() bool {
	for PlayerInd := range game.Players {
		if game.Players[PlayerInd].Pos >= 0 {
			return false
		}
	}
	return true
}

func getScore(currency int) int {
	koeff := rand.Intn(4)
	if currency == sell1 {
		return koeff
	}

	if currency == sell2 {
		return 4 + koeff
	}

	if currency == sell3 {
		return 8 + koeff
	}

	if currency == sell4 {
		return 12 + koeff
	}

	if currency == sellEmpty {
		return 0
	}

	return -999999
}

func (game *GameState) DoStep(
	actionNotEmpty, actionEmpty int,
	cubeVal int) (IsEndGame bool) {

	if game.IsAllOnBoat() {
		IsEndGame = true
		return
	}

	currPlayer := game.getCurrPlayerPtr()

	anotherPlayerPos := game.Players[game.PlayerInd].Pos

	stepVal := cubeVal - len(currPlayer.Currency)
	stepVal = max(0, stepVal)

	IsEndGame = false

	game.Oxygen -= len(currPlayer.Currency)
	if game.Oxygen <= 0 {
		IsEndGame = true
		return
	}

	posAfter := currPlayer.Pos
	currPlayer.Pos -= stepVal
	if currPlayer.Pos <= anotherPlayerPos && posAfter > anotherPlayerPos {
		currPlayer.Pos -= 1
	}

	if currPlayer.Pos < 0 {
		for _, currency := range currPlayer.Currency {
			currPlayer.Point += getScore(currency)
		}
		currPlayer.Currency = []int{sellEmpty}
		return
	}

	currCell := &(game.Field[currPlayer.Pos])
	switch *currCell {
	case sellEmpty:
		switch actionEmpty {
		case ActionHold:
			return
		case ActionPut:
			sort.Ints(currPlayer.Currency)
			if len(currPlayer.Currency) != 0 {
				*currCell = currPlayer.Currency[0]
				currPlayer.Currency = currPlayer.Currency[1:len(currPlayer.Currency)]
			}
		}
	default:
		switch actionNotEmpty {
		case ActionHold:
			return
		case ActionTake:
			currPlayer.Currency = append(currPlayer.Currency, *currCell)
			*currCell = sellEmpty
		}
	}

	if game.Oxygen <= 0 {
		IsEndGame = true
	}

	return
}

func PlayOnce(tracer io.Writer, results io.Writer) {
	steps := GetDefaultGame()
	steps[0] = []StepType{steps[0][rand.Intn(2)]}

	type stepArrType []StepType
	endGameSteps := make(map[int]stepArrType, StepsCount)

	stepCnt := 0
	for {
		res, isEnd := StepPropagate(&steps, stepCnt)

		if len(res) != 0 {
			endGameSteps[stepCnt+1] = res
		}
		stepCnt += 1
		if isEnd {
			break
		}
	}

	for stepNum := range endGameSteps {
		for scdKey := range endGameSteps[stepNum] {
			endGameStep := endGameSteps[stepNum][scdKey]
			trace := GetTrace(
				stepNum,
				endGameStep,
				steps,
			)

			endGameStep.state.Players[1].Point = max(27, endGameStep.state.Players[1].Point-4)
			fmt.Fprintf(results, "%v\t%v\t%v\n",
				endGameStep.state.Players[0].Point,
				endGameStep.state.Players[1].Point,
				endGameStep.state.Players[0].Point > endGameStep.state.Players[1].Point,
			)

			OutputTrace(tracer, trace)
		}
	}
}
