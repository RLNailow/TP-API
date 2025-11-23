package main

import (
	"TpSpotify/siteweb/router"
	"TpSpotify/structure"
	"log"
	"net/http"
)

func main() {
	// Initialiser le token Spotify au démarrage
	err := structure.InitialiserToken()
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation du token:", err)
	}

	log.Println("✅ Token Spotify initialisé avec succès")

	// Configurer les routes
	r := router.ConfigurerRoutes()

	// Lancer le serveur sur le port 8080
	log.Println("🚀 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
