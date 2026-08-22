package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/galexrt/protosetter"
)

func main() {
	singlechecker.Main(protosetter.NewAnalyzer(nil))
}
