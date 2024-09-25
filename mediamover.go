package main

import (
	"github.com/seancallaway/mediamover/cmd"
)

func main() {
	cmd.Execute()
	/*
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
			// Load configuration file
			config, err := NewConfig(c.String("config"))
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			// TODO: Get list of tiles to process.
			if c.Args().Len() == 0 {
				return cli.NewExitError("Missing required arguments.")
			}
			loadPath := c.Args().Get(0)

			// TODO: Copy/move file.

			fmt.Println("This doesn't do anything yet. It would load media from " + loadPath + ".")
			fmt.Println(config.TvRoot)
			return nil
		}

		err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
