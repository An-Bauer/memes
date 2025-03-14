package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDb() {
	//pswd := os.Getenv("MYSQL_PASSWORD")
	pswd := "rootPassword" // dev only!

	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/db")
	DB = db

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	fmt.Println("LOG: succesfully connected to db")
}

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

func GetAvailability(key string) (bool, error) {
	row := DB.QueryRow("SELECT available FROM db.memes WHERE db.memes.key = ?", key)

	var available bool
	err := row.Scan(&available)

	if err != nil {
		return false, err
	}

	return available, nil
}

func insert() {
	key := "abceig"
	batch := 1

	_, err := DB.Exec(" INSERT INTO db.memes (db.memes.key,db.memes.batch,db.memes.date,db.memes.available) VALUES (?,?,NOW(),1);", key, batch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("jo")
}
