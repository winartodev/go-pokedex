package userrepository

const (
	CreateTableUsersQuery = `
		CREATE TABLE IF NOT EXISTS  users (
			id int not null auto_increment,
			username varchar(255) not null,
			email varchar(255) not null,
			password text not null, 
			role int not null,
			primary key (id)
		);
	`

	InsertUserQuery = `
		INSERT INTO pokedex.users
			(
				username,
				email, 
				password,
				role
			) 
		VALUES
		(
			?, 
			?,
			?,
			?
		)
	`

	GetUserQuery = `
		SELECT 
			id,
			username,
			email,
			password,
			role
		FROM pokedex.users 
	`
)
