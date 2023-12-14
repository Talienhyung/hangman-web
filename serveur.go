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
	StrWord string
	Status  string
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html", "./templates/header.html", "./pages/info.html")
	if err != nil {
		log.Fatal(err)
	}
	infos.StrWord = string(infos.Hangman.Word)
	if infos.Hangman.IsThisTheWord(infos.StrWord) {
		infos.Status = "WIN"
	}
	template.Execute(w, infos)
}

func Login(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/login.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request, info Structure) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	letter := r.FormValue("input")

	if info.Hangman.MainMecanics(letter) {
		info.Status = "WIN"
	}
	info.StrWord = string(info.Hangman.Word)

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (myStruct *Structure) Init() {
	var infos HangManData
	infos.SetData()
	infos.SetWord(ReadAllDico())
	myStruct.Hangman = &infos
	myStruct.StrWord = string(infos.Word)
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
		hangmanHandler(w, r, myStruct)
	})

	http.HandleFunc("/login", Login)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
