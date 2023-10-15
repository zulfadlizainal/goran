// nrConversion is a package that converts 5G NR properties from 3GPP pre-defined tables or formulas.
package nrConversion

// QciQos represents data structures for 5QI table with its Qos requirements.
// The data structures associated to variable QciQosTable.
type QciQos struct {
	Qci                     int         // 5qi
	ResourceType            string      // gbr or non-gbr
	PriorityLevel           float64     // level
	PacketDelayBudget       int         // ms
	PacketErrorLoss         float64     // %
	MaximumBurstSize        interface{} // bytes
	DataRateAveragingWindow interface{} // ms
	Services                string      // services
}

// QciQosTable is based on 3GPP TS 23.501 - Table 5.7.4-1.
// The table associated to type QciQos.
var QciQosTable = []QciQos{
	{1, "GBR", 2, 100, 1e-2, nil, nil, "Conversational Voice"},
	{2, "GBR", 4, 150, 1e-3, nil, nil, "Conversational Video"},
	{3, "GBR", 3, 50, 1e-3, nil, nil, "Real Time Gaming"},
	{4, "GBR", 5, 300, 1e-6, nil, nil, "Non Conversational Video (Buffered Streaming)"},
	{65, "GBR", 0.7, 75, 1e-2, nil, nil, "Mission Critical Push to Talk Voice (MCPTT)"},
	{66, "GBR", 2, 100, 1e-2, nil, nil, "Non-Mission Critical Push to Talk Voice (MCPTT)"},
	{67, "GBR", 1.5, 100, 1e-3, nil, nil, "Mission Critical Video"},
	{75, "GBR", 2.5, 50, 1e-2, nil, nil, "V2X Messages"},
	{5, "Non-GBR", 1, 100, 1e-6, nil, nil, "IMS Signaling"},
	{6, "Non-GBR", 6, 300, 1e-6, nil, nil, "Video (Buffered Streaming), TCP Based"},
	{7, "Non-GBR", 7, 100, 1e-3, nil, nil, "Voice, Live Streaming, Interactive Gaming"},
	{8, "Non-GBR", 8, 300, 1e-6, nil, nil, "Video (Buffered Streaming), TCP Based"},
	{9, "Non-GBR", 9, 300, 1e-6, nil, nil, "Video (Buffered Streaming), TCP Based"},
	{69, "Non-GBR", 0.5, 60, 1e-6, nil, nil, "Mission Critical Delay Sensitive Signaling"},
	{70, "Non-GBR", 5.5, 200, 1e-6, nil, nil, "Mission Critical Data"},
	{79, "Non-GBR", 6.5, 50, 1e-2, nil, nil, "V2X Messages"},
	{80, "Non-GBR", 6.8, 10, 1e-6, nil, nil, "Low Latency eMMB applications, Augmented Reality"},
	{82, "GBR (Delay Critical)", 1.9, 10, 1e-4, 255, 2.0, "Discrete Automation"},
	{83, "GBR (Delay Critical)", 2.2, 10, 1e-4, 1358, 2.0, "Discrete Automation"},
	{84, "GBR (Delay Critical)", 2.4, 30, 1e-5, 1358, 2.0, "Intelligent Transport System"},
	{85, "GBR (Delay Critical)", 2.1, 6, 1e-5, 255, 2.0, "Electricity Distribution – High Volt, Remote Control"},
}

// QCIToType converts the 5QI to Resource Type based on QciQosTable.
//   - qci refers to Qci within the range of QciQosTable.
//   - The function will return ResourceType within the range of QciQosTable.
//   - The function will return -1 error if the input is out of range.
func QCIToType(qci int) string {
	for _, QCIRequirements := range QciQosTable {
		if QCIRequirements.Qci == qci {
			return QCIRequirements.ResourceType
		}
	}
	return "-1"
}

// QCIToPriority converts the 5QI to Priority Level based on QciQosTable.
//   - qci refers to Qci within the range of QciQosTable.
//   - The function will return PriorityLevel within the range of QciQosTable.
//   - The function will return -1 error if the input is out of range.
func QCIToPriority(qci int) float64 {
	for _, QCIRequirements := range QciQosTable {
		if QCIRequirements.Qci == qci {
			return QCIRequirements.PriorityLevel
		}
	}
	return -1
}

// QCIToPacketDelay converts the 5QI to Packet Delay Budget based on QciQosTable.
//   - qci refers to Qci within the range of QciQosTable.
//   - The function will return PacketDelayBudget within the range of QciQosTable.
//   - The function will return -1 error if the input is out of range.
func QCIToPacketDelay(qci int) int {
	for _, QCIRequirements := range QciQosTable {
		if QCIRequirements.Qci == qci {
			return QCIRequirements.PacketDelayBudget
		}
	}
	return -1
}

// QCIToPacketLoss converts the 5QI to Packet Error Loss based on QciQosTable.
//   - qci refers to Qci within the range of QciQosTable.
//   - The function will return PacketErrorLoss within the range of QciQosTable.
//   - The function will return -1 error if the input is out of range.
func QCIToPacketLoss(qci int) float64 {
	for _, QCIRequirements := range QciQosTable {
		if QCIRequirements.Qci == qci {
			return QCIRequirements.PacketErrorLoss
		}
	}
	return -1
}
