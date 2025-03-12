package server

import (
	"fmt"
	"memes/code/db"
	"os"
	"testing"
)

func TestServer(t *testing.T) {
	db.InitDb()
	defer db.DB.Close()

	runServer()
}

func TestFile(t *testing.T) {

	//path := filepath.Join("e:", "InProgress", "memes", "web", "index.html")

	path := "E:/InProgress/memes/web/index.html"
	info, err := os.Stat(path)
	fmt.Println("jo")
	fmt.Println(info, err)

}
