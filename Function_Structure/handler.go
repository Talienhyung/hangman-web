package hangmanweb

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	. "github.com/Talienhyung/hangman"
)

// hangmanHandler handles HTTP requests related to the Hangman game
func hangmanHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	// Check if the request method is not POST
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}

	// Get the letter input from the form
	letter := r.FormValue("input")

	// Check if the letter has been used before
	if info.Hangman.UsedVerif(letter) {
		info.Status = "USED"
	} else if letter != "STOP" && letter != "QUIT" { // The main mecanic function shuts down the game in the presence of these words, which is not necessary for this game version
		// Check the main game mechanics for the input letter
		if info.Hangman.MainMecanics(letter) {
			info.Status = "WIN"
		} else if info.Hangman.EndGame() {
			// Check if the game has ended
			if string(info.Hangman.Word) == info.Hangman.ToFind {
				info.Status = "WIN"
			} else {
				info.Status = "LOOSE"
			}
		} else {
			info.Status = ""
		}
	} else {
		info.Status = "FORBIDDEN"
	}

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// levelHandler handles HTTP requests related to changing the game level
func levelHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}
	action := r.FormValue("action")
	info.Reload(action)

	// Set the word for the Hangman based on the selected level
	switch action {
	case "Easy":
		info.Hangman.SetWord(ReadTheDico("words.txt"))
		info.Level = "Easy"
	case "Medium":
		info.Hangman.SetWord(ReadTheDico("words2.txt"))
		info.Level = "Medium"
	case "Hard":
		info.Hangman.SetWord(ReadTheDico("words3.txt"))
		info.Level = "Hard"
	}

	// Print the word to be guessed (for debugging purposes)
	fmt.Println(info.Hangman.ToFind)

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// connexionHandler handles HTTP requests related to user connection (signin and login)
func connexionHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}
	data := ReadAllData()
	action := r.FormValue("action")

	// Handle different cases based on the current status
	switch info.Status {
	case "CONNEXION":
		// If the status is CONNEXION, determine the next status based on the action
		switch action {
		case "Signin":
			info.Status = "SIGNIN"
		case "Login":
			info.Status = "LOGIN"
		}
	case "SIGNIN", "SIGNIN-ERROR":
		// If the status is SIGNIN or SIGNIN-ERROR, handle the actions accordingly
		if action == "Login" {
			info.Status = "LOGIN"
		} else {
			// Get user input values
			email := r.FormValue("email")
			username := r.FormValue("username")
			passw := r.FormValue("password")

			// Check if the email is already used
			if EmailAlreadyUsed(email, data) {
				info.Status = "SIGNIN-ERROR"
			} else {
				// Set new user data and upload it
				data = info.Data.SetNewUserData(email, passw, username, data)
				info.Data.UploadUserData(data)
				info.Status = ""
			}
		}
	case "LOGIN", "LOGIN-ERROR":
		// If the status is LOGIN or LOGIN-ERROR, handle the actions accordingly
		if action == "Signin" {
			info.Status = "SIGNIN"
		} else {
			// Get user input values
			email := r.FormValue("email")
			passw := r.FormValue("password")

			// Check login credentials
			if !Log(email, passw, data) {
				info.Status = "LOGIN-ERROR"
			} else {
				// Set user data based on the login and reset status
				info.Data.SetUserData(email, data)
				info.Status = ""
			}
		}
	}

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles(
		"./pages/index.html",
		"./templates/game.html",
		"./templates/connexion.html",
		"./templates/footer.html",
		"./templates/hangman.html",
		"./templates/header.html",
		"./templates/board.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func headerHandler(w http.ResponseWriter, r *http.Request, infos *Structure) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}
	action := r.FormValue("action")
	switch action {
	case "ScoreBoard":
		infos.Status = "SCOREBOARD"
		infos.Board.SetScoreBoard(infos.Data.MakeBoard(3, ReadAllData()))
		fmt.Println(infos.Board)
	case "Play":
		if string(infos.Hangman.Word) == infos.Hangman.ToFind {
			infos.Status = "WIN"
		} else if infos.Hangman.EndGame() {
			infos.Status = "LOOSE"
		} else {
			infos.Status = ""
		}
	case "Profil":
		infos.Status = "PROFIL"
	}
	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func boardHandler(w http.ResponseWriter, r *http.Request, infos *Structure) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
	}
	action := r.FormValue("action")
	infos.Board.ChangeBoardId(action)
	allData := ReadAllData()
	if infos.Board.Id == 2 {
		infos.Board.SetScoreBoard(infos.Data.MakeRatioBoard(allData))
	} else {
		infos.Board.SetScoreBoard(infos.Data.MakeBoard(infos.Board.Id+3, allData))
	}
	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
