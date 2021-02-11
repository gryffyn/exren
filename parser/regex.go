package parser

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/tiff"
)

// gfn-fRen -f '%DateTimeOriginal%-gryffyn.jpg' test.jpg

var validFields = []string{"ImageWidth", "ImageLength", "BitsPerSample", "Compression", "PhotometricInterpretation",
	"Orientation", "SamplesPerPixel", "PlanarConfiguration", "YCbCrSubSampling", "YCbCrPositioning", "XResolution",
	"YResolution", "ResolutionUnit", "DateTime", "ImageDescription", "Make", "Model", "Software", "Artist",
	"Copyright", "ExifIFDPointer", "GPSInfoIFDPointer", "InteroperabilityIFDPointer", "ExifVersion", "FlashpixVersion",
	"ColorSpace", "ComponentsConfiguration", "CompressedBitsPerPixel", "PixelXDimension", "PixelYDimension",
	"MakerNote", "UserComment", "RelatedSoundFile", "DateTimeOriginal", "DateTimeDigitized", "SubSecTime",
	"SubSecTimeOriginal", "SubSecTimeDigitized", "ImageUniqueID", "ExposureTime", "FNumber", "ExposureProgram",
	"SpectralSensitivity", "ISOSpeedRatings", "OECF", "ShutterSpeedValue", "ApertureValue", "BrightnessValue",
	"ExposureBiasValue", "MaxApertureValue", "SubjectDistance", "MeteringMode", "LightSource", "Flash", "FocalLength",
	"SubjectArea", "FlashEnergy", "SpatialFrequencyResponse", "FocalPlaneXResolution", "FocalPlaneYResolution",
	"FocalPlaneResolutionUnit", "SubjectLocation", "ExposureIndex", "SensingMethod", "FileSource", "SceneType",
	"CFAPattern", "CustomRendered", "ExposureMode", "WhiteBalance", "DigitalZoomRatio", "FocalLengthIn35mmFilm",
	"SceneCaptureType", "GainControl", "Contrast", "Saturation", "Sharpness", "DeviceSettingDescription",
	"SubjectDistanceRange"}
var dateFields = []string{"DateTime", "DateTimeOriginal", "DateTimeDigitized"}

func contains(arr []string, str string) bool {
	for _, n := range arr {
		if str == n {
			return true
		}
	}
	return false
}

func parseDate(t *tiff.Tag) string {
	inLay := "2006:01:02 15:04:05"
	outLay := "2006-01-02"
	n, _ := time.Parse(inLay, t.String())
	return n.Format(outLay)
}

func getTagValue(t *tiff.Tag) string {
	v, e := t.StringVal()
	if e != nil {
		log.Println(e)
	}
	return v
}

func sanitizeString(fieldname string, t *tiff.Tag) string {
	val := getTagValue(t)
	println(val)
	if contains(dateFields, fieldname) {
		return parseDate(t)
	}
	return val
}

func ParseFormat(str string, t Tags) string {
	r := regexp.MustCompile(`%\w+%`)
	p := r.ReplaceAllStringFunc(str,
		func(s string) string {
			raw := strings.Replace(s, "%", "", 2)
			return sanitizeString(raw, t[raw])
		})
	return p
}
