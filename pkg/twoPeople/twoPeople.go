package tools

import (
	"errors"
	"fmt"
	"github.com/NeverlandMJ/Tic-Tac-Toe/types"
	"math/rand"
	"strconv"
	"unicode"
)

var errChosen = errors.New("this position is choosen")
var errInvalidInput = errors.New("invalid input")

// PrintTable prints the table of Tic Tac Toe
func PrintTable(t types.Table) {

	for _, v := range t {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}
}

// Positions takes players' names and randomly chooses who will play first
func Positions() types.Ox {
	game := types.Ox{}
	fmt.Println("1st player please enter your name: ")
	fmt.Scan(&game.Player1)
	fmt.Println("2nd player please enter your name: ")
	fmt.Scan(&game.Player2)
	rand.Seed(10)
	turn := rand.Intn(2) // randomly chooses turn for players
	if turn == 0 {
		fmt.Println(game.Player1, " plays with X")
		fmt.Println(game.Player2, " plays with O")
		game.Turn = types.X
		game.Current = game.Player1
	} else {
		fmt.Println(game.Player2, " plays with X")
		fmt.Println(game.Player1, " plays with O")
		game.Turn = types.O
		game.Current = game.Player2
	}
	game.GameOver = false // start the game
	return game
}

// Check determines who won or is it draw
func Check(t types.Table) (bool, string) {
	x := types.X
	o := types.O
	Xcase1 := (t[0][0] == x && t[0][1] == x && t[0][2] == x)
	Ocase1 := (t[0][0] == o && t[0][1] == o && t[0][2] == o)
	Xcase2 := (t[1][0] == x && t[1][1] == x && t[1][2] == x)
	Ocase2 := (t[1][0] == o && t[1][1] == o && t[1][2] == o)
	Xcase3 := (t[2][0] == x && t[2][1] == x && t[2][2] == x)
	Ocase3 := (t[2][0] == o && t[2][1] == o && t[2][2] == o)

	Xcase4 := (t[0][0] == x && t[1][0] == x && t[2][0] == x)
	Ocase4 := (t[0][0] == o && t[1][0] == o && t[2][0] == o)
	Xcase5 := (t[0][1] == x && t[1][1] == x && t[2][1] == x)
	Ocase5 := (t[0][1] == o && t[1][1] == o && t[2][1] == o)
	Xcase6 := (t[0][2] == x && t[1][2] == x && t[2][2] == x)
	Ocase6 := (t[0][2] == o && t[1][2] == o && t[2][2] == o)

	Xcase7 := (t[0][2] == x && t[1][1] == x && t[2][0] == x)
	Ocase7 := (t[0][2] == o && t[1][1] == o && t[2][0] == o)
	Xcase8 := (t[0][0] == x && t[1][1] == x && t[2][2] == x)
	Ocase8 := (t[0][0] == o && t[1][1] == o && t[2][2] == o)

	draw := true

	for _, v1 := range t {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				if unicode.IsNumber(v3) {
					draw = false
				}
			}
		}
	}

	if Xcase1 || Xcase2 || Xcase3 || Xcase4 || Xcase5 || Xcase6 || Xcase7 || Xcase8 {
		return true, x
	} else if Ocase1 || Ocase2 || Ocase3 || Ocase4 || Ocase5 || Ocase6 || Ocase7 || Ocase8 {
		return true, o
	} else if draw {
		return true, types.Draw
	} else {
		return false, ""
	}
}

// Change changes values into x or o according to players choice
// if the position is already choosen it returns error "choosen"
func Change(t types.Table, input string, pl types.Ox) (types.Table, error) {
	x, o := types.X, types.O
	n, err := strconv.Atoi(input)
	if err != nil {
		return t, errInvalidInput
	}
	if n > 9 {
		return t, errInvalidInput
	} else {
		if pl.Turn == x {
			switch n {
			case 1:
				if t[0][0] == o || t[0][0] == x {
					return t, errChosen
				} else {
					t[0][0] = x
					return t, nil
				}

			case 2:
				if t[0][1] == o || t[0][1] == x {
					return t, errChosen
				} else {
					t[0][1] = x
					return t, nil
				}
			case 3:
				if t[0][2] == o || t[0][2] == x {
					return t, errChosen
				} else {
					t[0][2] = x
					return t, nil
				}
			case 4:
				if t[1][0] == o || t[1][0] == x {
					return t, errChosen
				} else {
					t[1][0] = x
					return t, nil
				}
			case 5:
				if t[1][1] == o || t[1][1] == x {
					return t, errChosen
				} else {
					t[1][1] = x
					return t, nil
				}
			case 6:
				if t[1][2] == o || t[1][2] == x {
					return t, errChosen
				} else {
					t[1][2] = x
					return t, nil
				}
			case 7:
				if t[2][0] == o || t[2][0] == x {
					return t, errChosen
				} else {
					t[2][0] = x
					return t, nil
				}
			case 8:
				if t[2][1] == o || t[2][1] == x {
					return t, errChosen
				} else {
					t[2][1] = x
					return t, nil
				}
			case 9:
				if t[2][2] == o || t[2][2] == x {
					return t, errChosen
				} else {
					t[2][2] = x
					return t, nil
				}
			}
		} else if pl.Turn == o {
			switch n {
			case 1:
				if t[0][0] == o || t[0][0] == x {
					return t, errChosen
				} else {
					t[0][0] = o
					return t, nil
				}

			case 2:
				if t[0][1] == o || t[0][1] == x {
					return t, errChosen
				} else {
					t[0][1] = o
					return t, nil
				}
			case 3:
				if t[0][2] == o || t[0][2] == x {
					return t, errChosen
				} else {
					t[0][2] = o
					return t, nil
				}
			case 4:
				if t[1][0] == o || t[1][0] == x {
					return t, errChosen
				} else {
					t[1][0] = o
					return t, nil
				}
			case 5:
				if t[1][1] == o || t[1][1] == x {
					return t, errChosen
				} else {
					t[1][1] = o
					return t, nil
				}
			case 6:
				if t[1][2] == o || t[1][2] == x {
					return t, errChosen
				} else {
					t[1][2] = o
					return t, nil
				}
			case 7:
				if t[2][0] == o || t[2][0] == x {
					return t, errChosen
				} else {
					t[2][0] = o
					return t, nil
				}
			case 8:
				if t[2][1] == o || t[2][1] == x {
					return t, errChosen
				} else {
					t[2][1] = o
					return t, nil
				}
			case 9:
				if t[2][2] == o || t[2][2] == x {
					return t, errChosen
				} else {
					t[2][2] = o
					return t, nil
				}
			}
		}
	}
	return t, nil
}

// SwapCurrentPlayer swaps turns of players
func SwapCurrentPlayer(OX types.Ox) types.Ox {
	if OX.Current == OX.Player1 {
		OX.Current = OX.Player2
	} else {
		OX.Current = OX.Player1
	}

	return OX
}

// Play plays the game
func Play() string {
	pl := Positions()
	table := types.Table{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	PrintTable(table)
	for !pl.GameOver {
		var ch string
		if pl.Turn == types.X {
		chooseX:
			fmt.Println(pl.Current, "(X): ")
			fmt.Scan(&ch)
			table, err := Change(table, ch, pl)
			if err != nil {
				fmt.Println(err)
				goto chooseX
			}
			IsgameOver, winner := Check(table)
			if IsgameOver && winner == types.Draw {
				pl.GameOver = true
				pl.Winner = winner
			} else if IsgameOver && winner != types.Draw {
				pl.GameOver = true
				pl.Winner = pl.Current
			} else {
				pl.GameOver = false
				pl.Turn = types.O
				pl = SwapCurrentPlayer(pl)
			}
		} else if pl.Turn == types.O {
		chooseO:
			fmt.Println(pl.Current, "(O): ")
			fmt.Scan(&ch)
			table, err := Change(table, ch, pl)
			if err != nil {
				fmt.Println(err)
				goto chooseO
			}
			IsgameOver, winner := Check(table)
			if IsgameOver && winner == types.Draw {
				pl.GameOver = true
				pl.Winner = winner
			} else if IsgameOver && winner != types.Draw {
				pl.GameOver = true
				pl.Winner = pl.Current
			} else {
				pl.GameOver = false
				pl.Turn = types.X
				pl = SwapCurrentPlayer(pl)
			}
		}
		PrintTable(table)
	}

	return pl.Winner
}
