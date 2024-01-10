package hangmanweb

import (
	"github.com/Talienhyung/hangman"
)

// Init initializes the Hangman game by setting up necessary data and printing the word to be guessed
func (webData *WebData) Init() {
	var infos hangman.HangManData
	infos.SetData()
	webData.Hangman = &infos
}

// Save updates user data accordingly
func (info *WebData) Save() {
	// Check if the player won the previous game
	if info.Status == "WIN" {
		// Update score and win statistics based on the game level
		switch info.Level {
		case "Easy":
			info.Data.Score += 1
			info.Data.WinEasy += 1
		case "Medium":
			info.Data.Score += 2
			info.Data.WinMedium += 1
		case "Hard":
			info.Data.Score += 3
			info.Data.WinHard += 1
		}

		// Increment overall win count
		info.Data.Win += 1

		// Update the best score if the current score is higher
		if info.Data.Score > info.Data.BestScore {
			info.Data.BestScore = info.Data.Score
		}
	} else if info.Status == "LOSE" {
		// If the player lost, increment the lose count and reset the score
		info.Data.Lose += 1
		info.Data.Score = 0
	}

	// Upload updated user data
	info.Data.UploadUserData(ReadAllData())
}

// Reload reloads the game state based on the provided level
func (info *WebData) Reload() {
	info.Hangman.SetData()

	// Reset game status
	info.Status = ""
}
