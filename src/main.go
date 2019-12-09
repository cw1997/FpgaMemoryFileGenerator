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

	"github.com/cw1997/FpgaMemoryFileGenerator/converter/ascii"
	"github.com/cw1997/FpgaMemoryFileGenerator/generator"
	"github.com/cw1997/FpgaMemoryFileGenerator/generator/coe"
	"github.com/cw1997/FpgaMemoryFileGenerator/generator/mif"
	"github.com/cw1997/FpgaMemoryFileGenerator/utils"

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

			&cli.IntFlag{
				Name:        "address-radix",
				Usage:       "Radix of address (Only be used for mif format)",
				Aliases:     []string{"ar"},
				Value:       16,
				DefaultText: "16 hex",
			},
			&cli.IntFlag{
				Name:        "depth",
				Usage:       "Depth of memory (Only be used for mif format)",
				Aliases:     []string{"d"},
				Value:       16,
				DefaultText: "16",
			},

			&cli.IntFlag{
				Name:        "width",
				Usage:       "Bit width of memory (Only be used for mif format)",
				Aliases:     []string{"w"},
				Value:       16,
				DefaultText: "16",
			},
			&cli.IntFlag{
				Name:        "data-radix",
				Usage:       "Radix of input data",
				Aliases:     []string{"dr"},
				Value:       16,
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
			dataSource := ctx.String("data-source")

			str, err := utils.ReadStringFromFile(dataSource)
			if err != nil {
				log.Fatalf("[Error] DataSource error, pleasc check the file path '%s'. %s.\n", dataSource, err)
			}

			var outputBytes []byte

			dataType := ctx.String("data-type")
			switch dataType {
			case "ascii":
				outputBytes = ascii.Converter(str)
				break

			}

			var g generator.Generator

			outputType := ctx.String("output-type")
			switch outputType {
			case "mif":
				depth := ctx.Int("depth")
				width := ctx.Int("width")
				addressRadix := ctx.Int("address-radix")
				dataRadix := ctx.Int("data-radix")
				g = mif.NewMifGenerator(depth, width, addressRadix, dataRadix)
				break

			case "coe":
				width := ctx.Int("width")
				dataRadix := ctx.Int("data-radix")
				g = coe.NewCoeGenerator(width, dataRadix)
				break

			default:
				log.Fatalf("[Error] Unsupport output type, " +
					"now only support Altera Memory Initialization File 'mif' and Xilinx memory content 'coe', " +
					"pleasc check the output type.\n")
				break

			}

			outputStr := g.Generate(outputBytes)
			outputPath := ctx.String("output-path")
			err = utils.WriteStrToFile(outputStr, outputPath)
			if err != nil {
				log.Fatalf("[Error] Save output result error, pleasc check the output file path '%s'.\n", outputPath)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
