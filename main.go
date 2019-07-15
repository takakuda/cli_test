package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "test app"
	app.Usage = "this app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("test") {
			fmt.Println(context.Args().Get(0) + "test")
		} else {
			fmt.Println(context.Args().Get(0))
		}

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "test, t",
			Usage: "Echo with test",
		},
	}

	app.Run(os.Args)
}
