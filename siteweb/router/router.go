package router

import (
	"TpSpotify/siteweb/controller"
	"net/http"
)

// ConfigurerRoutes définit toutes les routes de l'application
func ConfigurerRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Route page d'accueil
	mux.HandleFunc("/", controller.AccueilHandler)

	// Route pour les albums de Damso
	mux.HandleFunc("/album/damso", controller.DamsoHandler)

	// Route pour la track Maladresse de Laylow
	mux.HandleFunc("/track/laylow", controller.MaladresseHandler)

	// Servir les fichiers CSS
	fs := http.FileServer(http.Dir("siteweb/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}
