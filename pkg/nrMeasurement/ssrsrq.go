package nrMeasurement

import "math"

// Total SSB RB count constant (20 RB).
const ssbrbcount int = 20

// SS-RSRQ is defined as SS-RSRP / (RSSI / N) where N is the number SSB RB.
var ss_rsrq float64

// SsRsrq calculates SS-RSRQ based on SS-RSRP measurement, RSSI measurement, and total SSB RB count.
//   - ss_rsrp refers to SS-RSRP in dBm.
//   - rssi refers to RSSI within the total SSB RB count in dBm.
//   - The function will return SS-RSRQ value in dB.
func SsRsrq(ss_rsrp float64, rssi float64) float64 {

	ss_rsrq = (math.Pow(10, (ss_rsrp / 10))) / ((math.Pow(10, (rssi / 10))) / float64(ssbrbcount))

	return 10 * math.Log10(ss_rsrq)

}
