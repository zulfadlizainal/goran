`Experimental: I build this to learn about Go packages. Not ready to be forked for real use yet.`

# goran

goran is 4G/5G radio access network tools that normally used by telecom RF engineers. Written in [Go programming language](http://golang.org).

## Supported Packages

| Technology | Package-Name | Functions    | Purpose                   |
|------------|--------------|--------------|---------------------------|
| 5G NR      | nrdownlink   | Tbs()        | Calculate TBS             |
| 5G NR      | nruplink     | Pusch()      | Calculate PUSCH           |
| 5G NR      | nruplink     | UlSinr()     | Calculate PUSCH           |
| 5G NR      | nrscheduling | RoundRobin() | Calculate Throughput      |
| 5G NR      | nrpathloss   | Macro()      | Calculate Macro Path Loss |

## Installation

```bash
go get github.com/zulfadlizainal/goran
```

## Examples

TBD

## Documentations

TBD

## Roadmap

This project have no roadmap. This project taking a build to solve approach. No functions will be developed for non-existing problem.

## Contribution Guidelines

1. Write comprehensive comments for each functions.
2. Include test files for each functions.

## Licenses

[GNU AFFERO GENERAL PUBLIC LICENSE](https://github.com/zulfadlizainal/goran/blob/main/LICENSE)