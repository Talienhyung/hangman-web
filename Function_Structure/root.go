package hangmanweb

import (
	"net/http"
)

// Root initializes the game structure and sets up the handlers
func Root() {
	var myStruct Structure
	var userData Data
	var board ScoreBoard
	myStruct.Data = &userData
	myStruct.Board = &board
	myStruct.Theme = "brown"
	// Initialize the game
	myStruct.Init()

	myStruct.Status = "CONNEXION"

	// Define the root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, myStruct)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanHandler(w, r, &myStruct)
	})
	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		levelHandler(w, r, &myStruct)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		connexionHandler(w, r, &myStruct)
	})
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		headerHandler(w, r, &myStruct)
	})
	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		boardHandler(w, r, &myStruct)
	})
	http.HandleFunc("/theme", func(w http.ResponseWriter, r *http.Request) {
		themeHandler(w, r, &myStruct)
	})

}
