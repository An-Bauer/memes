package db

import (
	"fmt"
	"os"
	"testing"
)

func TestPswd(t *testing.T) {
	pswd := os.Getenv("MYSQL_PASSWORD")
	fmt.Println(pswd)
	fmt.Println("test")
}

func TestInit(t *testing.T) {
	db := initDb()
	defer db.Close()

	insert(db)
}

func TestAvailibility(t *testing.T) {
	db := initDb()
	defer db.Close()

	av1, err1 := getAvailability(db, "meme0000")
	fmt.Println(av1, err1)

	av2, err2 := getAvailability(db, "abc012")
	fmt.Println(av2, err2)
}
