package main

import (
	"fmt"
	"log"
	"os"

	"git.neveris.one/gryffyn/exren/parser"
	. "github.com/logrusorgru/aurora/v3"
	"github.com/urfave/cli/v2"
)

func Run() {
	var format string

	app := &cli.App{
		Name:    "exren",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Usage:       "Output format, including extension",
				Destination: &format,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			err := rename(format, c.Args().Get(0))
			return err
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func rename(format, path string) error {
	exifData := parser.NewExif(path)
	_ = exifData.Parse()
	newpath := parser.ParseFormat(format, exifData.Tags)
	fmt.Println(Bold(Green("✓")), Bold("Renamed to:"), newpath)
	return os.Rename(path, newpath)
}

func main() {
	Run()
}
