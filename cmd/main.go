package main

import (
	"fmt"

	twoPeople "github.com/NeverlandMJ/Tic-Tac-Toe/pkg/twoPeople"
	withcomputer "github.com/NeverlandMJ/Tic-Tac-Toe/pkg/withComputer"

	//"github.com/NeverlandMJ/Tic-Tac-Toe/types"
	"os"
	"strings"
)

func main() {
	var ans string
	fmt.Println("If you want to play the game type: yes")
	fmt.Scan(&ans)
	ans = strings.ToLower(ans)
	if ans == "yes" {
		kind := 0
		fmt.Println("👥(2) OR 💻(1): ")
		fmt.Scan(&kind)
		if kind == 2 {
			res := twoPeople.Play()
			if res == "" {
				fmt.Println("DRAW!")

			} else {
				fmt.Println("WINNER: ", res)

			}
		}else if kind == 1 {
			res := withcomputer.Play()
			if res == "" {
				fmt.Println("DRAW!")

			} else {
				fmt.Println("WINNER: ", res)

			}
		}

	} else {
		os.Exit(1)
	}

}
