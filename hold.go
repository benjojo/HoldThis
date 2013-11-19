package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "HoldThis - Holderr"
	app.Usage = "$ Hold <name of hold>"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 1 {
			str, _ := os.Getwd()
			println("Saving the directory", str, "as the name '", c.Args()[0], "'")
			// Now we need to save it.
		} else {
			println(app.Usage)
		}
	}

	app.Run(os.Args)
}
