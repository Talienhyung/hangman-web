package hangmanweb

import (
	. "github.com/Talienhyung/hangman"
)

// Init initializes the Hangman game by setting up necessary data and printing the word to be guessed.
func (myStruct *Structure) Init() {
	var infos HangManData
	infos.SetData()
	myStruct.Hangman = &infos
}

// Reload reloads the game state based on the provided level and updates user data accordingly.
func (info *Structure) Reload(level string) {
	info.Hangman.SetData()

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
	} else if info.Status == "LOOSE" {
		// If the player lost, increment the loose count and reset the score
		info.Data.Loose += 1
		info.Data.Score = 0
	}

	// Upload updated user data
	info.Data.UploadUserData(ReadAllData())

	// Reset game status
	info.Status = ""
}
