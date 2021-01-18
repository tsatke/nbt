package nbt

// Interface because there will be faster implementations than the most intuitive, probably slow one,
// but they're gonna be very memory expensive and I can't decide whether that tradeoff shouldn't be
// something that the user must decide.
// Also, I know that this interface is a bad abstraction, but I need this interface for a facade, not
// an abstraction.

// Mapper describes a component that can be used to support more complex
// mappings of known structures than marshalling could support.
type Mapper interface {
	Query(string) (Tag, error)
	MapByte(string, *int8) error
	MapInt(string, *int) error
	MapInt32(string, *int32) error
	MapLong(string, *int64) error
	MapString(string, *string) error
	MapByteArray(string, *[]int8) error
	MapShort(string, *int16) error
	MapIntArray(string, *[]int) error
	MapInt32Array(string, *[]int32) error
	MapLongArray(string, *[]int64) error
	MapList(string, func(int), func(int, Mapper) error) error
	// no MapCompound, since that is represented via queries
	MapCustom(string, func(Tag) error) error
}
