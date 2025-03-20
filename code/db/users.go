package db

import "fmt"

func AddUser(username string, hash []byte) error {
	res, err := DB.Exec("INSERT INTO db.users (db.users.username,db.users.hash) VALUES (?,?);", username, hash)
	if err != nil {
		return fmt.Errorf("error while adding user A(err:%v, username:%s)", err, username)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while adding user B(err:%v, username:%s)", err, username)
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
		return nil, fmt.Errorf("error while getting hash (username:%s)", username)
	}

	return hash, nil
}

func UpdateToken(username, token string) error {
	res, err := DB.Exec("UPDATE db.users SET db.users.token = ? WHERE (db.users.username = ?);", token, username)
	if err != nil {
		return fmt.Errorf("error while updating token A(err:%v, username:%s)", err, username)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while updating token B(err:%v, username:%s)", err, username)
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
		return "", fmt.Errorf("error while getting token (err:%v, username:%s)", err, username)
	}

	return token, nil
}

func CheckUserExistance(username string) (bool, error) {
	row := DB.QueryRow("SELECT EXISTS(SELECT * FROM db.users WHERE db.users.username = ?)", username)

	var exists bool
	err := row.Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error while checking user existance (err:%v, username:%s)", err, username)
	}

	return exists, nil
}
