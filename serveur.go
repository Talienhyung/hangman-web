package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	. "github.com/Talienhyung/hangman"
)

type Structure struct {
	Hangman *HangManData
	Status  string
	Win     int
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/game.html", "./templates/footer.html", "./templates/header.html", "./pages/info.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	if r.Method != http.MethodPost {
		info.Init()
	}

	letter := r.FormValue("input")
	if info.Hangman.UsedVerif(letter) {
		info.Status = "USED"
	} else {
		if info.Hangman.MainMecanics(letter) {
			info.Status = "WIN"
		} else if info.Hangman.EndGame() {
			if string(info.Hangman.Word) == info.Hangman.ToFind {
				info.Status = "WIN"
			} else {
				info.Status = "LOSE"
			}
		} else {
			info.Status = ""
		}
	}

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (myStruct *Structure) Init() {
	var infos HangManData
	infos.SetData()
	infos.SetWord(ReadAllDico())
	myStruct.Hangman = &infos
	fmt.Println(myStruct.Hangman.ToFind)
	myStruct.Status = ""
}

func main() {
	var myStruct Structure
	myStruct.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, myStruct)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanHandler(w, r, &myStruct)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	fss := http.FileServer(http.Dir("Ressources/"))
	http.Handle("/Ressources/", http.StripPrefix("/Ressources", fss))
	http.ListenAndServe(":8080", nil)
}
