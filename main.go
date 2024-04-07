package main

import (
	"github.com/KartoonYoko/errcheck/pkg/errcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		errcheck.ErrCheckAnalyzer,
	)
}
