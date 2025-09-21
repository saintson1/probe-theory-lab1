package main

import (
	"fmt"
	"os"
	game "prob-theory-lab1/src/lib/game"
	"strconv"
)

func main() {
	tracer, err := os.OpenFile(os.Getenv("PROB_THEORY_TRACE"), os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		fmt.Printf("file open error: %#v", err)
		panic(err)
	}
	defer tracer.Close()

	results, err := os.OpenFile(os.Getenv("PROB_THEORY_RES"), os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		fmt.Printf("file open error: %#v", err)
		panic(err)
	}
	defer results.Close()

	fmt.Fprintf(results, "red\tgreen\tres\n")
	count, err := strconv.Atoi(os.Getenv("PROB_THEORY_COUNT"))
	if err != nil {
		panic(err)
	}

	for range count {
		game.PlayOnce(tracer, results)
	}
}
