package main

import (
	mainio "./mainio"
	"fmt"
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
			target := mainio.Get(c.Args()[0]).Value
			//fmt.Println(target)
			os.Chdir(target)
			os.Setenv("PWD", target)
			os.Setenv("pwd", target)
			fmt.Println("cd", os.Getenv("PWD"))

		} else {
			fmt.Println(app.Usage)
		}
	}

	app.Run(os.Args)
}
