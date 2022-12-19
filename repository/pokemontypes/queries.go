package pokemontyperepository

const (
	InsertPokemonTypeQuery = `
		INSERT INTO pokedex.pokemon_types
		(
			pokemon_id,
			types_id
		) 
		VALUE 
		(
			?,
			?
		)
	`

	GetPokemonTypesByPokemonIDQuery = `
	SELECT
		pokemon_types.id,
		pokemon_types.pokemon_id,
		pokemon_types.types_id,
		types.name
	FROM pokedex.pokemon_types
	JOIN types ON types.id = pokemon_types.types_id
	WHERE pokemon_id = ?
	`

	UpdatePokemonTokenQuery = `
		UPDATE pokedex.pokemon_types
		SET 
			pokemon_id = ?,
			types_id = ?
		WHERE id = ? 
	`

	DeletePokemonTypeByPokemonIDQuery = `
		DELETE FROM pokedex.pokemon_types 
		WHERE pokemon_id = ?
	`

	DeletePokemonTypeByIDQuery = `
		DELETE FROM pokedex.pokemon_types 
		WHERE id = ?
	`
)
