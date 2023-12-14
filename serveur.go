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
	Status  string
	Win     int
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html", "./templates/header.html", "./pages/info.html")
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
		hangmanHandler(w, r, &myStruct)
	})

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
