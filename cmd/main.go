package main

import (
	"module/pkg/nr/nrdownlink"
	"module/pkg/nr/nrpathloss"
	"module/pkg/nr/nrscheduling"
	"module/pkg/nr/nruplink"
)

func main() {

	nrdownlink.Pss()
	nrdownlink.Tbs()
	nruplink.Pusch()
	nruplink.UlSinr()
	nrscheduling.RoundRobin()
	nrpathloss.Macro()

}
