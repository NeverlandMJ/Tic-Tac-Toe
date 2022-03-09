package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	//"unicode/utf8"
)

type table [][]string

func PrintTable (t table)  {
	
	for _, v := range t {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}

	
}

type ox struct {
	player1 string
	player2 string
	trun string
	gameOver bool
	winner string
	current string
}
const (
	x = "X"
	o = "O"
)
func Positions() ox {
	game := ox {}
	fmt.Println("1st player please enter your name: ")
	fmt.Scan(&game.player1)
	fmt.Println("2nd player please enter your name: ")
	fmt.Scan(&game.player2)
	turn := rand.Intn(2) // randomly chooses turn for players
	if turn == 1 {
		fmt.Println(game.player1, " is X")
		fmt.Println(game.player2, " is O")
		game.trun = x
		game.current = game.player1
	}else {
		fmt.Println(game.player2, " is X")
		fmt.Println(game.player1, " is O")
		game.trun = o
		game.current = game.player2
	}
	game.gameOver = false
	return game

}

func Check(t table) (bool, string) {
	Xcase1 := (t[0][0]==x && t[0][1]==x && t[0][2]==x) 
	Ocase1 := (t[0][0]==o && t[0][1]==o && t[0][2]==o)
	Xcase2 := (t[1][0]==x && t[1][1]==x && t[1][2]==x)
	Ocase2 := (t[1][0]==o && t[1][1]==o && t[1][2]==o)
	Xcase3 := (t[2][0]==x && t[2][1]==x && t[2][2]==x)
	Ocase3 := (t[2][0]==o && t[2][1]==o && t[2][2]==o)

	Xcase4 := (t[0][0]==x && t[1][0]==x && t[2][0]==x) 
	Ocase4 := (t[0][0]==o && t[1][0]==o && t[2][0]==o)
	Xcase5 := (t[0][1]==x && t[1][1]==x && t[2][1]==x) 
	Ocase5 := (t[0][1]==o && t[1][1]==o && t[2][1]==o)
	Xcase6 := (t[0][2]==x && t[1][2]==x && t[2][2]==x) 
	Ocase6 := (t[0][2]==o && t[1][2]==o && t[2][2]==o)

	Xcase7 := (t[0][2]==x && t[1][1]==x && t[2][0]==x) 
	Ocase7 := (t[0][2]==o && t[1][1]==o && t[2][0]==o)
	Xcase8 := (t[0][0]==x && t[1][1]==x && t[2][2]==x) 
	Ocase8 := (t[0][0]==o && t[1][1]==o && t[2][2]==o)

	if Xcase1 || Xcase2 || Xcase3 || Xcase4 || Xcase5 || Xcase6 || Xcase7 || Xcase8 {
		return true, x
	}else if Ocase1 || Ocase2 || Ocase3 || Ocase4 || Ocase5 || Ocase6 || Ocase7 || Ocase8 {
		return true, o
	}else {
		return false, ""
	}

	
}

func Change(t table, n int, pl ox) (table, error) {
	
	if pl.trun == x {
		switch n{
		case 1: 
			if t[0][0] == o || t[0][0] == x {
				return t, errors.New("choosen")
			}else {
				t[0][0] = x
				return t, nil
			}
			
		case 2:
			if t[0][1] == o || t[0][1] == x {
				return t, errors.New("choosen")
			}else {
				t[0][1] = x
				return t, nil
			}
		case 3:
			if t[0][2] == o || t[0][2] == x {
				return t, errors.New("choosen")
			}else {
				t[0][2] = x
				return t, nil
			}
		case 4:
			if t[1][0] == o || t[1][0] == x {
				return t, errors.New("choosen")
			}else {
				t[1][0] = x
				return t, nil
			}
		case 5:
			if t[1][1] == o || t[1][1] == x {
				return t, errors.New("choosen")
			}else {
				t[1][1] = x
				return t, nil
			}
		case 6:
			if t[1][2] == o || t[1][2] == x {
				return t, errors.New("choosen")
			}else {
				t[1][2] = x
				return t, nil
			}
		case 7:
			if t[2][0] == o || t[2][0] == x {
				return t, errors.New("choosen")
			}else {
				t[2][0] = x
				return t, nil
			}
		case 8:
			if t[2][1] == o || t[2][1] == x {
				return t, errors.New("choosen")
			}else {
				t[2][1] = x
				return t, nil
			}
		case 9:
			if t[2][2] == o || t[2][2] == x {
				return t, errors.New("choosen")
			}else {
				t[2][2] = x
				return t, nil
			}
		}
	}else if pl.trun == o {
		switch n{
			case 1: 
				if t[0][0] == o || t[0][0] == x {
					return t, errors.New("choosen")
				}else {
					t[0][0] = o
					return t, nil
				}
				
			case 2:
				if t[0][1] == o || t[0][1] == x {
					return t, errors.New("choosen")
				}else {
					t[0][1] = o
					return t, nil
				}
			case 3:
				if t[0][2] == o || t[0][2] == x {
					return t, errors.New("choosen")
				}else {
					t[0][2] = o
					return t, nil
				}
			case 4:
				if t[1][0] == o || t[1][0] == x {
					return t, errors.New("choosen")
				}else {
					t[1][0] = o
					return t, nil
				}
			case 5:
				if t[1][1] == o || t[1][1] == x {
					return t, errors.New("choosen")
				}else {
					t[1][1] = o
					return t, nil
				}
			case 6:
				if t[1][2] == o || t[1][2] == x {
					return t, errors.New("choosen")
				}else {
					t[1][2] = o
					return t, nil
				}
			case 7:
				if t[2][0] == o || t[2][0] == x {
					return t, errors.New("choosen")
				}else {
					t[2][0] = o
					return t, nil
				}
			case 8:
				if t[2][1] == o || t[2][1] == x {
					return t, errors.New("choosen")
				}else {
					t[2][1] = o
					return t, nil
				}
			case 9:
				if t[2][2] == o || t[2][2] == x {
					return t, errors.New("choosen")
				}else {
					t[2][2] = o
					return t, nil
				}
			}
	}

	return t, nil
}


func SwapCurrentPlayer(OX ox) ox {
	if OX.current == OX.player1 {
		OX.current = OX.player2
	}else {
		OX.current = OX.player1
	}

	return OX
}

func Play() string {
	pl := Positions()
	table := table {
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	PrintTable(table)
	for !pl.gameOver {
		var ch int
		if pl.trun == x {
			chooseX:
			fmt.Println(pl.current, ": ")
			fmt.Scan(&ch)
			table, err := Change(table, ch, pl)
			if err != nil {
				fmt.Println(err)
				goto chooseX
			}
			res, _ := Check(table)
			if res {
				pl.gameOver = true
				pl.winner = pl.player1
			}else {
				pl.gameOver = false
				pl.trun = o
				pl = SwapCurrentPlayer(pl)
			}
		}else if pl.trun == o {
			chooseO:
			fmt.Println(pl.player2, ": ")
			fmt.Scan(&ch)
			table, err := Change(table, ch, pl)
			if err != nil {
				fmt.Println(err)
				goto chooseO
			}
			res, _ := Check(table)
			if res {
				pl.gameOver = true
				pl.winner = pl.player2
			}else {
				pl.gameOver = false
				pl.trun = x
				pl = SwapCurrentPlayer(pl)
			}
		}
		PrintTable(table)		
		
	}

	return pl.winner
}

func main() {
	var ans string
	fmt.Println("If you want to play the game type: yes")
	fmt.Scan(&ans)
	ans = strings.ToLower(ans)
	if ans == "yes" {
		table := table {
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}
		PrintTable(table)
		res := Play()
		fmt.Println("Winner is ", res)
	}else {
		os.Exit(1)
	}


}