// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/8 6:07
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	AppName               = "FpgaMemoryFileGenerator"
	MajorVersionNumber    = 0
	MinorVersionNumber    = 1
	RevisionVersionNumber = 0
)

var (
	//timeStr = time.Now().Format("2006-01-02 15:04:05")
	BuildTime = time.Now().Format("20060102150405")
)

func main() {
	//log.Println("ok")

	version := fmt.Sprintf("%d.%d.%d Build%s", MajorVersionNumber, MinorVersionNumber, RevisionVersionNumber, BuildTime)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("FpgaMemoryFileGenerator V%s", c.App.Version)
	}

	app := &cli.App{
		Name:     AppName,
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "cw1997",
				Email: "changwei1006@gmail.com",
			},
		},
		Copyright: "(c) 2019 cw1997.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output-type",
				Usage:       "Type of output file",
				Aliases:     []string{"ot"},
				Value:       "mif",
				DefaultText: "mif",
			},
			&cli.StringFlag{
				Name:        "output-path",
				Usage:       "Output full path and file name",
				Aliases:     []string{"op"},
				Value:       "./output",
				DefaultText: "current path, filename: ./output",
			},

			&cli.StringFlag{
				Name:        "radix-address",
				Usage:       "Radix of address (Only be used for mif format)",
				Aliases:     []string{"ra"},
				Value:       "16",
				DefaultText: "16 hex",
			},
			&cli.StringFlag{
				Name:        "depth",
				Usage:       "Depth of memory (Only be used for mif format)",
				Aliases:     []string{"d"},
				Value:       "16",
				DefaultText: "16",
			},

			&cli.StringFlag{
				Name:        "width",
				Usage:       "Bit width of memory (Only be used for mif format)",
				Aliases:     []string{"w"},
				Value:       "16",
				DefaultText: "16",
			},
			&cli.StringFlag{
				Name:        "radix-data",
				Usage:       "Radix of input data",
				Aliases:     []string{"rd"},
				Value:       "Radix of data",
				DefaultText: "16 hex",
			},

			&cli.StringFlag{
				Name:        "data-source",
				Usage:       "Input data source",
				Aliases:     []string{"ds"},
				Value:       "./input",
				DefaultText: "current path, filename: ./input",
			},
			&cli.StringFlag{
				Name:        "data-type",
				Usage:       "Type of input file",
				Aliases:     []string{"dt"},
				Value:       "ascii",
				DefaultText: "ascii",
			},
		},
		Action: func(ctx *cli.Context) error {
			//if !ctx.Bool("ginger-crouton") {
			//	return cli.Exit("Ginger croutons are not in the soup", 86)
			//}
			log.Println(ctx)
			log.Println(ctx.String("output-path"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
