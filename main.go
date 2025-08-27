package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	Answer int
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
			CorrectAnswer, _ := toInt(record[5])
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: CorrectAnswer,
		}

		g.Questions = append(g.Questions, question)
	}

 }
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index + 1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j + 1, option)
		}
		
		fmt.Println("Digite a alternativa: ")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])
			if err != nil{
                fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Println("Parabéns! Você acertou.")
			g.PlayerPoints += 10
		}else {
			fmt.Println("Ops! Você errou.")
			fmt.Println("----------------------------------")
		}

		
	}
}

func main() {
    game := &GameState{PlayerPoints: 0}
    go game.ProcessCSV()
	game.Init()
	game.Run()

	fmt.Printf("Fim de jogo. Você fez %d pontos\n", game.PlayerPoints)
	
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("não é permitido caractere que não seja número. Por favor, insira um número")
	}
	return i, nil
}


