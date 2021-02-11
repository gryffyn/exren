package main

import (
	"log"
	"os"

	"git.neveris.one/gryffyn/exren/parser"
)

func main() {
	format := `%DateTimeOriginal%-gryffyn.jpg`
	path := "test.jpg"
	exifData := parser.NewExif(path)
	_ = exifData.Parse()
	newpath := parser.ParseFormat(format, exifData.Tags)
	println("PATH: " + path)
	println("NEWPATH: " + newpath)
	err := os.Rename(path, newpath)
	if err != nil {
		log.Fatal(err)
	}
}
