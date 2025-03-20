package qrcode

import (
	"fmt"
	"memes/code/encode"

	"github.com/skip2/go-qrcode/bitset"
	reedsolomon "github.com/skip2/go-qrcode/reedsolomon"
)

func Encode(text []rune, eccLevel EccLevel) ([]bool, error) {
	capacity, err := dataCapacity(eccLevel)
	if err != nil {
		return []bool{}, err
	}

	bits := make([]bool, 0, capacity*8)

	bits = append(bits, false, false, true, false)              // mode alphanumeric
	bits = append(bits, encode.IntToBoolSlice(len(text), 9)...) // count

	data, err := encode.EncodeChars(text) // data
	if err != nil {
		return []bool{}, err
	}
	bits = append(bits, data...)

	end, err := ending(len(bits), capacity*8) //ending
	if err != nil {
		return []bool{}, err
	}
	bits = append(bits, end...)

	bits = append(bits, zeros(len(bits))...)              // zeros
	bits = append(bits, pading(len(bits), capacity*8)...) // pading

	return reedsolomon.Encode(bitset.New(bits...), 26-capacity).Bits(), nil
}

func ending(length, capacity int) ([]bool, error) {
	if length > capacity {
		return []bool{}, fmt.Errorf("capacity (%d) to small for data (%d)", capacity, length)
	}

	return make([]bool, min(4, capacity-length)), nil
}

func zeros(length int) []bool {
	return make([]bool, 7-(length-1)%8)
}

func pading(length, capacity int) []bool {
	n := (capacity - length) / 8

	a := []bool{true, true, true, false, true, true, false, false}
	b := []bool{false, false, false, true, false, false, false, true}

	bits := make([]bool, 0)

	for i := range n {
		if i%2 == 0 {
			bits = append(bits, a...)
		} else {
			bits = append(bits, b...)
		}
	}

	return bits
}
