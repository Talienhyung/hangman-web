package hangmanweb

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Root initializes the game structure and sets up the handlers
func Root(router *mux.Router) {
	var webData WebData
	var userData Data
	var board ScoreBoard
	webData.Data = &userData
	webData.Board = &board
	webData.Theme = "green"
	// Initialize the game
	webData.Init()

	webData.Status = "CONNEXION"

	// Define the root handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, webData)
	})
	router.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanHandler(w, r, &webData)
	})
	router.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		levelHandler(w, r, &webData)
	})
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		connexionHandler(w, r, &webData)
	})
	router.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		disconnectHandler(w, r, &webData)
	})
	router.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		headerHandler(w, r, &webData)
	})
	router.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		boardHandler(w, r, &webData)
	})
	router.HandleFunc("/theme", func(w http.ResponseWriter, r *http.Request) {
		themeHandler(w, r, &webData)
	})
}
