<span style="line-height: 1;">
<small>

## goran

`goran` is 4G/5G radio access network tools that commonly used by telecom RF engineers. Written in Go.

<br>
<img src="\assets\logo.png" width=30% height=30% />
<br>

#### Supported Packages and Functions

See packages details at this [link](https://github.com/zulfadlizainal/goran/blob/main/packages.csv).

| Technology | Package      | Function                      |
|------------|--------------|-------------------------------|
| 5G NR      | nrConversion | NumerologyToScs()             |
| 5G NR      | nrConversion | NumerologyToSymbolPerSlot()   |
| 5G NR      | nrConversion | NumerologyToSlotPerSubframe() |
| 5G NR      | nrConversion | NumerologyToSlotPerFrame()    |
| 5G NR      | nrConversion | QCIToType()                   |
| 5G NR      | nrConversion | QCIToPriority()               |
| 5G NR      | nrConversion | QCIToPacketDelay()            |
| 5G NR      | nrConversion | QCIToPacketLoss()             |
| 5G NR      | nrConversion | BandwidthToRB()               |

#### Installation

```bash
cd <project-directory> # Go to project directory.
go mod init <module-name> # Initialize Go modules.
go get github.com/zulfadlizainal/goran # Download `goran` package.
```

#### Examples

To be updated.

#### Contribution Guidelines

1. Build only to improve RF engineer workflows - don't create functions for non-existing problems.
2. Write comprehensive comments for each functions - comments will be in official docs automatically.
3. Include test for each functions - test values should cover range of desired and undesired inputs.
4. Include error control for undesired inputs - return value for error is flexible.

#### Test Packages

```bash
git clone github.com/zulfadlizainal/goran.git # Clone repository.
cd goran/test # Go to test directory.
go test # Run Go test.
```

#### Version Control

Version X.Y.Z<br>
X = Major releases, Y = Minor releases, Z = Bug fixes.

#### Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE)

</small>
</span>
