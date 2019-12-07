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
	BuildTime = time.Now().Format("2006-01-02 15:04:05")
)

func main() {
	log.Println("ok")

	version := fmt.Sprintf("%d.%d.%d Build %s", MajorVersionNumber, MinorVersionNumber, RevisionVersionNumber, BuildTime)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("all2mif version all2mif%s", c.App.Version)
	}

	app := &cli.App{
		Name:    AppName,
		Version: version,
	}
	app.Run(os.Args)

}
