package main

import (
	"github.com/sibyakin/indenter/http"
)

func main() {
	app := http.NewJSON()
	app.Run(":8080", 4)
}
