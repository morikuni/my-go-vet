package main

import (
	"github.com/MakeNowJust/enumcase"
	"github.com/gostaticanalysis/typeswitch"
	"golang.org/x/tools/go/analysis/unitchecker"

	mygovet "github.com/morikuni/my-go-vet"
)

func main() {
	unitchecker.Main(
		append(mygovet.StandardVetTools,
			typeswitch.Analyzer,
			enumcase.Analyzer,
		)...,
	)
}
