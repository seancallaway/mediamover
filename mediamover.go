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

	app.Action = func(c *cli.Context) error {
		fmt.Println("This doesn't do anything yet. Sorry.")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
