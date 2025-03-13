package main

import (
	"memes/code/db"
	"memes/code/qrcode"
	"memes/code/server"
	"memes/code/svg"
	"testing"
)

func TestDraw(t *testing.T) {
	matrix, err := qrcode.CreateQR([]rune("google.de"), qrcode.Q, 6)
	if err != nil {
		t.Errorf("%v", err)
	}

	svg.DrawQR(matrix, "test.svg")
}

func TestServer(t *testing.T) {
	db.InitDb()
	defer db.DB.Close()

	server.RunServer()
}
