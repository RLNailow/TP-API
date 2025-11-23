package structure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Variables globales pour le token
var (
	tokenSpotify    string
	expirationToken time.Time
)

// credentials Spotify
const (
	CLIENT_ID     = "be57ff2fae44482a86e19a72d2ec2d66"
	CLIENT_SECRET = "162ec34616394ab99786217a8f7d65e4"
)

// InitialiserToken obtient le premier token au démarrage
func InitialiserToken() error {
	return obtenirNouveauToken()
}

// obtenirNouveauToken récupère un nouveau token depuis l'API Spotify
func obtenirNouveauToken() error {
	// Préparer les données de la requête
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	// Créer la requête POST
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("erreur création requête: %w", err)
	}

	// Ajouter l'authentification Basic
	req.SetBasicAuth(CLIENT_ID, CLIENT_SECRET)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Envoyer la requête
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur envoi requête: %w", err)
	}
	defer resp.Body.Close()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lecture réponse: %w", err)
	}

	// Vérifier le status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API status %d: %s", resp.StatusCode, string(body))
	}

	// Parser le JSON
	var resultat map[string]interface{}
	if err := json.Unmarshal(body, &resultat); err != nil {
		return fmt.Errorf("erreur parsing JSON: %w", err)
	}

	// Extraire le token et la durée d'expiration
	tokenSpotify = resultat["access_token"].(string)
	expiresIn := int(resultat["expires_in"].(float64))

	// Définir l'expiration (on rafraîchit 5 minutes avant pour être sûr)
	expirationToken = time.Now().Add(time.Duration(expiresIn-300) * time.Second)

	return nil
}

// verifierToken vérifie si le token est encore valide, sinon le rafraîchit
func verifierToken() error {
	if time.Now().After(expirationToken) {
		return obtenirNouveauToken()
	}
	return nil
}

// RequeteSpotify effectue une requête GET vers l'API Spotify
func RequeteSpotify(urlAPI string) ([]byte, error) {
	// Vérifier et rafraîchir le token si nécessaire
	if err := verifierToken(); err != nil {
		return nil, fmt.Errorf("erreur vérification token: %w", err)
	}

	// Créer la requête GET
	req, err := http.NewRequest("GET", urlAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("erreur création requête: %w", err)
	}

	// Ajouter le token dans le header
	req.Header.Set("Authorization", "Bearer "+tokenSpotify)

	// Envoyer la requête
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erreur envoi requête: %w", err)
	}
	defer resp.Body.Close()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lecture réponse: %w", err)
	}

	// Vérifier le status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
