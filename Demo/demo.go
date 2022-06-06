package main

import (
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

func main() {
	// set default concurrency number
	default_concurrency_num := runtime.NumCPU()

	// use cli for command line options
	var cliapp = &cli.App{
		Name:  "downloader",
		Usage: "Multi-threaded dowloader",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "'URL' from where the download begins",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "concurrency",
				Aliases: []string{"cn"},
				Value:   default_concurrency_num,
				Usage:   "Set default concurrency number",
			},
		},
	}
	err := cliapp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}