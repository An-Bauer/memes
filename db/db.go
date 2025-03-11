package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func initDb() *sql.DB {
	//pswd := os.Getenv("MYSQL_PASSWORD")
	pswd := "rootPassword" // dev only!

	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/db")
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	fmt.Println("succesfully connected to db!")

	return db
}

func insert(db *sql.DB) {
	key := "abceih"
	batch := 1

	_, err := db.Exec(" INSERT INTO db.memes (db.memes.key,db.memes.batch,db.memes.date) VALUES (?,?,NOW());", key, batch)
	if err != nil {
		fmt.Println("mist!")
		fmt.Println(err)
	}
}
