package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gli"
	app.Version = Version
	app.Usage = ""
	app.Author = "Yusuke Hatanaka"
	app.Email = "arbalestimp@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
