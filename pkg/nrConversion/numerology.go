package nrConversion

// NumerologyFreqDomain represents data structures for Numerology table with its frequency domain conversion.
// The data structures associated to variable NumerologyScsTable.
type NumerologyFreqDomain struct {
	Numerology int // numerology
	Scs        int // kHz
}

// NumerologyScsTable is based on 3GPP TS 38.300 Table 5.1-1.
// The table associated to type NumerologyFreqDomain.
var NumerologyScsTable = []NumerologyFreqDomain{
	{0, 15},
	{1, 30},
	{2, 60},
	{3, 120},
	{4, 240},
}

// NumerologyToScs converts the Numerology to Sub Carrier Spacing based on NumerologyScsTable.
//   - numerology refers to Numerology within the range of NumerologyScsTable.
//   - The function will return Scs within the range of NumerologyScsTable.
//   - The function will return -1 error if the input is out of range.
func NumerologyToScs(numerology int) int {
	for _, NumerologyFreqDomain := range NumerologyScsTable {
		if NumerologyFreqDomain.Numerology == numerology {
			return NumerologyFreqDomain.Scs
		}
	}
	return -1
}

// NumerologyTimeDomain represents data structures for Numerology table with its time domain conversion.
// The data structures associated to variable NumerologyTimeTable.
type NumerologyTimeDomain struct {
	Numerology      int // numerology
	SymbolPerSlot   int // count
	SlotPerSubframe int // count
	SlotPerFrame    int // count
}

// NumerologyTimeTable is based on 3GPP TS 38.211 Table 4.3.2.
// The table associated to type NumerologyTimeDomain.
var NumerologyTimeTable = []NumerologyTimeDomain{
	{0, 14, 1, 10},
	{1, 14, 2, 20},
	{2, 14, 4, 40},
	{3, 14, 8, 80},
	{4, 14, 16, 160},
	{5, 14, 32, 320},
	{6, 14, 64, 640},
}

// NumerologyToSymbolPerSlot converts the Numerology to Symbol/Slot based on NumerologyTimeTable.
//   - numerology refers to Numerology within the range of NumerologyTimeTable.
//   - The function will return SymbolPerSlot within the range of NumerologyTimeTable.
//   - The function will return -1 error if the input is out of range.
func NumerologyToSymbolPerSlot(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SymbolPerSlot
		}
	}
	return -1
}

// NumerologyToSlotPerSubframe converts the Numerology to Slot/Subframe based on NumerologyTimeTable.
//   - numerology refers to Numerology within the range of NumerologyTimeTable.
//   - The function will return SlotPerSubframe within the range of NumerologyTimeTable.
//   - The function will return -1 error if the input is out of range.
func NumerologyToSlotPerSubframe(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SlotPerSubframe
		}
	}
	return -1
}

// NumerologyToSlotPerFrame converts the Numerology to Slot/Frame based on NumerologyTimeTable.
//   - numerology refers to Numerology within the range of NumerologyTimeTable.
//   - The function will return SlotPerFrame within the range of NumerologyTimeTable.
//   - The function will return -1 error if the input is out of range.
func NumerologyToSlotPerFrame(numerology int) int {
	for _, NumerologyTimeDomain := range NumerologyTimeTable {
		if NumerologyTimeDomain.Numerology == numerology {
			return NumerologyTimeDomain.SlotPerFrame
		}
	}
	return -1
}
