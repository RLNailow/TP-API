package controller

import (
	"TpSpotify/structure"
	"html/template"
	"log"
	"net/http"
)

// AccueilHandler affiche la page d'accueil
func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/welcome.html")
	if err != nil {
		log.Println("Erreur chargement template welcome:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// DamsoHandler gère la page des albums de Damso
func DamsoHandler(w http.ResponseWriter, r *http.Request) {
	// Appeler la fonction qui récupère les albums de Damso
	albums, err := structure.RecupererAlbumsDamso()
	if err != nil {
		log.Println("Erreur récupération albums Damso:", err)
		http.Error(w, "Erreur lors de la récupération des albums", http.StatusInternalServerError)
		return
	}

	// Préparer les données pour le template
	data := structure.PageDamso{
		Albums: albums,
	}

	// Charger et exécuter le template
	tmpl, err := template.ParseFiles("templates/damso.html")
	if err != nil {
		log.Println("Erreur chargement template damso:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

// MaladresseHandler gère la page de la track Maladresse
func MaladresseHandler(w http.ResponseWriter, r *http.Request) {
	// Appeler la fonction qui récupère les infos de Maladresse
	musique, err := structure.RecupererMaladresse()
	if err != nil {
		log.Println("Erreur récupération Maladresse:", err)
		http.Error(w, "Erreur lors de la récupération de la musique", http.StatusInternalServerError)
		return
	}

	// Charger et exécuter le template
	tmpl, err := template.ParseFiles("templates/maladresse.html")
	if err != nil {
		log.Println("Erreur chargement template maladresse:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, musique)
}
