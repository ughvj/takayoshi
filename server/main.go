package main

import (
	"github.com/ughvj/takayoshi/processing"
)

func main() {
	e := processing.Init()
	e.Start(":2434")
}
