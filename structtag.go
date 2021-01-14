package nbt

import "strings"

const (
	structTag          = "nbt"
	structTagIgnore    = "-"
	structTagOmitempty = "omitempty"
)

type sTag struct {
	name      string
	ignore    bool
	omitempty bool
}

func parseStructTag(in string) (tag sTag) {
	frags := strings.Split(in, ",")
	for _, frag := range frags {
		switch frag {
		case structTagIgnore:
			tag.ignore = true
		case structTagOmitempty:
			tag.omitempty = true
		default:
			tag.name = frag
		}
	}
	return
}
