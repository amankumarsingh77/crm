package repository

const (
	createUser = `INSERT INTO users (first_name, last_name, email, password_hash, role, phone , is_active,
	               		 created_at, updated_at)
						VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6, $7, now(), now()) 
						RETURNING *`
	updateUser = `UPDATE users 
						SET first_name = COALESCE(NULLIF($1, ''), first_name),
						    last_name = COALESCE(NULLIF($2, ''), last_name),
						    email = COALESCE(NULLIF($3, ''), email),
						    role = COALESCE(NULLIF($4, ''), role),
						    phone = COALESCE(NULLIF($5, ''), phone),
						    is_active = COALESCE(NULLIF($6, ''), is_active),
						    updated_at = now()
						WHERE user_id = $7
						RETURNING *
				`
	deleteUserQuery = `DELETE FROM users WHERE user_id = $1`

	getUserQuery = `SELECT id, first_name, last_name, email, role, phone, created_at, updated_at, last_login  
					 FROM users 
					 WHERE id = $1`
	getUserByEmail = `SELECT id , email, password_hash, first_name, last_name, phone, role, is_active, last_login, created_at, updated_at
						FROM users WHERE email = $1`
	getTotalCount = "SELECT COUNT(id) FROM users WHERE first_name ILIKE '%' || $1 || '%' or last_name ILIKE '%' || $1 || '%' "
	getTotal      = `SELECT COUNT(id) FROM users`
	getUsers      = `SELECT id, first_name, last_name, email, role, phone, created_at, updated_at, last_login
				 	 FROM users 
				 	 ORDER BY COALESCE(NULLIF($1, ''), first_name) OFFSET $2 LIMIT $3`
)
