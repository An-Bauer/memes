package qrcode

import "fmt"

func dataCapacity(eccLevel EccLevel) (int, error) {
	switch eccLevel {
	case L:
		return 19, nil
	case M:
		return 16, nil
	case Q:
		return 13, nil
	case H:
		return 9, nil
	}
	return 0, fmt.Errorf("invalid eccLevel (%d)", eccLevel)
}

func formatBits(eccLevel EccLevel, mask int) ([]bool, error) {
	switch eccLevel {
	case L:
		switch mask {
		case 0:
			return []bool{true, true, true, false, true, true, true, true, true, false, false, false, true, false, false}, nil
		case 1:
			return []bool{true, true, true, false, false, true, false, true, true, true, true, false, false, true, true}, nil
		case 2:
			return []bool{true, true, true, true, true, false, true, true, false, true, false, true, false, true, false}, nil
		case 3:
			return []bool{true, true, true, true, false, false, false, true, false, false, true, true, true, false, true}, nil
		case 4:
			return []bool{true, true, false, false, true, true, false, false, false, true, false, true, true, true, true}, nil
		case 5:
			return []bool{true, true, false, false, false, true, true, false, false, false, true, true, false, false, false}, nil
		case 6:
			return []bool{true, true, false, true, true, false, false, false, true, false, false, false, false, false, true}, nil
		case 7:
			return []bool{true, true, false, true, false, false, true, false, true, true, true, false, true, true, false}, nil
		}
	case M:
		switch mask {
		case 0:
			return []bool{true, false, true, false, true, false, false, false, false, false, true, false, false, true, false}, nil
		case 1:
			return []bool{true, false, true, false, false, false, true, false, false, true, false, false, true, false, true}, nil
		case 2:
			return []bool{true, false, true, true, true, true, false, false, true, true, true, true, true, false, false}, nil
		case 3:
			return []bool{true, false, true, true, false, true, true, false, true, false, false, true, false, true, true}, nil
		case 4:
			return []bool{true, false, false, false, true, false, true, true, true, true, true, true, false, false, true}, nil
		case 5:
			return []bool{true, false, false, false, false, false, false, true, true, false, false, true, true, true, false}, nil
		case 6:
			return []bool{true, false, false, true, true, true, true, true, false, false, true, false, true, true, true}, nil
		case 7:
			return []bool{true, false, false, true, false, true, false, true, false, true, false, false, false, false, false}, nil
		}
	case Q:
		switch mask {
		case 0:
			return []bool{false, true, true, false, true, false, true, false, true, false, true, true, true, true, true}, nil
		case 1:
			return []bool{false, true, true, false, false, false, false, false, true, true, false, true, false, false, false}, nil
		case 2:
			return []bool{false, true, true, true, true, true, true, false, false, true, true, false, false, false, true}, nil
		case 3:
			return []bool{false, true, true, true, false, true, false, false, false, false, false, false, true, true, false}, nil
		case 4:
			return []bool{false, true, false, false, true, false, false, true, false, true, true, false, true, false, false}, nil
		case 5:
			return []bool{false, true, false, false, false, false, true, true, false, false, false, false, false, true, true}, nil
		case 6:
			return []bool{false, true, false, true, true, true, false, true, true, false, true, true, false, true, false}, nil
		case 7:
			return []bool{false, true, false, true, false, true, true, true, true, true, false, true, true, false, true}, nil
		}
	case H:
		switch mask {
		case 0:
			return []bool{false, false, true, false, true, true, false, true, false, false, false, true, false, false, true}, nil
		case 1:
			return []bool{false, false, true, false, false, true, true, true, false, true, true, true, true, true, false}, nil
		case 2:
			return []bool{false, false, true, true, true, false, false, true, true, true, false, false, true, true, true}, nil
		case 3:
			return []bool{false, false, true, true, false, false, true, true, true, false, true, false, false, false, false}, nil
		case 4:
			return []bool{false, false, false, false, true, true, true, false, true, true, false, false, false, true, false}, nil
		case 5:
			return []bool{false, false, false, false, false, true, false, false, true, false, true, false, true, false, true}, nil
		case 6:
			return []bool{false, false, false, true, true, false, true, false, false, false, false, true, true, false, false}, nil
		case 7:
			return []bool{false, false, false, true, false, false, false, false, false, true, true, true, false, true, true}, nil
		}
	}
	return []bool{}, fmt.Errorf("invalid eccLevel (%d) or mask (%d)", eccLevel, mask)
}
