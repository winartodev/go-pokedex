package entity

// Attributes PokemonDB
type PokemonDB struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Species  string `db:"species"`
	Catched  int64  `db:"catched"`
	Metadata string `db:"metadata"`
}

// Attributes Pokemon
type Pokemon struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Types       []int64 `json:"types"`
	Catched     int64   `json:"catched"`
	ImageURL    string  `json:"image_url,omitempty"`
	Description string  `json:"description,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
	Height      float64 `json:"height,omitempty"`
	Stats       Stats   `json:"stats,omitempty"`
}

type PokemonDetail struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Species     string   `json:"species"`
	Types       []string `json:"types"`
	Catched     int64    `json:"catched"`
	ImageURL    string   `json:"image_url,omitempty"`
	Description string   `json:"description,omitempty"`
	Weight      float64  `json:"weight,omitempty"`
	Height      float64  `json:"height,omitempty"`
	Stats       Stats    `json:"stats,omitempty"`
}

// Attributes PokemonList
type PokemonList struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Types    []string `json:"types"`
	Catched  int64    `json:"catched"`
	ImageURL string   `json:"image_url"`
}

// Attributes Stats
type Stats struct {
	HP     int64 `json:"hp"`
	Attack int64 `json:"attack"`
	Def    int64 `json:"def"`
	Speed  int64 `json:"speed"`
}
