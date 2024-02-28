package main

import (
	"github.com/tenderly/nitro/linters/koanf"
	"github.com/tenderly/nitro/linters/pointercheck"
	"github.com/tenderly/nitro/linters/rightshift"
	"github.com/tenderly/nitro/linters/structinit"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		koanf.Analyzer,
		pointercheck.Analyzer,
		rightshift.Analyzer,
		structinit.Analyzer,
	)
}
