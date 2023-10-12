# goran

goran is 4G/5G radio access network tools that commonly used by telecom RF engineers. Written in Go.

### Supported Packages

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

### Getting Started

```bash
go mod init <module-name>
go get github.com/zulfadlizainal/goran
```

### Examples

TBD

```bash

```

### Contribution Guidelines

1. Build only functions that could improve RF engineer workflows - Don't create functions for non-existing problems.
2. Write comprehensive comments for each functions - Comments will be part of the official docs automatically.
3. Include test files for each functions - Test values should loop a range of desired and undesired inputs.
4. Include error control for undesired inputs - return value for error is flexible.

### Version Control

Version X.Y.Z

1. X increment by 1 for major releases. E.g. New technology support, major refactoring, etc.
2. X = 0 for pre-release (non production ready)
3. Y increment by 1 for minor releases. E.g. New package support, new function support, new feature support, etc.
4. Z increment by 1 for bug fixes.

### Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE)
