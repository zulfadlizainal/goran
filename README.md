<span style="line-height: 1;">
<small>

## goran

`goran` is 5G radio access network tools that commonly used by telecom RF engineers written in Go. Visit `goran` project landing page here -> [go.dev///goran](https://pkg.go.dev/github.com/zulfadlizainal/goran).

<br>
<img src="https://raw.githubusercontent.com/zulfadlizainal/goran/main/assets/logo.png" width=30% height=30% />
<br>

#### Documentations

Visit official documentations [here](https://pkg.go.dev/github.com/zulfadlizainal/goran/pkg).

#### Supported Packages

All supported packages and functions maintained by this library is listed [here](https://raw.githubusercontent.com/zulfadlizainal/goran/main/docs/packages_index.md).

```markdown
| Package      | Function                      |
|--------------|-------------------------------|
| nrConversion | NumerologyToScs()             |
| nrConversion | NumerologyToSymbolPerSlot()   |
| nrConversion | NumerologyToSlotPerSubframe() |
| nrConversion | NumerologyToSlotPerFrame()    |
| nrConversion | QCIToType()                   |
| nrConversion | QCIToPriority()               |
| nrConversion | QCIToPacketDelay()            |
| nrConversion | QCIToPacketLoss()             |
| nrConversion | BandwidthToRB()               |
```

#### Examples

To be updated.

#### Installation

```bash
cd <project-directory> # Go to project directory.
go mod init <module-name> # Initialize Go modules.
go get github.com/zulfadlizainal/goran # Download `goran` package.
```

#### Test Packages

```bash
git clone github.com/zulfadlizainal/goran.git # Clone repository.
cd goran/test # Go to test directory.
go test # Run Go test.
```

#### Contribution Guidelines

```markdown
# What to Build
1. Flexible. No explicit roadmap. Build functions to address common RF engineers workflow.
2. Avoid functions creation for non-existing problems.

# Coding Standards
1. Include comprehensive comments for the codes. Documentation is automated based on the comments.
2. Include test for each functions with desired and undesired input.
3. Include error control for undesired inputs. Return value for error is flexible but need to specify.
4. Lint code using official go.dev linter.
```

#### Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE)

</small>
</span>
