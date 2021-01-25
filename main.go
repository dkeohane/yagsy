package main

import (
	"github.com/dkeohane/yagsy/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run()
}
