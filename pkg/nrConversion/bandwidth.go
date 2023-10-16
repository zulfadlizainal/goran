package nrConversion

import "errors"

// BandwidthRb represents data structures for Bandwidth table with its RB count.
// The data structures associated to variable BandwidthRbTable.
type BandwidthRb struct {
	Bw          int         // MHz
	RbFr1Scs15  interface{} // rb count fr1 scs 15
	RbFr1Scs30  interface{} // rb count fr1 scs 30
	RbFr1Scs60  interface{} // rb count fr1 scs 60
	RbFr2Scs60  interface{} // rb count fr2 scs 60
	RbFr2Scs120 interface{} // rb count fr2 scs 120
}

// BandwidthRbTable is based on 3GPP TS 38.101-1 Table 5.3.2-1.
// The table associated to type BandwidthRb.
var BandwidthRbTable = []BandwidthRb{
	{5, 25, 11, nil, nil, nil},
	{10, 52, 24, 11, nil, nil},
	{15, 79, 38, 18, nil, nil},
	{20, 106, 51, 24, nil, nil},
	{25, 133, 65, 31, nil, nil},
	{30, 160, 78, 38, nil, nil},
	{40, 216, 106, 51, nil, nil},
	{50, 270, 133, 65, 66, 32},
	{60, nil, 162, 79, nil, nil},
	{80, nil, 217, 107, nil, nil},
	{100, nil, 273, 135, 132, 66},
	{200, nil, nil, nil, 264, 132},
	{400, nil, nil, nil, nil, 264},
}

// Frequency Ranges constant.
const (
	Sub6   = "Sub6"   // FR1
	MmWave = "mmWave" // FR2
)

// Sub Carrier Spacing constant.
const (
	Scs15  = 15  // kHz
	Scs30  = 30  // kHz
	Scs60  = 60  // kHz
	Scs120 = 120 // kHz
)

// BandwidthToRB converts the Bandwidth to RB count based on BandwidthRbTable.
//   - bandwidth refers to Bw within the range of BandwidthRbTable.
//   - fr refers to Frequency Ranges const.
//   - scs refers to Sub Carrier Spacing const.
//   - The function will return RB count within the range of BandwidthRbTable.
func BandwidthToRB(bandwidth int, fr string, scs int) (interface{}, error) {
	for _, BandwidthRb := range BandwidthRbTable {
		if BandwidthRb.Bw == bandwidth {
			if fr == Sub6 {
				if scs == Scs15 {
					return BandwidthRb.RbFr1Scs15, nil
				} else if scs == Scs30 {
					return BandwidthRb.RbFr1Scs30, nil
				} else if scs == Scs60 {
					return BandwidthRb.RbFr1Scs60, nil
				}
			} else if fr == MmWave {
				if scs == Scs60 {
					return BandwidthRb.RbFr2Scs60, nil
				} else if scs == Scs120 {
					return BandwidthRb.RbFr2Scs120, nil
				}
			}
		}
	}
	return nil, errors.New("out of range")
}
