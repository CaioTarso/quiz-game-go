package main

import "fmt"

type GameState struct {
	PlayerName string
	PlayerPoints string
	Questions []Question
}

type Question struct {
	Text string
	Options []string
	CorrectAnswer int
}

func main() {
    fmt.Println("Só um teste, boy")
}


