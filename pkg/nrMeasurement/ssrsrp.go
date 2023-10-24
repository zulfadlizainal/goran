package nrMeasurement

import "math"

// Total SSS RE count constant (12 Sc x 20 RB).
const sssrecount int = 12 * 20

// SS-RSRP is the average received power of SSS RE across all bandwidth.
// SS-RSRP is calculated by dividing the total received power in all SSS RE by the total count of SSS RE.
var ss_rsrp float64

// SsRsrp calculates SS-RSRP based on total SSS RE received power.
//   - sssretotalrcvpow refers to total SSS RE received power in mW.
//   - The function will return SS-RSRP value in dBm.
func SsRsrp(sssretotalrcvpow float64) float64 {

	ss_rsrp = sssretotalrcvpow / float64(sssrecount)

	return 10 * math.Log10(ss_rsrp)

}
