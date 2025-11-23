package structure

import (
	"encoding/json"
	"fmt"
)

// RecupererMaladresse récupère les informations de la track "Maladresse" de Laylow
func RecupererMaladresse() (Musique, error) {
	// ID de la track "Maladresse" sur Spotify
	trackID := "3C8fNulR5KxiOXD3Ql7zJo"

	// Construire l'URL de l'API
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", trackID)

	// Faire la requête via spotify.go
	body, err := RequeteSpotify(urlAPI)
	if err != nil {
		return Musique{}, fmt.Errorf("erreur requête Maladresse: %w", err)
	}

	// Parser le JSON
	var resultat map[string]interface{}
	if err := json.Unmarshal(body, &resultat); err != nil {
		return Musique{}, fmt.Errorf("erreur parsing JSON: %w", err)
	}

	// Extraire le nom de la musique
	nomMusique := getString(resultat, "name")

	// Extraire le lien Spotify
	lienSpotify := ""
	if externalURLs, ok := resultat["external_urls"].(map[string]interface{}); ok {
		lienSpotify = getString(externalURLs, "spotify")
	}

	// Extraire l'artiste
	nomArtiste := ""
	if artists, ok := resultat["artists"].([]interface{}); ok && len(artists) > 0 {
		if firstArtist, ok := artists[0].(map[string]interface{}); ok {
			nomArtiste = getString(firstArtist, "name")
		}
	}

	// Extraire les infos de l'album
	nomAlbum := ""
	dateSortie := ""
	imageURL := ""

	if albumData, ok := resultat["album"].(map[string]interface{}); ok {
		nomAlbum = getString(albumData, "name")
		dateSortie = getString(albumData, "release_date")

		// Extraire l'image de l'album
		if images, ok := albumData["images"].([]interface{}); ok && len(images) > 0 {
			if firstImage, ok := images[0].(map[string]interface{}); ok {
				imageURL = getString(firstImage, "url")
			}
		}
	}

	// Créer la structure Musique
	musique := Musique{
		Nom:         nomMusique,
		Artiste:     nomArtiste,
		Album:       nomAlbum,
		Image:       imageURL,
		DateSortie:  dateSortie,
		LienSpotify: lienSpotify,
	}

	return musique, nil
}
