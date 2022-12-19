package typesrepository

const (
	CreateTableTypesQuery = `
		CREATE TABLE IF NOT EXISTS types (
			id int not null auto_increment,
			name varchar(255) not null,
			primary key (id)
		);
	`

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
