package pokemonrepository

const (
	GetPokemonQuery = `
		SELECT 
			pokemons.id, 
			pokemons.name, 
			pokemons.species, 
			pokemons.catched,
			pokemons.metadata
		FROM pokedex.pokemons
		JOIN pokedex.pokemon_types 
			ON pokemons.id = pokemon_types.pokemon_id
	`

	InsertPokemonQuery = `
		INSERT INTO pokedex.pokemons
		(
			name,
			species,
			catched,
			metadata
		) VALUES (
			?,
			?,
			?,
			?
		)
	`

	UpdatePokemonQuery = `
		UPDATE pokedex.pokemons 
		SET
			name = ?,
			species = ?,
			catched = ?,
			metadata = ?
		WHERE id = ?
	`

	DeletePokemonQuery = `
		DELETE FROM pokedex.pokemons 
		WHERE id = ?
	`
)
