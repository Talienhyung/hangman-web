package main

import (
	"fmt"
	h "hangmanweb/BDD"
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
	testBDD()
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

func testBDD() {
	var user h.Data
	data := h.ReadAllData()
	if !h.Log("tan@gmail.com", "jesuistan", data) {
		fmt.Println("Pas de compte")
		data = user.SetNewUserData("tan@gmail.com", "jesuistan", "Tan", data)
	} else {
		user.SetUserData("tan@gmail.com", data)
		user.Win += 1
	}
	user.UploadUserData(data)
	fmt.Println(user.Win)
}
