package typesrepository

const (
	GetTypesQuery = `
		SELECT
			id,
			name
		FROM pokedex.types
	`

	InsertTypeQuery = `
		INSERT INTO pokedex.types 
		(
			name
		) 
		VALUES 
		(
			?
		)
	`

	UpdateTypeQuery = `
		UPDATE pokedex.types  
		SET 
			name = ?
		WHERE id = ? 
	`
)
