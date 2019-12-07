package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	NAME = "FpgaMemoryFileGenerator"
	MAJOR_VERSION_NUMBER = 0
	MINOR_VERSION_NUMBER = 1
	REVISION_VERSION_NUMBER = 0
)

var (
	//timeStr = time.Now().Format("2006-01-02 15:04:05")
	BUILD_TIME = time.Now().Format("2006-01-02 15:04:05")
)

func main() {
	log.Println("ok")

	version := fmt.Sprintf("%d.%d.%d Build %s", MAJOR_VERSION_NUMBER, MINOR_VERSION_NUMBER, REVISION_VERSION_NUMBER, BUILD_TIME)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("all2mif version all2mif%s", c.App.Version)
	}

	app := &cli.App{
		Name: NAME,
		Version: version,
	}
	app.Run(os.Args)


}
