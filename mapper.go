package nbt

// Interface because there will be faster implementations than the most intuitive, probably slow one,
// but they're gonna be very memory expensive and I can't decide whether that tradeoff shouldn't be
// something that the user must decide.
// Also, I know that this interface is a bad abstraction, but I need this interface for a facade, not
// an abstraction.

// Mapper describes a component that can be used to support more complex
// mappings of known structures than marshalling could support.
type Mapper interface {
	// Query will execute the given query string on the tag in this mapper.
	// The interpretation of the query is implementation specific.
	Query(string) (Tag, error)
	// MapByte will interpret the tag under the given query path as byte and
	// store it under the given *int8, or return an error if the tag under the
	// path is not a byte tag.
	MapByte(string, *int8) error
	// MapShort will interpret the tag under the given query path as short and
	// store it under the given *int16, or return an error if the tag under the
	// path is not a short tag.
	MapShort(string, *int16) error
	// MapInt will interpret the tag under the given query path as int and
	// store it under the given *int, or return an error if the tag under the
	// path is not an int tag.
	MapInt(string, *int) error
	// MapInt32 works just as MapInt, but converts the given int to an int32.
	// This can be useful if you have a lot of ints (which are 4-byte in NBT),
	// but want to save the extra 4 bytes if you're on a 64bit arch.
	MapInt32(string, *int32) error
	// MapLong will interpret the tag under the given query path as int64 and
	// store it under the given *int64, or return an error if the tag under the
	// path is not a long tag.
	MapLong(string, *int64) error
	// MapFloat will interpret the tag under the given query path as float32 and
	// store it under the given *float32, or return an error if the tag under the
	// path is not a long tag.
	MapFloat(string, *float32) error
	// MapDouble will interpret the tag under the given query path as float64 and
	// store it under the given *float64, or return an error if the tag under the
	// path is not a long tag.
	MapDouble(string, *float64) error
	// MapString will interpret the tag under the given query path as string and
	// store it under the given *string, or return an error if the tag under the
	// path is not a string tag.
	MapString(string, *string) error
	// MapByteArray will interpret the tag under the given query path as bytearray and
	// store it under the given *[]int8, or return an error if the tag under the
	// path is not a bytearray tag.
	MapByteArray(string, *[]int8) error
	// MapIntArray will interpret the tag under the given query path as intarray and
	// store it under the given *[]int, or return an error if the tag under the
	// path is not a intarray tag.
	MapIntArray(string, *[]int) error
	// MapInt32Array is the array equivalent to MapInt32.
	MapInt32Array(string, *[]int32) error
	// MapLongArray will interpret the tag under the given query path as longarray and
	// store it under the given *[]int64, or return an error if the tag under the
	// path is not a longarray tag.
	MapLongArray(string, *[]int64) error
	// MapList will interpret the tag under the given query path as list. It will return
	// an error if that tag is not a list (also returns an error if the path points to
	// an array). Before calling the mapping function, it will call the initializer function
	// once with the size of the list, allowing for preallocation. After that, it will call
	// the mapping function for every element in the list, with the mapper containing only
	// the list element at index i. i is the zero-based index of an element in the list.
	MapList(query string, initializer func(int), mapping func(i int, mapper Mapper) error) error
	// MapCustom is equivalent to calling Query, and then calling the given function with the tag
	// udner the query, or return an error if any.
	MapCustom(string, func(Tag) error) error
}
