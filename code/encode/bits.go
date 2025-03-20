package encode

func EncodeChars(text []rune) ([]bool, error) {
	bits := make([]bool, 0)

	for i := range len(text) / 2 {
		n1, err := RuneToInt(text[2*i])
		if err != nil {
			return []bool{}, err
		}

		n2, err := RuneToInt(text[2*i+1])
		if err != nil {
			return []bool{}, err
		}

		n := 45*n1 + n2

		bits = append(bits, IntToBoolSlice(n, 11)...)
	}

	if len(text)%2 == 1 {
		n, err := RuneToInt(text[len(text)-1])
		if err != nil {
			return []bool{}, err
		}

		bits = append(bits, IntToBoolSlice(n, 6)...)
	}

	return bits, nil
}

func IntToBoolSlice(num, n int) []bool {
	bits := make([]bool, 0, n)

	for j := range n {
		bits = append(bits, (num/(1<<(n-1-j)))%2 == 1)
	}

	return bits
}
