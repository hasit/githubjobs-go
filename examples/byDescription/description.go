package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	positions, err := githubjobs.GetPositions("go", "", false)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range positions {
		fmt.Println(p)
	}
}
