package hangmanweb

import . "github.com/Talienhyung/hangman"

type Structure struct {
	Hangman *HangManData
	Data    *Data
	Status  string
	Level   string
}

type Data struct {
	Username  string
	Email     string
	Password  string
	Win       int
	Loose     int
	Score     int
	BestScore int
	WinHard   int
	WinMedium int
	WinEasy   int
}
