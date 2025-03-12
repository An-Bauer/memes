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

	fmt.Println("succesfully connected to db!")
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
