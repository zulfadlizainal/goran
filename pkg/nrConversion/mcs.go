package nrConversion

import "errors"

// McsIndex represents data structures for MCS table with its Modulation requirements.
// The data structures associated to variable McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
type McsIndex struct {
	Mcs                int     // Mcs index
	Qm                 int     // modulation order bits/symbol
	R                  float64 // target code rate
	SpectralEfficiency float64 // Qm x R
}

// McsIndexTableOne is based on 3GPP TS 38.214 - Table 5.1.3.1-1.
// The table associated to type McsIndex.
// The table will return 0 for reserved R and SpectralEfficiency value.
var McsIndexTableOne = []McsIndex{
	{0, 2, 0.1171875, 0.2344},
	{1, 2, 0.153320313, 0.3066},
	{2, 2, 0.188476563, 0.3770},
	{3, 2, 0.245117188, 0.4902},
	{4, 2, 0.30078125, 0.6016},
	{5, 2, 0.370117188, 0.7402},
	{6, 2, 0.438476563, 0.8770},
	{7, 2, 0.513671875, 1.0273},
	{8, 2, 0.587890625, 1.1758},
	{9, 2, 0.663085938, 1.3262},
	{10, 4, 0.33203125, 1.3281},
	{11, 4, 0.369140625, 1.4766},
	{12, 4, 0.423828125, 1.6953},
	{13, 4, 0.478515625, 1.9141},
	{14, 4, 0.540039063, 2.1602},
	{15, 4, 0.6015625, 2.4063},
	{16, 4, 0.642578125, 2.5703},
	{17, 6, 0.427734375, 2.5664},
	{18, 6, 0.455078125, 2.7305},
	{19, 6, 0.504882813, 3.0293},
	{20, 6, 0.553710938, 3.3223},
	{21, 6, 0.6015625, 3.6094},
	{22, 6, 0.650390625, 3.9023},
	{23, 6, 0.702148438, 4.2129},
	{24, 6, 0.75390625, 4.5234},
	{25, 6, 0.802734375, 4.8164},
	{26, 6, 0.852539063, 5.1152},
	{27, 6, 0.888671875, 5.3320},
	{28, 6, 0.92578125, 5.5547},
	{29, 2, 0.0, 0.0},
	{30, 4, 0.0, 0.0},
	{31, 6, 0.0, 0.0},
}

// McsIndexTableTwo is based on 3GPP TS 38.214 - Table 5.1.3.1-2.
// The table associated to type McsIndex.
// The table will return 0 for reserved R and SpectralEfficiency value.
var McsIndexTableTwo = []McsIndex{
	{0, 2, 0.1171875, 0.2344},
	{1, 2, 0.188476563, 0.377},
	{2, 2, 0.30078125, 0.6016},
	{3, 2, 0.438476563, 0.877},
	{4, 2, 0.587890625, 1.1758},
	{5, 4, 0.369140625, 1.4766},
	{6, 4, 0.423828125, 1.6953},
	{7, 4, 0.478515625, 1.9141},
	{8, 4, 0.540039063, 2.1602},
	{9, 4, 0.6015625, 2.4063},
	{10, 4, 0.642578125, 2.5703},
	{11, 6, 0.455078125, 2.7305},
	{12, 6, 0.504882813, 3.0293},
	{13, 6, 0.553710938, 3.3223},
	{14, 6, 0.6015625, 3.6094},
	{15, 6, 0.650390625, 3.9023},
	{16, 6, 0.702148438, 4.2129},
	{17, 6, 0.75390625, 4.5234},
	{18, 6, 0.802734375, 4.8164},
	{19, 6, 0.852539063, 5.1152},
	{20, 8, 0.666503906, 5.332},
	{21, 8, 0.694335938, 5.5547},
	{22, 8, 0.736328125, 5.8906},
	{23, 8, 0.778320313, 6.2266},
	{24, 8, 0.821289063, 6.5703},
	{25, 8, 0.864257813, 6.9141},
	{26, 8, 0.895019531, 7.1602},
	{27, 8, 0.92578125, 7.4063},
	{28, 2, 0.0, 0.0},
	{29, 4, 0.0, 0.0},
	{30, 6, 0.0, 0.0},
	{31, 8, 0.0, 0.0},
}

// McsIndexTableThree is based on 3GPP TS 38.214 - Table 5.1.3.1-3.
// The table associated to type McsIndex.
// The table will return 0 for reserved R and SpectralEfficiency value.
var McsIndexTableThree = []McsIndex{
	{0, 2, 0.029296875, 0.0586},
	{1, 2, 0.0390625, 0.0781},
	{2, 2, 0.048828125, 0.0977},
	{3, 2, 0.0625, 0.125},
	{4, 2, 0.076171875, 0.1523},
	{5, 2, 0.096679688, 0.1934},
	{6, 2, 0.1171875, 0.2344},
	{7, 2, 0.153320313, 0.3066},
	{8, 2, 0.188476563, 0.377},
	{9, 2, 0.245117188, 0.4902},
	{10, 2, 0.30078125, 0.6016},
	{11, 2, 0.370117188, 0.7402},
	{12, 2, 0.438476563, 0.877},
	{13, 2, 0.513671875, 1.0273},
	{14, 2, 0.587890625, 1.1758},
	{15, 4, 0.33203125, 1.3281},
	{16, 4, 0.369140625, 1.4766},
	{17, 4, 0.423828125, 1.6953},
	{18, 4, 0.478515625, 1.9141},
	{19, 4, 0.540039063, 2.1602},
	{20, 4, 0.6015625, 2.4063},
	{21, 6, 0.427734375, 2.5664},
	{22, 6, 0.455078125, 2.7305},
	{23, 6, 0.504882813, 3.0293},
	{24, 6, 0.553710938, 3.3223},
	{25, 6, 0.6015625, 3.6094},
	{26, 6, 0.650390625, 3.9023},
	{27, 6, 0.702148438, 4.2129},
	{28, 6, 0.0, 0.0},
	{29, 4, 0.0, 0.0},
	{30, 6, 0.0, 0.0},
	{31, 8, 0.0, 0.0},
}

// MCS table constant.
const (
	TableOne   = 1 // McsIndexTableOne
	TableTwo   = 2 // McsIndexTableTwo
	TableThree = 3 // McsIndexTableThree
)

// McsToQm converts the MCS Index to Modulation Order based on McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - mcs refers to Mcs within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - tablenum refers to MCS table constant.
//   - The function will return Qm within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
func McsToQm(mcs int, tablenum int) (interface{}, error) {
	if tablenum == TableOne {
		for _, McsIndex := range McsIndexTableOne {
			if McsIndex.Mcs == mcs {
				return McsIndex.Qm, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableTwo {
		for _, McsIndex := range McsIndexTableTwo {
			if McsIndex.Mcs == mcs {
				return McsIndex.Qm, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableThree {
		for _, McsIndex := range McsIndexTableThree {
			if McsIndex.Mcs == mcs {
				return McsIndex.Qm, nil
			}
		}
		return nil, errors.New("out of range")
	}
	return nil, errors.New("out of range")
}

// McsToR converts the MCS Index to Target Code Rate based on McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - mcs refers to Mcs within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - tablenum refers to MCS table constant.
//   - The function will return R within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
func McsToR(mcs int, tablenum int) (interface{}, error) {
	if tablenum == TableOne {
		for _, McsIndex := range McsIndexTableOne {
			if McsIndex.Mcs == mcs {
				return McsIndex.R, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableTwo {
		for _, McsIndex := range McsIndexTableTwo {
			if McsIndex.Mcs == mcs {
				return McsIndex.R, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableThree {
		for _, McsIndex := range McsIndexTableThree {
			if McsIndex.Mcs == mcs {
				return McsIndex.R, nil
			}
		}
		return nil, errors.New("out of range")
	}
	return nil, errors.New("out of range")
}

// McsToSe converts the MCS Index to Spectral Efficiency based on McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - mcs refers to Mcs within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
//   - tablenum refers to MCS table constant.
//   - The function will return SpectralEfficiency within the range of McsIndexTableOne, McsIndexTableTwo, or McsIndexTableThree.
func McsToSe(mcs int, tablenum int) (interface{}, error) {
	if tablenum == TableOne {
		for _, McsIndex := range McsIndexTableOne {
			if McsIndex.Mcs == mcs {
				return McsIndex.SpectralEfficiency, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableTwo {
		for _, McsIndex := range McsIndexTableTwo {
			if McsIndex.Mcs == mcs {
				return McsIndex.SpectralEfficiency, nil
			}
		}
		return nil, errors.New("out of range")
	} else if tablenum == TableThree {
		for _, McsIndex := range McsIndexTableThree {
			if McsIndex.Mcs == mcs {
				return McsIndex.SpectralEfficiency, nil
			}
		}
		return nil, errors.New("out of range")
	}
	return nil, errors.New("out of range")
}
