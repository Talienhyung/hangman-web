package hangmanweb

import (
	"net/http"
)

// Root initializes the game structure and sets up the handlers
func Root() {
	var myStruct Structure
	var userData Data
	myStruct.Data = &userData

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
}
