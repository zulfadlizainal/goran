<span style="line-height: 1;">
<small>

## goran

`goran` is 5G NR radio access network tools that commonly used by telecom RF engineers. Visit `goran` project landing page here -> [go.dev///goran](https://pkg.go.dev/github.com/zulfadlizainal/goran). The project is written in Go.

[![CI](https://github.com/zulfadlizainal/goran/actions/workflows/go_pkgtest.yaml/badge.svg)](https://github.com/zulfadlizainal/goran/actions/workflows/go_pkgtest.yaml)

<br>
<img src="https://raw.githubusercontent.com/zulfadlizainal/goran/main/assets/logo.png" width=30% height=30% />
<br>

#### Documentations

Visit official documentations page [here](https://pkg.go.dev/github.com/zulfadlizainal/goran/pkg).

#### Supported Packages

[Packages Index](https://raw.githubusercontent.com/zulfadlizainal/goran/main/docs/packages_index.md) list.

#### Getting Started

Installation:

```bash
# Go to project directory.
cd <project-directory> 

# Initialize Go modules.
go mod init <module-name> 

# Download `goran` package.
go get github.com/zulfadlizainal/goran 
```

Using it in your project:

```go
package main

import (
	"fmt"
	"log"

	// Import goran package
	"github.com/zulfadlizainal/goran/pkg/nrConversion"
)

func main() {

	// Define the input parameters
	bandwidth := 20     // MHz
	freqrange := "Sub6" // FR1
	scs := 30           // 60 kHz

	// Call the function associated within the goran package
	rbCount, err := nrConversion.BandwidthToRB(bandwidth, freqrange, scs)

	// Stop the program if the function output error
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the result
	fmt.Printf("RB count = %vRB\n", rbCount)

}
```

Once in a while, update the package to the latest version:

```bash
go get -u github.com/zulfadlizainal/goran
```

#### Examples

To be updated.

#### Test Packages

```bash
git clone github.com/zulfadlizainal/goran.git # Clone repository.
cd goran/test # Go to test directory.
go test # Run Go test.
```

#### Contribution Guidelines

```markdown
# What to Build
1. Flexible. No explicit roadmap. 
2. Build functions to address common RF engineers workflow.
3. Avoid functions creation for non-existing problems.

# Coding Standards
1. Include comprehensive comments for the codes. Documentation is automated based on the comments.
2. Include test for each functions with desired and undesired input.
3. Include error control for undesired inputs. Return value for error is flexible but need to specify.
4. Lint code using official go.dev linter.
```

#### Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE).

</small>
</span>
