package server

import (
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	initServer()
}

func TestFile(t *testing.T) {

	//path := filepath.Join("e:", "InProgress", "memes", "web", "index.html")

	path := "E:/InProgress/memes/web/index.html"
	info, err := os.Stat(path)
	fmt.Println("jo")
	fmt.Println(info, err)

}
