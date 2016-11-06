package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	positions, err := githubjobs.GetPositionsByCoordinates("47.6062100", "-122.3320700")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(positions)
	for _, p := range positions {
		fmt.Println(p)
	}
}
