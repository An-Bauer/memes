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

func TestDBFunc(t *testing.T) {
	InitDb()
	defer DB.Close()

	fmt.Println(CheckUserExistance(""))

}

func TestGetHash(t *testing.T) {
	InitDb()
	defer DB.Close()
	fmt.Println(GetHash("anton"))
}
