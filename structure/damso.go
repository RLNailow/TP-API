package structure

import (
	"encoding/json"
	"fmt"
)

// RecupererAlbumsDamso récupère tous les albums de Damso depuis l'API Spotify
func RecupererAlbumsDamso() ([]Album, error) {
	// ID de Damso sur Spotify
	damsoID := "2UwqpfQtNuhBwviIC0f2ie"

	// Construire l'URL de l'API
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums?include_groups=album&market=FR&limit=50", damsoID)

	// Faire la requête via spotify.go
	body, err := RequeteSpotify(urlAPI)
	if err != nil {
		return nil, fmt.Errorf("erreur requête albums Damso: %w", err)
	}

	// Parser le JSON
	var resultat map[string]interface{}
	if err := json.Unmarshal(body, &resultat); err != nil {
		return nil, fmt.Errorf("erreur parsing JSON: %w", err)
	}

	// Extraire les albums
	items, ok := resultat["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("format de réponse invalide")
	}

	// Créer la liste des albums
	var albums []Album
	for _, item := range items {
		albumData, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Récupérer l'image
		imageURL := ""
		if images, ok := albumData["images"].([]interface{}); ok && len(images) > 0 {
			if firstImage, ok := images[0].(map[string]interface{}); ok {
				if url, ok := firstImage["url"].(string); ok {
					imageURL = url
				}
			}
		}

		// Créer l'album
		album := Album{
			Nom:            getString(albumData, "name"),
			Image:          imageURL,
			DateSortie:     getString(albumData, "release_date"),
			NombreMusiques: getInt(albumData, "total_tracks"),
		}

		albums = append(albums, album)
	}

	return albums, nil
}

// Fonctions utilitaires pour extraire des valeurs du JSON
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key].(string); ok {
		return val
	}
	return ""
}

func getInt(data map[string]interface{}, key string) int {
	if val, ok := data[key].(float64); ok {
		return int(val)
	}
	return 0
}
