package main

import (
	"memes/code/qrcode"
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
