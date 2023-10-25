<span style="line-height: 1;">
<small>

## goran

`goran` is 5G NR radio access network packages that commonly used in the workflow of telecom RF engineers for simulations and calculations. Visit `goran` project landing page here -> [go.dev///goran](https://pkg.go.dev/github.com/zulfadlizainal/goran).

[![CI](https://github.com/zulfadlizainal/goran/actions/workflows/go_pkgtest.yaml/badge.svg)](https://github.com/zulfadlizainal/goran/actions/workflows/go_pkgtest.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zulfadlizainal/goran.svg)](https://pkg.go.dev/github.com/zulfadlizainal/goran)

<br>
<img src="https://raw.githubusercontent.com/zulfadlizainal/goran/main/assets/logo.png" width=20% height=20% />
<br>

#### Documentations

[![Static Badge](https://img.shields.io/badge/goran-Docs-blue)](https://pkg.go.dev/github.com/zulfadlizainal/goran/pkg)

Visit official documentations page [here](https://pkg.go.dev/github.com/zulfadlizainal/goran/pkg). 

#### Supported Packages

[![Static Badge](https://img.shields.io/badge/Packages-8A2BE2)](https://github.com/zulfadlizainal/goran#supported-packages)

```markdown

| Technology | Package       | Function                      | Purpose                                          |
|------------|---------------|-------------------------------|--------------------------------------------------|
| 5G NR      | nrConversion  | BandwidthToRB()               | Converts Bandwidth (MHz) to RB (Count)           |
| 5G NR      | nrConversion  | McsToQm()                     | Converts MCS to Modulation Order (bits/Symbol)   |
| 5G NR      | nrConversion  | McsToR()                      | Converts MCS to Target Code Rate (R)             |
| 5G NR      | nrConversion  | McsToSe()                     | Converts MCS to Spectral Efficiency (bps/Hz)     |
| 5G NR      | nrConversion  | NumerologyToScs()             | Converts Numerology (µ) to SCS (kHz)             |
| 5G NR      | nrConversion  | NumerologyToSlotPerFrame()    | Converts Numerology (µ) to Slot/Frame (Count)    |
| 5G NR      | nrConversion  | NumerologyToSlotPerSubframe() | Converts Numerology (µ) to Slot/Subframe (Count) |
| 5G NR      | nrConversion  | NumerologyToSymbolPerSlot()   | Converts Numerology (µ) to Symbol/Slot (Count)   |
| 5G NR      | nrConversion  | QCIToPacketDelay()            | Converts 5QI to Packet Delay (ms)                |
| 5G NR      | nrConversion  | QCIToPacketLoss()             | Converts 5QI to Packet Loss Rate (%)             |
| 5G NR      | nrConversion  | QCIToPriority()               | Converts 5QI to Priority                         |
| 5G NR      | nrConversion  | QCIToType()                   | Converts 5QI to Bit Rate Type                    |
| 5G NR      | nrDownlink    | Tbs()                         | Calculates Transport Block Size (Bytes)          |
| 5G NR      | nrMeasurement | SsRsrp()                      | Calculates SS-RSRP (dBm)                         |
| 5G NR      | nrMeasurement | SsRsrq()                      | Calculates SS-RSRQ (dB)                          |
| 5G NR      | nrMeasurement | SsSinr()                      | Calculates SS-SINR (dB)                          |
| 5G NR      | nrPathloss    | FreeSpace()                   | Generates Free Space Path Loss (dB)              |

```

#### Getting Started

[![Static Badge](https://img.shields.io/badge/Install-8A2BE2)](https://github.com/zulfadlizainal/goran#getting-started)

```bash
# Go to project directory.
cd <project-directory> 

# Initialize Go modules.
go mod init <module-name> 

# Download `goran` package.
go get github.com/zulfadlizainal/goran 

# Get updated package. (Optional)
go get -u github.com/zulfadlizainal/goran
```

[![Static Badge](https://img.shields.io/badge/Use-8A2BE2)](https://github.com/zulfadlizainal/goran#getting-started)

```go
package main

import (
	"fmt"
	"log"

	// Import goran package
	"github.com/zulfadlizainal/goran/pkg/nrConversion"
)

func main() {

	bandwidth := 20
	freqrange := "Sub6" 
	scs := 30 

	// Call the function
	rbCount, err := nrConversion.BandwidthToRB(bandwidth, freqrange, scs)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("RB count = %vRB\n", rbCount)

}
```

#### Verification

[![Static Badge](https://img.shields.io/badge/Test-8A2BE2)](https://github.com/zulfadlizainal/goran#verification)

```bash
git clone github.com/zulfadlizainal/goran.git # Clone repository.
cd goran/pkg/<package-name> # Go to test directory.
go test # Run Go test.
```

#### Contribution

[![Static Badge](https://img.shields.io/badge/Roadmap-8A2BE2)](https://github.com/zulfadlizainal/goran#contribution)

```markdown
# What to Build

1. Flexible. No explicit roadmap. 
2. Build functions to address common RF engineers workflow.
3. Avoid functions creation for non-existing problems.
```

```markdown
# Wish List

1. Generate urban and rural path loss.
2. Calculate RSSI.
3. Calculate PRACH power control, PUSCH power control. 
4. Calculate impact of measurement gap to downlink throughput.
5. Examples of goran packages usage.
```

[![Static Badge](https://img.shields.io/badge/Coding-8A2BE2)](https://github.com/zulfadlizainal/goran#contribution)

```markdown
# Guidelines

1. No dependencies outside Go standard library.
2. Include comprehensive comments for the codes. Documentation is automated based on the comments.
3. Include test for each functions with desired and undesired input.
4. Include error control. Return value for error is flexible but need to specify.
```

<br><a href="https://www.buymeacoffee.com/zulfadlizainal" target="blank"><img src="https://cdn.ko-fi.com/cdn/kofi2.png?v=2" alt="Buy Me A Coffee" height="37.5" width="127.5"></a>

#### Licenses

GNU AFFERO GPL Version 3

</small>
</span>
