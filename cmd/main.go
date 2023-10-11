package main

import (
	"github.com/zulfadlizainal/goran/pkg/nr/nrdownlink"
	"github.com/zulfadlizainal/goran/pkg/nr/nrpathloss"
	"github.com/zulfadlizainal/goran/pkg/nr/nrscheduling"
	"github.com/zulfadlizainal/goran/pkg/nr/nruplink"
)

func main() {

	nrdownlink.Pss()
	nrdownlink.Tbs()
	nruplink.Pusch()
	nruplink.UlSinr()
	nrscheduling.RoundRobin()
	nrpathloss.Macro()

}
