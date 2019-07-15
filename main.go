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
		fmt.Println(context.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
