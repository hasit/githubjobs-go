package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositionByID("4dfece3c-97c4-11e6-97f0-6745f96d2097")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
