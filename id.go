package nbt

//go:generate stringer -linecomment -type=ID

// ID is a tag id, such as IDTagEnd or IDTagLongArray.
type ID byte

const (
	// IDTagEnd indicates the end for a compound tag.
	IDTagEnd ID = iota // TagEnd

	// IDTagByte is a signed byte (-128 to 127).
	IDTagByte // TagByte

	// IDTagShort is a signed short (-32768 to 32767).
	IDTagShort // TagShort

	// IDTagInt is a signed int (-2147483648 to 2147483647).
	IDTagInt // TagInt

	// IDTagLong is a signed long (-9223372036854775808 to 9223372036854775807).
	IDTagLong // TagLong

	// IDTagFloat is a signed float32 (IEEE-754).
	IDTagFloat // TagFloat

	// IDTagDouble is a signed float64 (IEEE-754).
	IDTagDouble // TagDouble

	// IDTagByteArray is a array of signed bytes.
	IDTagByteArray // TagByteArray

	// IDTagString is a UTF-8 string.
	IDTagString // TagString

	// IDTagList is a list of nameless tags, all tags are of the same type.
	IDTagList // TagList

	// IDTagCompound is a list of named tags.
	IDTagCompound // TagCompound

	// IDTagIntArray is a array for int(4bytes).
	IDTagIntArray // TagIntArray

	// IDTagLongArray is a array for long(8bytes).
	IDTagLongArray // TagLongArray

	// NumIDTags is the amount of ID tags that are known.
	NumIDTags // amount of ID tags
)
