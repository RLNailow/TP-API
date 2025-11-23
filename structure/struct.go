package structure

// Album représente un album avec ses informations
type Album struct {
	Nom            string
	Image          string
	DateSortie     string
	NombreMusiques int
}

// PageDamso contient la liste des albums pour la page Damso
type PageDamso struct {
	Albums []Album
}

// Musique représente une musique avec ses informations
type Musique struct {
	Nom         string
	Artiste     string
	Album       string
	Image       string
	DateSortie  string
	LienSpotify string
}
