package main

import (
	"memes/qrcode"
	"memes/svg"
	"testing"
)

func TestDraw(t *testing.T) {
	matrix, err := qrcode.CreateQR([]rune("HELLO WORLD"), qrcode.Q, 6)
	if err != nil {
		t.Errorf("%v", err)
	}

	svg.DrawQR(matrix, "test.svg")
}
