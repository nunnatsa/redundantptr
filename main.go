package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/nunnatsa/redundantptr/redundantptrpkg"
)

func main() {
	singlechecker.Main(redundantptrpkg.NewAnalyzer())
}
