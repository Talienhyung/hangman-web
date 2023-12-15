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
	Data    *h.Data
	StrWord string
	Status  string
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html", "./templates/header.html", "./pages/info.html", "./templates/connexion.html")
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
				info.Status = "LOOSE"
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
}

func (data *Structure) Reload() {
	data.Hangman.SetData()
	data.Hangman.SetWord(ReadAllDico())
	data.StrWord = string(data.Hangman.Word)
	data.StrWord = string(data.Hangman.Word)
	fmt.Println(data.Hangman.ToFind)
}

func relaodHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	action := r.FormValue("action")

	if action == "Reload" {
		info.Reload()
	}
	if info.Status == "WIN" {
		info.Data.Win += 1
		info.Data.Score += 1
		if info.Data.Score > info.Data.BestScore {
			info.Data.BestScore = info.Data.Score
		}
	} else if info.Status == "LOOSE" {
		info.Data.Loose += 1
		info.Data.Score = 0
	}
	info.Data.UploadUserData(h.ReadAllData())
	info.Status = ""

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func connexionHandler(w http.ResponseWriter, r *http.Request, info *Structure) {

	data := h.ReadAllData()
	action := r.FormValue("action")
	if info.Status == "CONNEXION" {
		switch action {
		case "Signin":
			info.Status = "SIGNIN"
		case "Login":
			info.Status = "LOGIN"
		}
	} else if info.Status == "SIGNIN" || info.Status == "SIGNIN-ERROR" {
		email := r.FormValue("email")
		username := r.FormValue("username")
		passw := r.FormValue("password")

		if h.EmailAlreadyUsed(email, data) {
			info.Status = "SIGNIN-ERROR"
		} else {
			data = info.Data.SetNewUserData(email, passw, username, data)
			info.Data.UploadUserData(data)
			info.Status = ""
		}
	} else if info.Status == "LOGIN" || info.Status == "LOGIN-ERROR" {
		email := r.FormValue("email")
		passw := r.FormValue("password")
		if !h.Log(email, passw, data) {
			info.Status = "LOGIN-ERROR"
		} else {
			info.Data.SetUserData(email, data)
			info.Status = ""
		}
	}
	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	var myStruct Structure
	var userData h.Data
	myStruct.Data = &userData
	myStruct.Init()
	myStruct.Status = "CONNEXION"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, myStruct)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanHandler(w, r, &myStruct)
	})

	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		relaodHandler(w, r, &myStruct)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		connexionHandler(w, r, &myStruct)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
