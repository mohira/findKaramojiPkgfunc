package findKaramojiPkgfunc

import (
	"golang.org/x/tools/go/analysis"
)

const doc = "findKaramojiPkgfunc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findKaramojiPkgfunc",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
