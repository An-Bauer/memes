package qrcode

import "fmt"

func CreateQR(content []rune, eccLevel EccLevel, mask int) (Matrix, error) {
	bits, err := encode(content, eccLevel)
	if err != nil {
		return Matrix{}, err
	}
	if len(bits) != 208 {
		return Matrix{}, fmt.Errorf("data has length %d, not 208", len(bits))
	}

	format, err := formatBits(eccLevel, mask)
	if err != nil {
		return Matrix{}, err
	}

	raw := pattern(bits, format, true, false)

	matrix, err := applyMask(raw, mask)
	if err != nil {
		return Matrix{}, err
	}

	return matrix, nil
}

func pattern(a []bool, m []bool, b, w bool) [21][21]bool {
	return ([21][21]bool{
		{b, b, b, b, b, b, b, w, m[14], a[137], a[136], a[135], a[134], w, b, b, b, b, b, b, b},
		{b, w, w, w, w, w, b, w, m[13], a[139], a[138], a[133], a[132], w, b, w, w, w, w, w, b},
		{b, w, b, b, b, w, b, w, m[12], a[141], a[140], a[131], a[130], w, b, w, b, b, b, w, b},
		{b, w, b, b, b, w, b, w, m[11], a[143], a[142], a[129], a[128], w, b, w, b, b, b, w, b},
		{b, w, b, b, b, w, b, w, m[10], a[145], a[144], a[127], a[126], w, b, w, b, b, b, w, b},
		{b, w, w, w, w, w, b, w, m[9], a[147], a[146], a[125], a[124], w, b, w, w, w, w, w, b},
		{b, b, b, b, b, b, b, w, b, w, b, w, b, w, b, b, b, b, b, b, b},
		{w, w, w, w, w, w, w, w, m[8], a[149], a[148], a[123], a[122], w, w, w, w, w, w, w, w},
		{m[0], m[1], m[2], m[3], m[4], m[5], b, m[6], m[7], a[151], a[150], a[121], a[120], m[7], m[8], m[9], m[10], m[11], m[12], m[13], m[14]},
		{a[201], a[200], a[199], a[198], a[185], a[184], w, a[183], a[182], a[153], a[152], a[119], a[118], a[73], a[72], a[71], a[70], a[25], a[24], a[23], a[22]},
		{a[203], a[202], a[197], a[196], a[187], a[186], b, a[181], a[180], a[155], a[154], a[117], a[116], a[75], a[74], a[69], a[68], a[27], a[26], a[21], a[20]},
		{a[205], a[204], a[195], a[194], a[189], a[188], w, a[179], a[178], a[157], a[156], a[115], a[114], a[77], a[76], a[67], a[66], a[29], a[28], a[19], a[18]},
		{a[207], a[206], a[193], a[192], a[191], a[190], b, a[177], a[176], a[159], a[158], a[113], a[112], a[79], a[78], a[65], a[64], a[31], a[30], a[17], a[16]},
		{w, w, w, w, w, w, w, w, b, a[161], a[160], a[111], a[110], a[81], a[80], a[63], a[62], a[33], a[32], a[15], a[14]},
		{b, b, b, b, b, b, b, w, m[6], a[163], a[162], a[109], a[108], a[83], a[82], a[61], a[60], a[35], a[34], a[13], a[12]},
		{b, w, w, w, w, w, b, w, m[5], a[165], a[164], a[107], a[106], a[85], a[84], a[59], a[58], a[37], a[36], a[11], a[10]},
		{b, w, b, b, b, w, b, w, m[4], a[167], a[166], a[105], a[104], a[87], a[86], a[57], a[56], a[39], a[38], a[9], a[8]},
		{b, w, b, b, b, w, b, w, m[3], a[169], a[168], a[103], a[102], a[89], a[88], a[55], a[54], a[41], a[40], a[7], a[6]},
		{b, w, b, b, b, w, b, w, m[2], a[171], a[170], a[101], a[100], a[91], a[90], a[53], a[52], a[43], a[42], a[5], a[4]},
		{b, w, w, w, w, w, b, w, m[1], a[173], a[172], a[99], a[98], a[93], a[92], a[51], a[50], a[45], a[44], a[3], a[2]},
		{b, b, b, b, b, b, b, w, m[0], a[175], a[174], a[97], a[96], a[95], a[94], a[49], a[48], a[47], a[46], a[1], a[0]},
	})
}

func getDataMask() Matrix {
	a := make([]bool, 208)
	for i := range a {
		a[i] = true
	}

	m := make([]bool, 15)

	return pattern(a, m, false, false)
}
