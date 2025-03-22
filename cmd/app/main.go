package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	for range 100 {
		fmt.Println(rand.Text())
	}
}
