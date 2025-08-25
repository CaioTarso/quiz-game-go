package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func (g *GameState) Init() {
    fmt.Println("Olá! Seja bem-vindo(a) ao Quiz Game em Go!")
	fmt.Println("Por favor, insira seu nome: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a String")
	}

	g.PlayerName = name

	fmt.Printf("Ok! Vamos começar o Quiz, %s", g.PlayerName)
}

func main() {
    gameteste := &GameState{}
	gameteste.Init()
}


