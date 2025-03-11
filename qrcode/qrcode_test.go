package qrcode

import (
	"testing"
)

func TestQR(t *testing.T) {
	matrix, err := CreateQR([]rune("hello world"), Q, 6)
	if err != nil {
		t.Errorf("%v", err)
	}

	correctAnswer := Matrix{{true, true, true, true, true, true, true, false, false, false, false, true, false, false, true, true, true, true, true, true, true}, {true, false, false, false, false, false, true, false, true, true, false, false, true, false, true, false, false, false, false, false, true}, {true, false, true, true, true, false, true, false, false, true, false, true, true, false, true, false, true, true, true, false, true}, {true, false, true, true, true, false, true, false, true, true, true, true, true, false, true, false, true, true, true, false, true}, {true, false, true, true, true, false, true, false, true, true, false, true, false, false, true, false, true, true, true, false, true}, {true, false, false, false, false, false, true, false, false, true, false, false, true, false, true, false, false, false, false, false, true}, {true, true, true, true, true, true, true, false, true, false, true, false, true, false, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, true, true, false, true, true, false, false, false, false, false, false, false, false}, {false, true, false, true, true, true, true, false, true, true, false, false, true, true, true, false, true, true, false, true, false}, {true, false, true, true, true, true, false, true, false, false, false, false, true, true, true, true, false, true, true, true, false}, {false, false, true, false, true, false, true, true, false, false, false, true, false, false, true, true, false, false, false, false, false}, {true, false, true, true, false, true, false, false, false, true, false, true, true, false, false, false, true, true, false, false, false}, {true, true, false, true, true, true, true, true, true, true, true, false, true, true, true, false, true, true, true, true, true}, {false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, true, false, false, false}, {true, true, true, true, true, true, true, false, false, true, true, false, false, true, true, false, false, true, true, true, true}, {true, false, false, false, false, false, true, false, true, false, true, false, false, true, false, false, true, false, true, true, true}, {true, false, true, true, true, false, true, false, true, true, false, true, false, false, true, false, false, false, true, true, true}, {true, false, true, true, true, false, true, false, true, false, true, true, true, false, false, false, true, false, true, false, false}, {true, false, true, true, true, false, true, false, false, true, false, false, false, false, true, false, false, false, false, true, true}, {true, false, false, false, false, false, true, false, true, true, true, false, false, true, true, true, false, false, true, true, false}, {true, true, true, true, true, true, true, false, false, true, false, true, false, false, false, false, false, false, false, true, false}}

	if matrix != correctAnswer {
		t.Error("incorrect QRCode")
	}
}

//func printBits(bits []bool) {
//for i, v := range bits {
//if v {
//fmt.Print("1")
//} else {
//fmt.Print("0")
//}

//if i%8 == 7 {
//fmt.Print(" ")
//}
//}
//}
