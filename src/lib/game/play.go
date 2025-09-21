package game_simulation

import (
	"math/rand"
)

type StepType struct {
	prevInd int
	state   GameState
}

func (step StepType) Copy() StepType {
	return StepType{
		prevInd: step.prevInd,
		state:   step.state.Copy(),
	}
}

func GetDefaultGame() [][]StepType {
	steps := make([][]StepType, StepsCount)
	steps[0] = append(steps[0],
		StepType{
			state:   InitStartGame(0),
			prevInd: -1,
		},
		StepType{
			state:   InitStartGame(1),
			prevInd: -1,
		},
	)
	return steps
}

func GetDesision() (actionNotEmptyInd, actionEmptyInd, cubeVal1, cubeVal2 int) {
	actionNotEmptyInd = desisonMap["NotEmptyCase"][rand.Intn(len(desisonMap["NotEmptyCase"]))]
	actionEmptyInd = desisonMap["EmptyCase"][rand.Intn(len(desisonMap["EmptyCase"]))]
	cubeVal1 = desisonMap["Cube"][rand.Intn(len(desisonMap["Cube"]))]
	cubeVal2 = desisonMap["Cube"][rand.Intn(len(desisonMap["Cube"]))]

	// actionNotEmptyInd = desisonMap["NotEmptyCase"][0]
	// actionEmptyInd = desisonMap["EmptyCase"][0]
	// cubeValue = desisonMap["Cube"][0] + desisonMap["Cube"][0]
	return
}

func StepPropagate(steps *[][]StepType, currStep int) ([]StepType, bool) {
	finishSteps := []StepType{}

	if len((*steps)[currStep]) == 0 {
		return []StepType{}, true
	}

	for stateInd := range (*steps)[currStep] {
		if (*steps)[currStep][stateInd].state.IsEndGame {
			continue
		}

		currState := (*steps)[currStep][stateInd].Copy()
		currState.prevInd = stateInd
		actionNotEmptyInd, actionEmptyInd, cubeVal1, cubeVal2 := GetDesision()

		currState.state.CubeVals = []int{cubeVal1, cubeVal2}
		if actionEmptyInd == ActionHold {
			currState.state.EmptyCellDesision = "Hold"
		} else {
			currState.state.EmptyCellDesision = "Put"
		}

		if actionNotEmptyInd == ActionHold {
			currState.state.NotEmptyCellDesision = "Hold"
		} else {
			currState.state.NotEmptyCellDesision = "Take"
		}

		isEnd := currState.state.DoStep(
			actionNotEmptyInd, actionEmptyInd, cubeVal1+cubeVal2,
		)

		if isEnd {
			// fmt.Printf("finish game:\ngame state: %+v\nstep count: %v\nstep ind %v\n", currState, currStep, stateInd)
			currState.state.IsEndGame = true
			finishSteps = append(finishSteps, currState)
		} else if currState.state.IsAllOnBoat() {
			// fmt.Printf("finish game:\ngame state: %+v\nstep count: %v\nstep ind %v\n\n", currState, currStep, stateInd)
			currState.state.IsEndGame = true
			finishSteps = append(finishSteps, currState)
		} else {
			// fmt.Printf("game state: %+v\ngame step: %v\n\n", (*steps)[currStep][stateInd], currStep)
		}
		(*steps)[currStep+1] = append((*steps)[currStep+1], currState)

	}
	return finishSteps, false
}
