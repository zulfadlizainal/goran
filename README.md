<span style="line-height: 1.05;">
<small>

## goran

`goran` is 4G/5G radio access network tools that commonly used by telecom RF engineers. Written in Go.

#### Supported Packages and Functions

See packages details at this [link](https://github.com/zulfadlizainal/goran/blob/main/packages.csv).

| Technology | Package      | Function                      | Purpose                                          |
|------------|--------------|-------------------------------|--------------------------------------------------|
| 5G NR      | nrConversion | NumerologyToScs()             | Converts Numerology (µ) to SCS (kHz)             |
| 5G NR      | nrConversion | NumerologyToSymbolPerSlot()   | Converts Numerology (µ) to Symbol/Slot (Count)   |
| 5G NR      | nrConversion | NumerologyToSlotPerSubframe() | Converts Numerology (µ) to Slot/Subframe (Count) |
| 5G NR      | nrConversion | NumerologyToSlotPerFrame()    | Converts Numerology (µ) to Slot/Frame (Count)    |
| 5G NR      | nrConversion | QCIToType()                   | Converts 5QI to Bit Rate Type                    |
| 5G NR      | nrConversion | QCIToPriority()               | Converts 5QI to Priority                         |
| 5G NR      | nrConversion | QCIToPacketDelay()            | Converts 5QI to Packet Delay (ms)                |
| 5G NR      | nrConversion | QCIToPacketLoss()             | Converts 5QI to Packet Loss Rate (%)             |
| 5G NR      | nrConversion | BandwidthToRB()               | Converts Bandwidth (MHz) to RB (Count)           |

#### Getting Started

Go to project directory.

```bash
cd <project-directory>
```

Initialize Go modules.

```bash
go mod init <module-name>
```

Download `goran` package.

```bash
go get github.com/zulfadlizainal/goran
```

#### Examples

To be added.

#### Contribution Guidelines

1. Build only to improve RF engineer workflows - don't create functions for non-existing problems.
2. Write comprehensive comments for each functions - comments will be in official docs automatically.
3. Include test for each functions - test values should cover range of desired and undesired inputs.
4. Include error control for undesired inputs - return value for error is flexible.

#### Test

Clone repository.

```bash
git clone github.com/zulfadlizainal/goran
```

Go to test directory.

```bash
cd goran/test
```

Run Go test.

```bash
go test
```

#### Version Control

Version X.Y.Z

X = Major releases
Y = Minor releases
Z = Bug fixes

#### Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE)

</small>
</span>
