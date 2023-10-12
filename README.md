<span style="line-height: 1;">
<small>

## goran

`goran` is 4G/5G radio access network tools that commonly used by telecom RF engineers. Written as Go packages. Visit official landing page for `goran` [here](https://pkg.go.dev/github.com/zulfadlizainal/goran).

<br>
<img src="https://raw.githubusercontent.com/zulfadlizainal/goran/main/assets/logo.png" width=30% height=30% />
<br>

#### Documentations

Visit official documentations at [link](https://pkg.go.dev/github.com/zulfadlizainal/goran/pkg).

#### Supported Packages

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

- Build only to improve RF engineer workflows - don't create functions for non-existing problems.
- Write comprehensive comments for each functions - comments will be in official docs automatically.
- Include test for each functions - test values should cover range of desired and undesired inputs.
- Include error control for undesired inputs - return value for error is flexible.

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
