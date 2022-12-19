package userrepository

const (
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
