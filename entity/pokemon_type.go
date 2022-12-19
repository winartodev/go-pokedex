package entity

// Attributes PokemonType
type PokemonType struct {
	ID        int64 `db:"id"`
	PokemonID int64 `db:"pokemon_id"`
	TypeID    int64 `db:"types_id"`
	Name      string
}
