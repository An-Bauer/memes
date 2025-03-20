package db

import "fmt"

func AddUser(username string, hash []byte) error {
	res, err := DB.Exec("INSERT INTO db.users (db.users.username,db.users.hash) VALUES (?,?);", username, hash)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows != 1 {
		return fmt.Errorf("update status did not update exactly one row (rows:%d)", nRows)
	}
	return nil
}

func GetHash(username string) ([]byte, error) {
	row := DB.QueryRow("SELECT db.users.hash FROM db.users WHERE db.users.username = ?", username)

	var hash []byte
	err := row.Scan(&hash)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

func UpdateToken(username, token string) error {
	res, err := DB.Exec("UPDATE db.users SET db.users.token = ? WHERE (db.users.username = ?);", token, username)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows != 1 {
		return fmt.Errorf("update status did not update exactly one row (rows:%d)", nRows)
	}
	return nil
}

func GetToken(username string) (string, error) {
	row := DB.QueryRow("SELECT db.users.token FROM db.users WHERE db.users.username = ?", username)

	var token string
	err := row.Scan(&token)

	if err != nil {
		return "", err
	}

	return token, nil
}
