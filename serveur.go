package main

import (
	hangmanweb "hangmanweb/Function_Structure"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	// Obtenez le répertoire actuel du programme
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Créez des chemins absolus pour les répertoires statiques
	staticDir := filepath.Join(dir, "static")
	ressourcesDir := filepath.Join(dir, "Ressources")

	// Initialisez un routeur Gorilla mux
	router := mux.NewRouter()

	// Ajoutez votre gestionnaire de routes hangmanweb.Root() ici si nécessaire
	hangmanweb.Root(router)

	// Serveur de fichiers statiques
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	router.PathPrefix("/Ressources/").Handler(http.StripPrefix("/Ressources/", http.FileServer(http.Dir(ressourcesDir))))

	// Démarrez le serveur en utilisant l'environnement PORT ou le port 8080 par défaut
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
