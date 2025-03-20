package encode

import (
	"fmt"
)

func IntToRune(n int64) (rune, error) {
	switch n {
	case 0:
		return '0', nil
	case 1:
		return '1', nil
	case 2:
		return '2', nil
	case 3:
		return '3', nil
	case 4:
		return '4', nil
	case 5:
		return '5', nil
	case 6:
		return '6', nil
	case 7:
		return '7', nil
	case 8:
		return '8', nil
	case 9:
		return '9', nil
	case 10:
		return 'a', nil
	case 11:
		return 'b', nil
	case 12:
		return 'c', nil
	case 13:
		return 'd', nil
	case 14:
		return 'e', nil
	case 15:
		return 'f', nil
	case 16:
		return 'g', nil
	case 17:
		return 'h', nil
	case 18:
		return 'i', nil
	case 19:
		return 'j', nil
	case 20:
		return 'k', nil
	case 21:
		return 'l', nil
	case 22:
		return 'm', nil
	case 23:
		return 'n', nil
	case 24:
		return 'o', nil
	case 25:
		return 'p', nil
	case 26:
		return 'q', nil
	case 27:
		return 'r', nil
	case 28:
		return 's', nil
	case 29:
		return 't', nil
	case 30:
		return 'u', nil
	case 31:
		return 'v', nil
	case 32:
		return 'w', nil
	case 33:
		return 'x', nil
	case 34:
		return 'y', nil
	case 35:
		return 'z', nil
	case 36:
		return ' ', nil
	case 37:
		return '$', nil
	case 38:
		return '%', nil
	case 39:
		return '*', nil
	case 40:
		return '+', nil
	case 41:
		return '-', nil
	case 42:
		return '.', nil
	case 43:
		return '/', nil
	case 44:
		return ':', nil
	default:
		return '0', fmt.Errorf("can't convert %d to rune", n)
	}
}
