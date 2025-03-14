package main

import (
	"memes/code/db"
	"memes/code/qrcode"
	"memes/code/server"
	"memes/code/svg"
	"testing"
)

func TestDraw(t *testing.T) {
	matrix, err := qrcode.CreateQR([]rune("192.168.178.53:8/meme0002"), qrcode.L, 6)
	if err != nil {
		t.Errorf("%v", err)
	}

	svg.DrawQR(matrix, "meme0002.svg")
}

func TestServer(t *testing.T) {
	db.InitDb()
	defer db.DB.Close()

	server.RunServer()
}
