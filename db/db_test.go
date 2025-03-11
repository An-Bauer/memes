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
