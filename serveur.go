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
	Status  string
	Level   string
}

func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles("./index.html", "./templates/game.html", "./templates/connexion.html", "./templates/footer.html", "./templates/hangman.html", "./templates/header.html", "./pages/info.html")
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

	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (myStruct *Structure) Init() {
	var infos HangManData
	infos.SetData()
	infos.SetWord(ReadAllDico())
	myStruct.Hangman = &infos
	fmt.Println(myStruct.Hangman.ToFind)
}

func (info *Structure) Reload(level string) {
	info.Hangman.SetData()
	if info.Status == "WIN" {
		switch info.Level {
		case "Easy":
			info.Data.Score += 1
			info.Data.WinEasy += 1
		case "Medium":
			info.Data.Score += 2
			info.Data.WinMedium += 1
		case "Hard":
			info.Data.Score += 3
			info.Data.WinHard += 1
		}

		info.Data.Win += 1
		if info.Data.Score > info.Data.BestScore {
			info.Data.BestScore = info.Data.Score
		}
	} else if info.Status == "LOOSE" {
		info.Data.Loose += 1
		info.Data.Score = 0
	}
	info.Data.UploadUserData(h.ReadAllData())
	info.Status = ""
}

func levelHandler(w http.ResponseWriter, r *http.Request, info *Structure) {
	action := r.FormValue("action")

	info.Reload(action)
	switch action {
	case "Easy":
		info.Hangman.SetWord(ReadTheDico("words.txt"))
		info.Level = "Easy"
	case "Medium":
		info.Hangman.SetWord(ReadTheDico("words2.txt"))
		info.Level = "Medium"
	case "Hard":
		info.Hangman.SetWord(ReadTheDico("words3.txt"))
		info.Level = "Hard"
	}
	fmt.Println(info.Hangman.ToFind)
	// Redirect back to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func connexionHandler(w http.ResponseWriter, r *http.Request, info *Structure) {

	data := h.ReadAllData()
	action := r.FormValue("action")
	switch info.Status {
	case "CONNEXION":
		switch action {
		case "Signin":
			info.Status = "SIGNIN"
		case "Login":
			info.Status = "LOGIN"
		}
	case "SIGNIN", "SIGNIN-ERROR":
		if action == "Login" {
			info.Status = "LOGIN"
		} else {
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
		}
	case "LOGIN", "LOGIN-ERROR":
		if action == "Signin" {
			info.Status = "SIGNIN"
		} else {
			email := r.FormValue("email")
			passw := r.FormValue("password")
			if !h.Log(email, passw, data) {
				info.Status = "LOGIN-ERROR"
			} else {
				info.Data.SetUserData(email, data)
				info.Status = ""
			}
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

	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		levelHandler(w, r, &myStruct)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		connexionHandler(w, r, &myStruct)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	fss := http.FileServer(http.Dir("Ressources/"))
	http.Handle("/Ressources/", http.StripPrefix("/Ressources", fss))
	http.ListenAndServe(":8080", nil)
}
