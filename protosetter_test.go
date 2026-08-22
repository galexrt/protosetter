package protosetter_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/galexrt/protosetter"
)

func Test(t *testing.T) {
	cfg := &protosetter.Config{}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, protosetter.NewAnalyzer(cfg))

	analysistest.Run(t, testdata, protosetter.NewAnalyzer(cfg), "./proto/...")
}
