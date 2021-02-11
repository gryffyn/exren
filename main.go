package main

import (
	"log"
	"os"

	"git.neveris.one/gryffyn/gfn-fRen/parser"
)

func main() {
	format := `%DateTimeOriginal%-gryffyn.jpg`
	path := "test.jpg"
	exifData := parser.NewExif(path)
	_ = exifData.Parse()
	println(exifData.Tags["DateTimeOriginal"])
	err := os.Rename(path, parser.ParseFormat(format, exifData.Tags))
	if err != nil {
		log.Fatal(err)
	}
}

// gfn-fRen -f '%DateTimeOriginal%-gryffyn.jpg' test.jpg
