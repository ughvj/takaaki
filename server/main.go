package main

import (
	"github.com/ughvj/takaaki/processing"
)

func main() {
	e := processing.Init()
	e.Logger.Fatal(e.Start(":2434"))
}
