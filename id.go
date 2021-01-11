package nbt

//go:generate stringer -linecomment -type=ID

type ID byte

const (
	// IDTagEnd signifies end for compound tag
	// Payload: none
	IDTagEnd ID = iota // TagEnd

	// IDTagByte is a signed byte (-128 to 127)
	// Payload is 1byte
	IDTagByte // TagByte

	// IDTagShort is a signed short (-32768 to 32767)
	// Payload is 2 bytes
	IDTagShort // TagShort

	// IDTagInt is a signed int (-2147483648 to 2147483647)
	// Payload is 4 bytes
	IDTagInt // TagInt

	// IDTagLong is a signed long (-9223372036854775808 to 9223372036854775807)
	// Payload is 8 bytes
	IDTagLong // TagLong

	// IDTagFloat is a signed single float (IEEE-754)
	// Payload is 4 bytes
	IDTagFloat // TagFloat

	// IDTagDouble is a signed single double (IEEE-754)
	// Payload is 8 bytes
	IDTagDouble // TagDouble

	// IDTagByteArray is a array of signed bytes
	// Payload is 4 bytes(len of data with signed int) + len bytes
	IDTagByteArray // TagByteArray

	// IDTagString is a UTF-8 string (max 32767 bytes)
	// Payload is 2 bytes (len of data with short) + len bytes
	IDTagString // TagString

	// IDTagList is a list of nameless tags, all tags need same type.
	// Payload is 1 byte(tag type with byte) + 4 bytes(len with signed int) + len bytes
	IDTagList // TagList

	// IDTagCompound is a list of named tags.
	IDTagCompound // TagCompound

	// IDTagIntArray is a array for int(4bytes).
	IDTagIntArray // TagIntArray

	// IDTagLongArray is a array for long(8bytes).
	IDTagLongArray // TagLongArray

	NumIDTags // amount of ID tags
)
