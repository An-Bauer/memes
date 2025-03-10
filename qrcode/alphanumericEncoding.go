package qrcode

import (
	"fmt"
)

func encodeChars(text []rune) ([]bool, error) {
	bits := make([]bool, 0)

	for i := range len(text) / 2 {
		n1, err := runeToInt(text[2*i])
		if err != nil {
			return []bool{}, err
		}

		n2, err := runeToInt(text[2*i+1])
		if err != nil {
			return []bool{}, err
		}

		n := 45*n1 + n2

		bits = append(bits, intToBoolSlice(n, 11)...)
	}

	if len(text)%2 == 1 {
		n, err := runeToInt(text[len(text)-1])
		if err != nil {
			return []bool{}, err
		}

		bits = append(bits, intToBoolSlice(n, 6)...)
	}

	return bits, nil
}

func intToBoolSlice(num, n int) []bool {
	bits := make([]bool, 0, n)

	for j := range n {
		bits = append(bits, (num/(1<<(n-1-j)))%2 == 1)
	}

	return bits
}

func runeToInt(r rune) (int, error) {
	switch r {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	case 'A':
		return 10, nil
	case 'B':
		return 11, nil
	case 'C':
		return 12, nil
	case 'D':
		return 13, nil
	case 'E':
		return 14, nil
	case 'F':
		return 15, nil
	case 'G':
		return 16, nil
	case 'H':
		return 17, nil
	case 'I':
		return 18, nil
	case 'J':
		return 19, nil
	case 'K':
		return 20, nil
	case 'L':
		return 21, nil
	case 'M':
		return 22, nil
	case 'N':
		return 23, nil
	case 'O':
		return 24, nil
	case 'P':
		return 25, nil
	case 'Q':
		return 26, nil
	case 'R':
		return 27, nil
	case 'S':
		return 28, nil
	case 'T':
		return 29, nil
	case 'U':
		return 30, nil
	case 'V':
		return 31, nil
	case 'W':
		return 32, nil
	case 'X':
		return 33, nil
	case 'Y':
		return 34, nil
	case 'Z':
		return 35, nil
	case ' ':
		return 36, nil
	case '$':
		return 37, nil
	case '%':
		return 38, nil
	case '*':
		return 39, nil
	case '+':
		return 40, nil
	case '-':
		return 41, nil
	case '.':
		return 42, nil
	case '/':
		return 43, nil
	case ':':
		return 44, nil
	}

	return 0, fmt.Errorf("illegal rune: %q", r)
}
