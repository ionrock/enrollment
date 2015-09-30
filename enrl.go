package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Boom"
	app.Usage = "Make an explosive entrance"

	hosts := cli.StringSlice{}
	
	app.Flags = []cli.Flag {
		cli.StringSliceFlag{
			Name: "host, H",
			Value: *hosts,
			Usage: "A host name to listen on",
		},
	}
	
	app.Action = func(c *cli.Context) {
		println(c.String("host"))
	}

	app.Run(os.Args)
}
