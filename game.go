package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type game struct {
	moves    []string
	whosTurn int
}

func (g game) start() {
	g.moves = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	g.whosTurn = 1
	g.loop()
}

func (g game) gameover() bool {
	//game is won
	if g.gameIsWon(1) {
		fmt.Println("The game is over. Player 1 has won.")
		return true
	}
	if g.gameIsWon(2) {
		fmt.Println("The game is over. Player 2 has won.")
		return true
	}

	//is board full
	if g.boardIsFull() {
		fmt.Println("The game is over. There is no winner.")
		return true
	}

	//else
	return false
}

func (g game) gameIsWon(player int) bool {
	symbol := "X"
	if player == 2 {
		symbol = "O"
	}

	//rows
	if g.moves[0] == symbol && g.moves[1] == symbol && g.moves[2] == symbol {
		return true
	}
	if g.moves[3] == symbol && g.moves[4] == symbol && g.moves[5] == symbol {
		return true
	}
	if g.moves[6] == symbol && g.moves[7] == symbol && g.moves[8] == symbol {
		return true
	}

	//columns
	if g.moves[0] == symbol && g.moves[3] == symbol && g.moves[6] == symbol {
		return true
	}
	if g.moves[1] == symbol && g.moves[4] == symbol && g.moves[7] == symbol {
		return true
	}
	if g.moves[2] == symbol && g.moves[5] == symbol && g.moves[8] == symbol {
		return true
	}

	//diagonal
	if g.moves[0] == symbol && g.moves[4] == symbol && g.moves[8] == symbol {
		return true
	}
	if g.moves[2] == symbol && g.moves[4] == symbol && g.moves[6] == symbol {
		return true
	}

	return false
}

func (g game) boardIsFull() bool {
	for _, v := range g.moves {
		if v == " " {
			return false
		}
	}

	return true
}

func (g game) loop() {
	for !g.gameover() {
		g.draw(true)
		g.prompt(fmt.Sprintf("Player %v it is your turn. Please select a position (0 to 8): ", g.whosTurn))

		if g.whosTurn == 1 {
			g.whosTurn = 2
		} else {
			g.whosTurn = 1
		}
	}
	g.draw(false)
}

func (g game) prompt(msg string) {
	var i int

	fmt.Printf(msg)
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
		g.prompt(fmt.Sprintf("An error occured: %v. Try again: ", err.Error()))
		return
	}

	if i < 0 || i > 8 {
		g.prompt("That is not a vaid position. Please select a position between 0 and 8: ")
		return
	}

	if g.moves[i] != " " {
		g.prompt("That position is already filled. Please select another position between 0 and 8: ")
		return
	}

	//store players selection
	g.moves[i] = g.getPlayerSymbol()
}

func (g game) getPlayerSymbol() string {
	if g.whosTurn == 1 {
		return "X"
	}

	return "O"
}

func (g game) draw(clear bool) {
	if clear {
		CallClear()
	}

	for i, v := range g.moves {
		//top and middle lines
		if i == 0 || i == 3 || i == 6 {
			if i > 0 {
				fmt.Println("")
			}

			fmt.Println("-------------")
			fmt.Printf("|")
		}

		fmt.Printf(" %v |", v)

		//bottom line
		if i == 8 {
			fmt.Println("\n-------------")
		}
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
