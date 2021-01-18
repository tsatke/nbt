// Package nbt provides encoding and decoding functionality for minecraft's NBT format.
// This package implements support for both BigEndian and SmallEndian, i.e. for every
// binary.ByteOrder.
//
// To marshal a struct, use the marshal function as follows.
//
//	type myStruct struct {
//		X int `nbt:"x"`
//	}
//	...
//	_ = nbt.MarshalWriter(myWriter, binary.BigEndian, myStruct{5})
//
// The above example will write the following as NBT.
//
//	TagCompound(''): 1 entries
//	{
//		TagInt('x'): 5
//	}
//
// You can also manually encode a tag, as follows.
//
//	enc := nbt.NewEncoder(myWriter, binary.BigEndian)
//		_ = enc.WriteTag(nbt.NewCompoundTag("", []Tag{
//			nbt.NewIntTag("x", 5),
//		}),
//	)
//
// This will result in the same output on the writer as in the example above.
// To read this written NBT tag, there are a few options, the easiest one
// being unmarshalling.
//
//	var v myStruct
//	_ = nbt.UnmarshalReader(myReader, binary.BigEndian, &v)
//	// v == 5
//
// For (un-)marshalling, there is a struct tag, which supports naming, '-' (ignore while marshalling and unmarshalling)
// and 'omitempty', which ignores zero values while marshalling.
// For reading tags one by one from a reader, the process is similar to encoding.
//
//	dec := NewDecoder(myReader, binary.BigEndian)
//	tag, err := dec.ReadTag()
//	fmt.Println(nbt.ToString(tag))
//
// This will print whatever one tag was on the reader.
// nbt.ToString will print an NBT tag in the above used representation.
//
// Another way of decoding is using the nbt.Mapper. It is designed to unmarshal
// larger structures, for which unmarshalling is not flexible enough. It is more work
// than unmarshalling, but provides a clean and simple API for decoding. The idea is based
// on XPath, and a simple version of that is currently implemented.
//
//	TagCompound(''): 1 entries
//	{
//		TagCompound('first'): 1 entries
//		{
//			TagInt('x'): 5
//		}
//	}
//
// Using this NBT structure, the following code will extract the value 5 and finally print it.
//
//	var myInt int
//	mapper := nbt.NewSimpleMapper(myTag)
//	_ = mapper.MapInt("first.x", &myInt)
//	fmt.Println(myInt)
//
// This works for all NBT data types, including arrays and lists. For lists, the mapping function takes
// a function that is called before any mapping is done with the size of the list, which allows the user
// to preallocate a slice or similar. The following code decodes a list of int tags into an int array.
//
//	var myInts []int
//	_ = mapper.MapList("path.to.intlist", func(size int) {
//		myInts = make([]int, size)
//	}, func(i int, mapper Mapper) error {
//		return mapper.MapInt(&myInts[i])
//	})
//
// Any error returned will contain a detailed message, what caused the error. Examples are, that the root
// tag or any tag in the query path except the last element is not a compound, the query path does not
// exist, or a type didn't match.
package nbt
