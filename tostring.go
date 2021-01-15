package nbt

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

type lineWriter struct {
	indent string
	out    io.Writer
}

func (w *lineWriter) println(line string) {
	_, _ = io.WriteString(w.out, w.indent+line+"\n")
}

func (w *lineWriter) printSingleValueTag(id ID, name string, value interface{}) {
	w.println(fmt.Sprintf("%s('%s'): %v", id, name, value))
}

func (w *lineWriter) indentUp()   { w.indent += "\t" }
func (w *lineWriter) indentDown() { w.indent = w.indent[:len(w.indent)-1] }

func ToString(tag Tag) string {
	var buf bytes.Buffer
	w := &lineWriter{
		indent: "",
		out:    &buf,
	}
	toString(w, tag)
	return buf.String()
}

func toString(w *lineWriter, tag Tag) {
	switch tag.ID() {
	case IDTagEnd:
		w.println("(//end)")
	case IDTagByte:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Byte).Value))
	case IDTagShort:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Short).Value))
	case IDTagInt:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Int).Value))
	case IDTagLong:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Long).Value))
	case IDTagFloat:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Float).Value))
	case IDTagDouble:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*Double).Value))
	case IDTagString:
		w.println(fmt.Sprintf("%s('%s'): '%s'", tag.ID(), tag.Name(), tag.(*String).Value))
	case IDTagByteArray:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*ByteArray).Value))
	case IDTagIntArray:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*IntArray).Value))
	case IDTagLongArray:
		w.println(fmt.Sprintf("%s('%s'): %v", tag.ID(), tag.Name(), tag.(*LongArray).Value))
	case IDTagList:
		values := tag.(*List).Value
		w.println(fmt.Sprintf("%s('%s'): %d entries", tag.ID(), tag.Name(), len(values)))
		w.println("{")
		w.indentUp()
		for _, value := range values {
			toString(w, value)
		}
		w.indentDown()
		w.println("}")
	case IDTagCompound:
		values := tag.(*Compound).Value
		w.println(fmt.Sprintf("%s('%s'): %d entries", tag.ID(), tag.Name(), len(values)))
		w.println("{")
		w.indentUp()
		// sort keys alphabetically, at least in ToString
		var keys []string
		for k, _ := range values {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			toString(w, values[k])
		}
		w.indentDown()
		w.println("}")
	default:
		panic("unknown tag ID " + tag.ID().String())
	}
}
