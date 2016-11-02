package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositions("go", "", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p[0])
}
