package hangmanweb

import (
	"net/http"
)

// Root initializes the game structure and sets up the handlers
func Root() {
	var webData WebData
	var userData Data
	var board ScoreBoard
	webData.Data = &userData
	webData.Board = &board
	webData.Theme = "brown"
	// Initialize the game
	webData.Init()

	webData.Status = "CONNEXION"

	// Define the root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, webData)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanHandler(w, r, &webData)
	})
	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		levelHandler(w, r, &webData)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		connexionHandler(w, r, &webData)
	})
	http.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		disconnectHandler(w, r, &webData)
	})
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		headerHandler(w, r, &webData)
	})
	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		boardHandler(w, r, &webData)
	})
	http.HandleFunc("/theme", func(w http.ResponseWriter, r *http.Request) {
		themeHandler(w, r, &webData)
	})
}
