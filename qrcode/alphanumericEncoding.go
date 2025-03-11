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
	case 'a':
		return 10, nil
	case 'b':
		return 11, nil
	case 'c':
		return 12, nil
	case 'd':
		return 13, nil
	case 'e':
		return 14, nil
	case 'f':
		return 15, nil
	case 'g':
		return 16, nil
	case 'h':
		return 17, nil
	case 'i':
		return 18, nil
	case 'j':
		return 19, nil
	case 'k':
		return 20, nil
	case 'l':
		return 21, nil
	case 'm':
		return 22, nil
	case 'n':
		return 23, nil
	case 'o':
		return 24, nil
	case 'p':
		return 25, nil
	case 'q':
		return 26, nil
	case 'r':
		return 27, nil
	case 's':
		return 28, nil
	case 't':
		return 29, nil
	case 'u':
		return 30, nil
	case 'v':
		return 31, nil
	case 'w':
		return 32, nil
	case 'x':
		return 33, nil
	case 'y':
		return 34, nil
	case 'z':
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
