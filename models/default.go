package models

type PokemonDefaultResponse struct {
	Name          string `json:"name"`
	PokedexNumber int    `json:"pokedex_number"`
	Type1         string `json:"type1"`
	Generation    int    `json:"generation"`
}
