package main

import (
	"fmt"

	"github.com/hhjpin/gse"
)

var (
	text = "纽约时代广场, 纽约帝国大厦"

	seg gse.Segmenter
)

func main() {
	seg.LoadDict()

	hmm := seg.Cut(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutSearch(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutAll(text)
	fmt.Println("cut all: ", hmm)

	//
	hmm = seg.HMMCutMod(text)
	fmt.Println("hmm cut: ", hmm)
}
