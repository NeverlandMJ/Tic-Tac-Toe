package types


// Table stores table of the tic tac toe
type Table [][]string


// Ox is used to define main roles of the gane
type Ox struct {
	Player1 string // 1st player's name
	Player2 string // 2ng player's name
	Turn string    // shows whose turn now
	GameOver bool  // informs if game is over
	Winner string  // shows who is the winner
	Current string // shows whose playing now
}


// X and O
const (
	X = "X"
	O = "O"
	Draw = "draw"
)