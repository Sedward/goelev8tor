package main

import (
	"fmt"

	"github.com/Sedward/goelev8tor/internal/elevator"
)

func main() {
	fmt.Printf("Hello!")
	e := &elevator.Elevator{Direction: elevator.Staionary}
	fmt.Printf("My Elevator %+v\n", e)
}
