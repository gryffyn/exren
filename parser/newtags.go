package parser

import (
	"strings"
)

var newtags = []string{"ER_Make", "ER_Model", "ER_Hash"}

func formatERTag(str string) string {
	m := strings.ToTitle(str)
	ma := strings.Split(m, " ")
	return ma[len(ma)-1]
}

func parseNewTags(fieldname string, t Tags) string {
	if fieldname == "ER_Make" {
		return formatERTag(t[fieldname].String())
	} else if fieldname == "ER_Model" {
		return formatERTag(t[fieldname].String())
	} else if fieldname == "ER_Hash" {
		return hash(t)
	}
	return getValue(fieldname, t[fieldname])
}
