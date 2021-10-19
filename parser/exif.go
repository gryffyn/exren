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

func NewExif(path string) *Exif {
	e := new(Exif)
	e.Tags = make(Tags)
	e.open(path)
	return e
}

func (e *Exif) open(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
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
	x, err := goexif.Decode(e.File)
	if err != nil {return err}
	walker := walker{ed: x, Tags: Tags{}}
	err = x.Walk(&walker)
	e.Tags = walker.Tags
	return err
}
