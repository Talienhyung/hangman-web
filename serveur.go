package main

import (
	hangmanweb "hangmanweb/Function_Structure"
	"net/http"
)

func main() {
	hangmanweb.Root()
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	fss := http.FileServer(http.Dir("Ressources/"))
	http.Handle("/Ressources/", http.StripPrefix("/Ressources", fss))
	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
