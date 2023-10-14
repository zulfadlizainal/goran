package nrConversion

// NumerologyFreqDomain represents data structures for Numerology with its frequency domain properties.
type NumerologyFreqDomain struct {
	Numerology int
	Scs        int
}

// 3GPP TS 38.300 Table 5.1-1.
var NumerologyToScsTable = []NumerologyFreqDomain{
	{0, 15},
	{1, 30},
	{2, 60},
	{3, 120},
	{4, 240},
}

// NumerologyToScs converts the Numerology (µ) input to SCS (kHz) value based on the 3GPP TS 38.300 Table 5.1-1.
// The function will return -1 error if the input is out of range.
func NumerologyToScs(numerology int) int {
	for _, NumerologyFreqDomain := range NumerologyToScsTable {
		if NumerologyFreqDomain.Numerology == numerology {
			return NumerologyFreqDomain.Scs
		}
	}
	return -1
}

// NumerologyTimeDomain represents data structures for Numerology with its time domain properties.
type NumerologyTimeDomain struct {
	Numerology      int
	SymbolPerSlot   int
	SlotPerSubframe int
	SlotPerFrame    int
}

// 3GPP TS 38.211 Table 4.3.2.
var NumerologyToTimeTable = []NumerologyTimeDomain{
	{0, 14, 1, 10},
	{1, 14, 2, 20},
	{2, 14, 4, 40},
	{3, 14, 8, 80},
	{4, 14, 16, 160},
	{5, 14, 32, 320},
	{6, 14, 64, 640},
}

// NumerologyToSymbolPerSlot converts the Numerology (µ) input to Symbol/Slot (Count) value based on the 3GPP TS 38.211 Table 4.3.2.
// The function will return -1 error if the input is out of range.
func NumerologyToSymbolPerSlot(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyToTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SymbolPerSlot
		}
	}
	return -1
}

// NumerologyToSlotPerSubframe converts the Numerology (µ) input to Slot/Subframe (Count) value based on the 3GPP TS 38.211 Table 4.3.2.
// The function will return -1 error if the input is out of range.
func NumerologyToSlotPerSubframe(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyToTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SlotPerSubframe
		}
	}
	return -1
}

// NumerologyToSlotPerFrame converts the Numerology (µ) input to Slot/Frame (Count) value based on the 3GPP TS 38.211 Table 4.3.2.
// The function will return -1 error if the input is out of range.
func NumerologyToSlotPerFrame(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyToTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SlotPerFrame
		}
	}
	return -1
}
