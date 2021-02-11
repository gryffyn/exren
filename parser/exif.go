package parser

import (
	"io"
	"log"
	"os"

	goexif "github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type Tags map[string]*tiff.Tag

type Exif struct {
	File io.Reader
	Tags Tags
}

type walker struct {
	ed   *goexif.Exif
	Tags Tags
}

func (e *Exif) Open(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Print(err)
	}
	e.File = f
}

// Implements goexif.Walker()
func (w *walker) Walk(name goexif.FieldName, tag *tiff.Tag) error {
	w.Tags[string(name)] = tag
	return nil
}

// Extracts the exif data from the file, then parses it into tags.
func (e *Exif) Parse() error {
	var x *goexif.Exif
	var err error
	x, err = goexif.Decode(e.File)
	walker := &walker{ed: x, Tags: make(Tags)}
	err = x.Walk(walker)
	e.Tags = walker.Tags
	return err
}
