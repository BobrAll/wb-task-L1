package main

import (
	"fmt"
	"time"
)

// Human represents a person with basic attributes.
type Human struct {
	age        uint8
	isSleeping bool
}

// Sleep makes the Human sleep according to whether they are a student.
func (h *Human) Sleep() {
	fmt.Println("Sleeping...")

	h.isSleeping = true
	if h.IsStudent() {
		time.Sleep(8 * time.Second)
	} else {
		time.Sleep(8 * time.Hour)
	}
	h.isSleeping = false

	fmt.Println("Waking up...")
}

// IsStudent returns true if Human is a student based on age.
func (h *Human) IsStudent() bool {
	return h.age > 17 && h.age < 22
}

// Action embeds Human and adds additional behaviors.
type Action struct {
	Human
}

// PlayComputerGames simulates playing video games.
func (a *Action) PlayComputerGames() {
	fmt.Println("Playing GTA...")
	fmt.Println("Playing CS:GO...")
}

func main() {
	action := Action{Human{19, false}}
	action.Sleep()
	action.PlayComputerGames()
}
