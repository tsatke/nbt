package nbt

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestUnmarshalSuite(t *testing.T) {
	suite.Run(t, new(UnmarshalSuite))
}

type UnmarshalSuite struct {
	suite.Suite

	buf *bytes.Buffer
}

func (suite *UnmarshalSuite) SetupTest() {
	suite.buf = new(bytes.Buffer)
}

func (suite *UnmarshalSuite) writeTag(t Tag, order binary.ByteOrder) {
	suite.NoError(NewEncoder(suite.buf, order).WriteTag(t))
}

func (suite *UnmarshalSuite) TestUnmarshalReader_String() {
	var target string
	suite.writeTag(NewStringTag("myName", "myVal"), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues("myVal", target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Long() {
	var target int64
	suite.writeTag(NewLongTag("myName", 1234567890), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(1234567890, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_LongArray() {
	var target []int64
	suite.writeTag(NewLongArrayTag("myName", []int64{1, 2, 3, 4, 5, 9}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues([]int64{1, 2, 3, 4, 5, 9}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Int() {
	var target int32
	suite.writeTag(NewIntTag("myName", 12345), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(12345, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Short() {
	var target int16
	suite.writeTag(NewShortTag("myName", 12345), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(12345, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_IntArray() {
	var target []int32
	suite.writeTag(NewIntArrayTag("myName", []int32{1, 2, 3, 4, 5, 9}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues([]int32{1, 2, 3, 4, 5, 9}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Byte() {
	var target byte
	suite.writeTag(NewByteTag("myName", 17), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(17, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_ByteArray() {
	var target []int8
	suite.writeTag(NewByteArrayTag("myName", []int8{1, 2, 3, 4, 5, 9}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues([]int8{1, 2, 3, 4, 5, 9}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Double() {
	var target float64
	suite.writeTag(NewDoubleTag("myName", 1.5), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(1.5, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Float() {
	var target float32
	suite.writeTag(NewFloatTag("myName", 1.5), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(1.5, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_Compound() {
	type t struct {
		X string
		Y string
		Z string `nbt:"foobar"`
	}
	var target t
	suite.writeTag(NewCompoundTag("myName", []Tag{
		NewStringTag("X", "xVal"),
		NewStringTag("Y", "yVal"),
		NewStringTag("foobar", "zVal"),
	}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(t{
		X: "xVal",
		Y: "yVal",
		Z: "zVal",
	}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_NestedCompound() {
	type n struct {
		A, B string
	}
	type t struct {
		X string
		Y n
		Z string `nbt:"foobar"`
	}
	var target t
	suite.writeTag(NewCompoundTag("myName", []Tag{
		NewStringTag("X", "xVal"),
		NewCompoundTag("Y", []Tag{
			NewStringTag("A", "aVal"),
			NewStringTag("B", "bVal"),
		}),
		NewStringTag("foobar", "zVal"),
	}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(t{
		X: "xVal",
		Y: n{
			A: "aVal",
			B: "bVal",
		},
		Z: "zVal",
	}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_NestedCompoundPointer() {
	type n struct {
		A, B string
	}
	type t struct {
		X string
		Y *n
		Z string `nbt:"foobar"`
	}
	var target t
	suite.writeTag(NewCompoundTag("myName", []Tag{
		NewStringTag("X", "xVal"),
		NewCompoundTag("Y", []Tag{
			NewStringTag("A", "aVal"),
			NewStringTag("B", "bVal"),
		}),
		NewStringTag("foobar", "zVal"),
	}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(t{
		X: "xVal",
		Y: &n{
			A: "aVal",
			B: "bVal",
		},
		Z: "zVal",
	}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_List() {
	var target []string
	suite.writeTag(NewListTag("myName", []Tag{
		NewStringTag("0", "index0"),
		NewStringTag("1", "index1"),
		NewStringTag("2", "index2"),
	}, IDTagString), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues([]string{
		"index0",
		"index1",
		"index2",
	}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_StructTag_Ignore() {
	type t struct {
		X string
		Y string
		Z string `nbt:"-"`
	}
	var target t
	suite.writeTag(NewCompoundTag("myName", []Tag{
		NewStringTag("X", "xVal"),
		NewStringTag("Y", "yVal"),
		NewStringTag("Z", "zVal"),
	}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(t{
		X: "xVal",
		Y: "yVal",
		Z: "", // must be empty because ignored
	}, target)
}

func (suite *UnmarshalSuite) TestUnmarshalReader_StructTag_Omitempty() {
	type t struct {
		X string
		Y string
		Z string `nbt:"omitempty"`
	}
	var target t
	suite.writeTag(NewCompoundTag("myName", []Tag{
		NewStringTag("X", "xVal"),
		NewStringTag("Y", "yVal"),
		NewStringTag("Z", "zVal"),
	}), binary.BigEndian)
	suite.NoError(UnmarshalReader(suite.buf, binary.BigEndian, &target))
	suite.EqualValues(t{
		X: "xVal",
		Y: "yVal",
		Z: "zVal", // omitempty does't have an effect during unmarshalling
	}, target)
}
