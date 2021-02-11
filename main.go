package main

import (
	"log"
	"os"

	"git.neveris.one/gryffyn/gfn-fRen/parser"
)

func main() {
	format := `%DateTimeOriginal%-gryffyn.jpg`
	path := "test.jpg"
	exifData := parser.Exif{}
	exifData.Open(path)
	_ = exifData.Parse()
	err := os.Rename(path, parser.ParseFormat(format, exifData.Tags))
	if err != nil {
		log.Fatal(err)
	}
}

// gfn-fRen -n '%DateTimeOriginal%-gryffyn.jpg' test.jpg
