package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositionsByCoordinates("47.6062100", "-122.3320700")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
