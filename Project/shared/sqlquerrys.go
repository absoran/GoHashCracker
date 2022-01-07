package shared

// sql querrys for not repeat myself.
var SqlInsert_user = `
	INSERT INTO app_user (username, email, password)
	VALUES ($1, $2, $3)
	RETURNING id`

var SqlGetUserByUserName = `SELECT * FROM app_user WHERE username=`

var SqlCheckExist = `select exists(select 1 from app_user where username='`
