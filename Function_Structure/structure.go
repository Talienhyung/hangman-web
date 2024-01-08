package hangmanweb

import . "github.com/Talienhyung/hangman"

// Main structure that groups together all the structures and variables required for the hangman web to function properly
type WebData struct {
	Hangman *HangManData
	Data    *Data
	Board   *ScoreBoard
	Status  string
	Level   string
	Theme   string
}

// Structure that stores all user data that will be stored in the database
type Data struct {
	Username  string
	Email     string
	Password  string
	Win       int
	Lose      int
	Score     int
	BestScore int
	WinHard   int
	WinMedium int
	WinEasy   int
}

// Structure for the board
type ScoreBoard struct {
	Id       int
	First    string
	Second   string
	Third    string
	Sentence string
}
