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
	fmt.Println(DB)
	InitDb()
	fmt.Println(DB)
	defer DB.Close()

	insert()
}

func TestAvailibility(t *testing.T) {
	InitDb()
	defer DB.Close()

	av1, err1 := GetAvailability("meme0000")
	fmt.Println(av1, err1)

	av2, err2 := GetAvailability("abc012")
	fmt.Println(av2, err2)
}

func TestUpdateStatus(t *testing.T) {
	InitDb()
	defer DB.Close()

	fmt.Println(UpdateStatus("meme0000", 1))

}

func TestGetHash(t *testing.T) {
	InitDb()
	defer DB.Close()
	fmt.Println(GetHash("anton"))
}
