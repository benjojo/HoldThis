package main

import (
	mainio "./mainio"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "HoldThis - Jumperr"
	app.Usage = "$ Jump <name of hold>"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 1 {
			// Now we need to save it.
			os.Chdir(mainio.Get(c.Args()[0]).Value)
		} else {
			println(app.Usage)
		}
	}

	app.Run(os.Args)
}
