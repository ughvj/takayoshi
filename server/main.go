package main

import (
	"flag"

	"github.com/ughvj/takamori/processing"
)

func main() {
	isDryrun := flag.Bool("dryrun", true, "dryrun mode (default: true)")

	e := processing.Init(*isDryrun)
	e.Start(":2434")
}
