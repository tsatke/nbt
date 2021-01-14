package nbt

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMarshalSuite(t *testing.T) {
	suite.Run(t, new(MarshalSuite))
}

type MarshalSuite struct {
	suite.Suite
}

func (suite *MarshalSuite) expect(t Tag, val interface{}) {
	var buf bytes.Buffer
	suite.NoError(MarshalWriter(&buf, binary.BigEndian, val))
	dec := NewDecoder(&buf, binary.BigEndian)
	tag, err := dec.ReadTag()
	suite.NoError(err)
	suite.equalTag(t, tag)
}

func (suite *MarshalSuite) equalTag(expected, got Tag) {
	if expected == nil || got == nil {
		if expected == got {
			return
		}
		suite.Require().NotNil(expected)
		suite.Require().NotNil(got)
		return // not reached
	}

	suite.Equal(expected.ID(), got.ID())
	suite.Equal(expected.Name(), got.Name())
	switch expected.ID() {
	case IDTagByte:
		suite.Equal(expected.(*Byte).Value, got.(*Byte).Value)
	case IDTagShort:
		suite.Equal(expected.(*Short).Value, got.(*Short).Value)
	case IDTagInt:
		suite.Equal(expected.(*Int).Value, got.(*Int).Value)
	case IDTagIntArray:
		suite.Equal(expected.(*IntArray).Value, got.(*IntArray).Value)
	case IDTagLong:
		suite.Equal(expected.(*Long).Value, got.(*Long).Value)
	case IDTagLongArray:
		suite.Equal(expected.(*LongArray).Value, got.(*LongArray).Value)
	case IDTagList:
		expectedValues := expected.(*List).Value
		gotValues := got.(*List).Value
		for i := range expectedValues {
			suite.equalTag(expectedValues[i], gotValues[i])
		}
	case IDTagCompound:
		expectedMap := expected.(*Compound).Value
		gotMap := got.(*Compound).Value
		for k, v := range expectedMap {
			suite.equalTag(v, gotMap[k])
		}
	case IDTagString:
		suite.Equal(expected.(*String).Value, got.(*String).Value)
	default:
		suite.Failf("unknown tag ID", "tag %s unknown", expected.ID())
	}
}

func (suite *MarshalSuite) TestMarshalWriter_String() {
	suite.expect(NewStringTag("", "hello"), "hello")
}

func (suite *MarshalSuite) TestMarshalWriter_Byte() {
	suite.expect(NewByteTag("", 7), int8(7))
	suite.expect(NewByteTag("", 7), uint8(7))
}

func (suite *MarshalSuite) TestMarshalWriter_Short() {
	suite.expect(NewShortTag("", 7), int16(7))
	suite.expect(NewShortTag("", 7), uint16(7))
}

func (suite *MarshalSuite) TestMarshalWriter_Int() {
	suite.expect(NewIntTag("", 7), int32(7))
	suite.expect(NewIntTag("", 7), uint32(7))
}

func (suite *MarshalSuite) TestMarshalWriter_Long() {
	suite.expect(NewLongTag("", 7), int64(7))
	suite.expect(NewLongTag("", 7), uint64(7))
}

func (suite *MarshalSuite) TestMarshalWriter_IntArray() {
	suite.expect(NewIntArrayTag("", []int32{1, 2, 3}), []int32{1, 2, 3})
	suite.expect(NewIntArrayTag("", []int32{1, 2, 3}), []uint32{1, 2, 3})
}

func (suite *MarshalSuite) TestMarshalWriter_LongArray() {
	suite.expect(NewLongArrayTag("", []int64{1, 2, 3}), []int64{1, 2, 3})
	suite.expect(NewLongArrayTag("", []int64{1, 2, 3}), []uint64{1, 2, 3})
}

func (suite *MarshalSuite) TestMarshalWriter_List() {
	suite.expect(NewListTag("", []Tag{
		NewStringTag("", "a"),
		NewStringTag("", "b"),
	}, IDTagString), []string{"a", "b"})
}

func (suite *MarshalSuite) TestMarshalWriter_Compound() {
	type t struct {
		X string
		Y string
		Z string `nbt:"foobar"`
	}
	suite.expect(NewCompoundTag("", []Tag{
		NewStringTag("X", "a"),
		NewStringTag("Y", "b"),
		NewStringTag("foobar", "c"),
	}), t{
		X: "a",
		Y: "b",
		Z: "c",
	})
}

func (suite *MarshalSuite) TestMarshalWriter_CompoundNested() {
	type n struct {
		X string
		Y string
		Z string `nbt:"foobar"`
	}
	type t struct {
		X string
		Y string
		Z n
	}
	suite.expect(NewCompoundTag("", []Tag{
		NewStringTag("X", "a"),
		NewStringTag("Y", "b"),
		NewCompoundTag("Z", []Tag{
			NewStringTag("X", "a"),
			NewStringTag("Y", "b"),
			NewStringTag("foobar", "c"),
		}),
	}), t{
		X: "a",
		Y: "b",
		Z: n{
			X: "a",
			Y: "b",
			Z: "c",
		},
	})
}

func (suite *MarshalSuite) TestMarshalWriter_CompoundNestedPointer() {
	type n struct {
		X string
		Y string
		Z string `nbt:"foobar"`
	}
	type t struct {
		X string
		Y string
		Z *n
	}
	suite.expect(NewCompoundTag("", []Tag{
		NewStringTag("X", "a"),
		NewStringTag("Y", "b"),
		NewCompoundTag("Z", []Tag{
			NewStringTag("X", "a"),
			NewStringTag("Y", "b"),
			NewStringTag("foobar", "c"),
		}),
	}), t{
		X: "a",
		Y: "b",
		Z: &n{
			X: "a",
			Y: "b",
			Z: "c",
		},
	})
}
