package nrDownlink

import (
	"math"
)

// ModulationBitsPerSymbol represents a map of modulation type (key) with its modulation order (value).
var ModulationBitsPerSymbol = map[string]int{
	"pi/2-BPSK": 1,  //bits/symbol
	"QPSK":      2,  //bits/symbol
	"16QAM":     4,  //bits/symbol
	"64QAM":     6,  //bits/symbol
	"256QAM":    8,  //bits/symbol
	"1024QAM":   10, //bits/symbol
}

// TbsTable is based on 3GPP TS 38.214 Table 5.1.3.2-1.
// Table contain TB size values (Bytes).
var TbsTable = []int{
	24, 32, 40, 48, 56, 64, 72, 80, 88, 96, 104, 112, 120, 128, 136, 144, 152, 160,
	168, 176, 184, 192, 208, 224, 240, 256, 272, 288, 304, 320, 336, 352, 368, 384,
	408, 432, 456, 480, 504, 528, 552, 576, 608, 640, 672, 704, 736, 768, 808, 848,
	888, 928, 984, 1032, 1064, 1128, 1160, 1192, 1224, 1256, 1288, 1320, 1352, 1416,
	1480, 1544, 1608, 1672, 1736, 1800, 1864, 1928, 2024, 2088, 2152, 2216, 2280,
	2408, 2472, 2536, 2600, 2664, 2728, 2792, 2856, 2976, 3104, 3240, 3368, 3496,
	3624, 3752, 3824,
}

// Tbs calculates Transport Block Size based on formula specified in 3GPP TS 38.214 - 5.1.3.2.
//   - mod refers to modulation type key in ModulationBitsPerSymbol map.
//   - nlayers refers to number of transmission layers from 1 to 8.
//   - nrb refers to number of RBs.
//   - nreperrb refers to number of REs per RB.
//   - tcr refers to target code rate between 0 to 1.
//   - xoh refers to additional overhead.
//   - tbscaling refers to scaling factor between 0 to 1.
//   - Tbs function will return TB size (Bytes).
func Tbs(mod string, nlayers int, nrb int, nreperrb int, tcr float64, xoh int, tbscaling float64) int {

	if _, ok := ModulationBitsPerSymbol[mod]; !ok {
		panic("out of range: mod")
	}

	if nlayers < 1 || nlayers > 8 {
		panic("out of range: nlayers")
	}

	if nrb < 0 {
		panic("out of range: nrb")
	}

	if nreperrb < 0 {
		panic("out of range: nreperrb")
	}

	if tcr < 0 || tcr > 1 {
		panic("out of range: tcr")
	}

	if xoh < 0 {
		panic("out of range: xoh")
	}

	if tbscaling <= 0 || tbscaling > 1 {
		panic("out of range: tbscaling")
	}

	nre := min(156, (nreperrb-xoh)) * nrb
	ninfo := tbscaling * float64(nre) * tcr * float64(ModulationBitsPerSymbol[mod]) * float64(nlayers)

	var tbs int

	if nreperrb == 0 || nrb == 0 || nreperrb < xoh {

		tbs = 0

	} else {
		if ninfo <= 3824 {
			n := max(3, math.Floor(math.Abs(math.Log2(ninfo)-6)))
			ninfo_ := max(24, (math.Pow(2, n))*math.Floor(math.Abs(ninfo/(math.Pow(2, n)))))

			target := int(ninfo_)
			var closestValue int
			minDiff := math.MaxInt
			for _, value := range TbsTable {
				diff := int(math.Abs(float64(target - value)))
				if diff < minDiff {
					minDiff = diff
					closestValue = value
				}
			}

			ninfo = float64(closestValue)

		} else {

			n := math.Floor(math.Abs(math.Log2(ninfo-24) - 5))
			ninfo_ := max(3840, (math.Pow(2, n))*math.Round((ninfo-24)/(math.Pow(2, n))))

			if tcr <= 0.25 {

				c := math.Ceil(math.Abs((ninfo_ + 24) / (3816)))
				ninfo = (8 * c * math.Ceil(math.Abs((ninfo_+24)/(8*c)))) - 24

			} else {

				if ninfo_ >= 8424 {

					c := math.Ceil(math.Abs((ninfo_ + 24) / (8424)))
					ninfo = (8 * c * math.Ceil(math.Abs((ninfo_+24)/(8*c)))) - 24

				} else {

					ninfo = (8 * math.Ceil(math.Abs((ninfo_+4)/8))) - 24

				}

			}

		}

	}

	tbs = int(ninfo)

	return tbs
}
