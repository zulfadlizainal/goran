package nrMeasurement

import "math"

// SSB number of RE and symbol constant.
const (
	pss_re     int = 127 // re count
	sss_re     int = 127 // re count
	pbch_re    int = 576 // re count
	ssb_symbol int = 4   // symbol count
)

// SS-RSSI is the average received power of all SSB in each SSB symbols.
// SS-RSSI is calculated by dividing the total SSB RE power over total SSB symbols.
var ss_rssi float64

// SsRssi calculates SS-RSSI based on total SSB RE received power and SSB symbols.
//   - ss_rsrp refers to SS-RSRP value in dBm.
//   - The function will return SS-RSSI value in dBm.
func SsRssi(ss_rsrp float64) float64 {

	ss_rssi = float64(pss_re+sss_re+pbch_re) * math.Pow(10, (ss_rsrp/10)) / float64(ssb_symbol)

	return 10 * math.Log10(ss_rssi)

}
