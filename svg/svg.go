package svg

import (
	"os"

	svgLib "github.com/ajstarks/svgo"
)

func DrawQR(qrcode [21][21]bool, name string) {

	getColor := func(i, j int) bool {
		if i < 0 || i > 20 || j < 0 || j > 20 {
			return false //white
		}
		return qrcode[i][j]
	}

	cell := 10
	corner := 4

	file, err := os.Create(name)
	if err != nil {
		panic("fuck")
	}
	defer file.Close()
	canvas := svgLib.New(file)
	canvas.Start(21*cell, 21*cell)

	for i := range 22 {
		for j := range 22 {
			nw, no, sw, so := getColor(j-1, i-1), getColor(j, i-1), getColor(j-1, i), getColor(j, i)

			switch {
			case (nw && no && so) || (no && so && sw) || (so && sw && nw) || (sw && nw && no):
				//case (nw && so) || (no && sw):
				canvas.Rect(cell/2*(2*i-1), cell/2*(2*j-1), cell, cell, "fill:black; stroke:none") //voll
			case (!nw && !no && sw && so):
				canvas.Rect(cell/2*(2*i), cell/2*(2*j-1), cell/2, cell, "fill:black; stroke:none") //unten
			case (nw && !no && sw && !so):
				canvas.Rect(cell/2*(2*i-1), cell/2*(2*j-1), cell, cell/2, "fill:black; stroke:none") //links
			case (nw && no && !sw && !so):
				canvas.Rect(cell/2*(2*i-1), cell/2*(2*j-1), cell/2, cell, "fill:black; stroke:none") //oben
			case (!nw && no && !sw && so):
				canvas.Rect(cell/2*(2*i-1), cell/2*(2*j), cell, cell/2, "fill:black; stroke:none") //rechts
			}
		}
	}

	for i := range 21 {
		for j := range 21 {
			if getColor(j, i) {
				canvas.Roundrect(cell*i, cell*j, cell, cell, corner, corner, "fill:black; stroke:none")
			} else {
				canvas.Roundrect(cell*i, cell*j, cell, cell, corner, corner, "fill:white; stroke:none")
			}
		}
	}

	canvas.End()
}
