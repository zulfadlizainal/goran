<span style="line-height: 1.05;">
<small>

## goran

`goran` is 4G/5G radio access network tools that commonly used by telecom RF engineers. Written in Go.

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
