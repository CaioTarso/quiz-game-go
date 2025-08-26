package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type GameState struct {
	PlayerName string
	PlayerPoints int
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

	fmt.Printf("Ok! Vamos começar o Quiz, %s \n", g.PlayerName)
}

func (g *GameState) ProcessCSV() {
	f, error := os.Open("quizgolang.csv")
	if error != nil {
		panic("Erro ao ler arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler arquivo")
	}
    
	for index, record := range records {
		if index > 0 {
			question := Question{
				Text: record[0],
				Options: record[1:5],
				CorrectAnswer: toInt(record[5]),
		}

		g.Questions = append(g.Questions, question)
	}

 }
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Println(index+1, question.Text)
	}
}

func main() {
    gameteste := &GameState{PlayerPoints: 0}
	go gameteste.ProcessCSV()
	gameteste.Init()
	gameteste.Run()
	
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Erro ao converter string para int")
	}

	return i
}


