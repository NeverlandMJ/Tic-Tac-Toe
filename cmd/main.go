package main

import (
	
	"fmt"
	"os"
	"strings"
	"github.com/NeverlandMJ/Tic-Tac-Toe/types"
	"github.com/NeverlandMJ/Tic-Tac-Toe/pkg/tools"
	"github.com/kyokomi/emoji/v2"
)



func main() {
	var ans string
	fmt.Println("If you want to play the game type: yes")
	fmt.Scan(&ans)
	ans = strings.ToLower(ans)
	if ans == "yes" {
		table := types.Table {
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}
		tools.PrintTable(table)
		res := tools.Play()
		if res == "" {
			fmt.Println("DRAW!")
			
		}else {
			fmt.Println("WINNER: ", res)
			party := emoji.Sprint(":heart:")
			fmt.Println(strings.Repeat(party, 5))
		}
		
	}else {
		os.Exit(1)
	}


}