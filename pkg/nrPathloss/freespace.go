package nrPathloss

import "math"

// FreeSpace generate path loss from using Free Space Path Loss Model.
//   - distance refers to distance from the transmitter to the receiver (Km).
//   - freq refers to operating signal frequency (MHz).
//   - gaintx refers to transmitter antenna gain (dB).
//   - gainrx refers to receiver antenna gain (dB).
//   - The function will return path loss (dB).
func FreeSpace(distance, freq, gaintx, gainrx float64) float64 {

	pathloss := (20 * math.Log10(distance)) + (20 * math.Log10(freq)) + 32.44 - gaintx - gainrx

	return pathloss

}
