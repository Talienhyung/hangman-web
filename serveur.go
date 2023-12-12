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
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html", "./templates/header.html", "./pages/info.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func Info(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/info.html", "./templates/footer.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	var infos HangManData
	infos.SetData()
	infos.SetWord(ReadAllDico())
	var myStruct Structure
	myStruct.Hangman = &infos
	myStruct.StrWord = string(infos.Word)
	fmt.Println(myStruct.StrWord)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, myStruct)
	})
	http.HandleFunc("/hangman", Info)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
