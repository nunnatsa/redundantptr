package redundantptrpkg

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestAnalyser(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(), "samplecode")
}
