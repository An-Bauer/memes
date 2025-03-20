package db

import "fmt"

func GetStatus(key string) (int, error) {
	row := DB.QueryRow("SELECT status FROM db.memes WHERE db.memes.key = ?", key)

	var status int
	err := row.Scan(&status)

	if err != nil {
		return 0, err
	}

	return status, nil
}

func UpdateStatus(key string, status int) error {
	res, err := DB.Exec("UPDATE db.memes SET db.memes.status = ? WHERE (db.memes.key = ?);", status, key)
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
