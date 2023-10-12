package nrConversion

// BandwidthRB represents data structures for bandwidth resource block properties.
type BandwidthRb struct {
	ID          int
	RbFr1Scs15  interface{}
	RbFr1Scs30  interface{}
	RbFr1Scs60  interface{}
	RbFr2Scs60  interface{}
	RbFr2Scs120 interface{}
}

// 3GPP TS 38.101-1 Table 5.3.2-1.
var BwRbTable = []BandwidthRb{
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

// Frequency ranges constant
const (
	Sub6   = "Sub6"
	MmWave = "mmWave"
)

// BandwidthToRB converts the Bandwidth (MHz) input to RB (Count) value based on the 3GPP TS 38.101-1 Table 5.3.2-1.
// The function will return -1 error if the input is out of range.
func BandwidthToRB(bandwidth int, ranges string, scs int) interface{} {
	for _, BandwidthRb := range BwRbTable {
		if BandwidthRb.ID == bandwidth {
			if ranges == Sub6 {
				if scs == 15 {
					return BandwidthRb.RbFr1Scs15
				} else if scs == 30 {
					return BandwidthRb.RbFr1Scs30
				} else if scs == 60 {
					return BandwidthRb.RbFr1Scs60
				}
			} else if ranges == MmWave {
				if scs == 60 {
					return BandwidthRb.RbFr2Scs60
				} else if scs == 120 {
					return BandwidthRb.RbFr2Scs120
				}
			}
		}
	}
	return -1
}
