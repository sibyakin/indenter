package main

import (
	"os"

	"github.com/sibyakin/indenter"
)

func main() {
	j := indenter.NewJSON(os.Stdout)
	j.SetIndent(4)
	j.Run()
}
