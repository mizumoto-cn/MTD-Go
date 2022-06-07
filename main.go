package main

import (
	"log"
	"os"
	"runtime"

	"github.com/mizumoto-cn/MTD-Go/core"

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
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "save as file at output path",
			},
			&cli.IntFlag{
				Name:    "concurrency",
				Aliases: []string{"c"},
				Value:   default_concurrency_num,
				Usage:   "Set default concurrency number",
			},
		},
		Action: func(ctx *cli.Context) error {
			url := ctx.String("url")
			filename := ctx.String("output")
			concurrency_num := ctx.Int("concurrency")
			return core.NewDownloader(concurrency_num).Download(url, filename)
		},
	}

	err := cliapp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
