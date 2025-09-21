package game_simulation

import (
	"encoding/json"
	"fmt"
	"io"
)

func GetTrace(currStepNum int, currStep StepType, steps [][]StepType) (res []GameState) {
	if currStep.prevInd == -1 {
		res = append(res, currStep.state)
		return
	}

	res = append(GetTrace(currStepNum-1, steps[currStepNum-1][currStep.prevInd], steps), currStep.state)
	return
}

func OutputTrace(writer io.Writer, trace []GameState) {
	data, err := json.MarshalIndent(trace, "", "  ")
	if err != nil {
		fmt.Errorf("json marshaling error: %+v\n\n", err)
		panic(err)
	}

	_, err = fmt.Fprintf(writer, "%+v\n", string(data))
	if err != nil {
		fmt.Errorf("file write error: %+v\n\n", err)
		panic(err)
	}
}
