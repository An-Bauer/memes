package qrcode

import "fmt"

//func findBestMask(qrcode [21][21]bool) int {
//sym := arrayToSymbol(ApplyMask(qrcode, 0))
//score := sym.penaltyScore()
//fmt.Println(score)
//lowest := 0

//for i := range 7 {
//sym := arrayToSymbol(ApplyMask(qrcode, i+1))
//newScore := sym.penaltyScore()
//fmt.Println(newScore)
//if newScore < score {
//score = newScore
//lowest = i + 1
//}
//}
//return lowest
//}

//func arrayToSymbol(qrcode [21][21]bool) symbol {
//sym := newSymbol(21, 0)

//for i := range 21 {
//for j := range 21 {
//sym.set(j, i, qrcode[i][j])
//}
//}

//return *sym
//}

//func symbolToArray(sym *symbol) [21][21]bool {
//arr := [21][21]bool{}

//for i := range 21 {
//for j := range 21 {
//arr[i][j] = sym.get(i, j)
//}
//}

//return arr
//}

func applyMask(qrcode Matrix, mask int) (Matrix, error) {
	dataMask := getDataMask()

	invertBit := func(i, j, mask int) (bool, error) {
		switch mask {
		case 0:
			return (i+j)%2 == 0, nil
		case 1:
			return i%2 == 0, nil
		case 2:
			return j%3 == 0, nil
		case 3:
			return (i+j)%3 == 0, nil
		case 4:
			return (i/2+j/3)%2 == 0, nil
		case 5:
			return (i*j)%2+(i*j)%3 == 0, nil
		case 6:
			return ((i*j)%2+((i*j)%3))%2 == 0, nil
		case 7:
			return ((i+j)%2+((i*j)%3))%2 == 0, nil
		}
		return false, fmt.Errorf("invalid mask (%d)", mask)
	}

	for i := range 21 {
		for j := range 21 {
			if dataMask[i][j] {
				invert, err := invertBit(i, j, mask)
				if err != nil {
					return Matrix{}, err
				}
				qrcode[i][j] = (qrcode[i][j] != invert)
			}
		}
	}

	return qrcode, nil
}

//func penalty1(qrcode [21][21]bool) int {
//score := 0

//for i := range 21 {
//count := 1
//streakColor := qrcode[i][0]

//for j := range 20 {
//color := qrcode[i][j+1]

//if color == streakColor {
//count += 1
//} else {
//count = 1
//streakColor = color
//}

//if count == 5 {
//score += 3
//} else if count > 5 {
//score += 1
//}
//}
//fmt.Println(score)
//}

//return score
//}
