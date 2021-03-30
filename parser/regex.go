package parser

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/tiff"
)

// Valid EXIF fields
var ValidFields = []string{"ImageWidth", "ImageLength", "BitsPerSample", "Compression", "PhotometricInterpretation",
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

func containsTag(t Tags, str string) bool {
	if _, ok := t[str]; ok {
		return true
	}
	return false
}

func parseDate(str string) string {
	inLay := `"2006:01:02 15:04:05"`
	outLay := `2006-01-02`
	n, _ := time.Parse(inLay, str)
	return n.Format(outLay)
}

func hash(t Tags) string {
	var in string
	for _, n := range t {
		in = in + n.String()
	}
	h := md5.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))[0:8]
}

func getValue(fieldname string, t *tiff.Tag) string {
	val := t.String()
	if contains(dateFields, fieldname) {
		return parseDate(val)
	}
	return val
}

func sanitizeString(val string) string {
	val = strings.TrimSuffix(val, `"`)
	val = strings.TrimPrefix(val, `"`)
	val = strings.ReplaceAll(val, ` `, `_`)
	val = strings.ReplaceAll(val, `/`, `_`)
	return val
}

func ParseFormat(str string, t Tags) string {
	r := regexp.MustCompile(`%\w+%`)
	p := r.ReplaceAllStringFunc(str,
		func(s string) string {
			raw := strings.Replace(s, "%", "", 2)
			if !containsTag(t, raw) && !contains(newtags, raw) {
				log.Fatal("Tag '" + raw + "' not present in file EXIF data.")
			} else if raw == "Hash" {
				return sanitizeString(hash(t))
			}
			return sanitizeString(parseNewTags(raw, t))
		})
	return p
}
