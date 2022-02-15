package main

import (
	"findKaramojiPkgfunc"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findKaramojiPkgfunc.Analyzer) }
