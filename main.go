package main

import (
	"fmt"
	"os"
)

func main() {
	pswd := os.Getenv("MYSQL_PASSWORD")
	fmt.Println(pswd)
	fmt.Println("läuft")
}
