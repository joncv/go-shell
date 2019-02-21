package strings

import "strings"

var StuMap = map[string]interface{}{
	"Builder":  strings.Builder{},
	"Reader":   strings.Reader{},
	"Replacer": strings.Replacer{},
}

var FucMap = map[string]interface{}{
	"ToUpperSpecial": strings.ToUpperSpecial,
	"IndexByte":      strings.IndexByte,
	"IndexAny":       strings.IndexAny,
	"ToLowerSpecial": strings.ToLowerSpecial,
	"HasSuffix":      strings.HasSuffix,
	"ToLower":        strings.ToLower,
	"ToTitle":        strings.ToTitle,
	"LastIndexFunc":  strings.LastIndexFunc,
	"Index":          strings.Index,
	"ContainsAny":    strings.ContainsAny,
	"Replace":        strings.Replace,
	"SplitN":         strings.SplitN,
	"TrimSpace":      strings.TrimSpace,
	"TrimSuffix":     strings.TrimSuffix,
	"Fields":         strings.Fields,
	"TrimFunc":       strings.TrimFunc,
	"TrimLeft":       strings.TrimLeft,
	"TrimPrefix":     strings.TrimPrefix,
	"EqualFold":      strings.EqualFold,
	"Split":          strings.Split,
	"Contains":       strings.Contains,
	"SplitAfterN":    strings.SplitAfterN,
	"TrimLeftFunc":   strings.TrimLeftFunc,
	"TrimRight":      strings.TrimRight,
	"TrimRightFunc":  strings.TrimRightFunc,
	"Compare":        strings.Compare,
	"IndexRune":      strings.IndexRune,
	"HasPrefix":      strings.HasPrefix,
	"ToTitleSpecial": strings.ToTitleSpecial,
	"NewReplacer":    strings.NewReplacer,
	"LastIndex":      strings.LastIndex,
	"FieldsFunc":     strings.FieldsFunc,
	"NewReader":      strings.NewReader,
	"SplitAfter":     strings.SplitAfter,
	"IndexFunc":      strings.IndexFunc,
	"LastIndexAny":   strings.LastIndexAny,
	"ContainsRune":   strings.ContainsRune,
	"LastIndexByte":  strings.LastIndexByte,
	"Join":           strings.Join,
	"Map":            strings.Map,
	"Count":          strings.Count,
	"Repeat":         strings.Repeat,
	"ToUpper":        strings.ToUpper,
	"Title":          strings.Title,
	"Trim":           strings.Trim,
}
