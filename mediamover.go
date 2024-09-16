package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mediamover"
	app.Usage = "Rename and move movie and TV files"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		config, err := NewConfig(c.String("config"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("This doesn't do anything yet. Sorry.")
		fmt.Println(config.TvRoot)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
