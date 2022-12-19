package pokemonrepository

const (
	CreateTablePokemonQuery = `
		CREATE TABLE IF NOT EXISTS pokemons (
			id int not null auto_increment,
			name varchar(255) not null,
			species varchar(255) not null,
			catched int not null,
			metadata text,
			primary key (id)
		);
	`

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
