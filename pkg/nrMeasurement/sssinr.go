package nrMeasurement

import "math"

// SS-SINR is the ratio of the wanted signal power to the interference plus noise power measured at RE used for SSS.
// SS-SINR is defined as SS-RSRP / [(NCell SS-RSRP) + Noise]
var ss_sinr float64

// SsSinr calculates SS-SINR based on neighbor cell interference and noise figure.
//   - ss_rsrp refers to serving cell SS-RSRP in dBm.
//   - ncell_ss_rsrp refers to neighbor cell SS-RSRP in dBm.
//   - scs refers to SSB sub carrier spacing in kHz.
//   - noisefig refers to noise other than thermal noise in dB. Common value = 7dB.
//   - The function will return SS-SINR value in dB.
func SsSinr(ss_rsrp float64, ncell_ss_rsrp float64, scs int, noisefig float64) float64 {

	if scs != 15 && scs != 30 && scs != 120 && scs != 240 {
		panic("out of range: supported scs for SSB are 15, 30, 120, 240")
	}

	thermalnoise := -174 + 10*math.Log10(float64(scs*1000))
	totalnoise := thermalnoise + noisefig

	ss_sinr = math.Pow(10, (ss_rsrp/10)) / ((math.Pow(10, (ncell_ss_rsrp / 10))) + math.Pow(10, (totalnoise/10)))

	return 10 * math.Log10(ss_sinr)
}
