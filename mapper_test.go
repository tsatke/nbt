package nbt

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSimpleMapper(t *testing.T) {
	suite.Run(t, &MapperSuite{
		gen: NewSimpleMapper,
	})
}

type MapperSuite struct {
	suite.Suite

	gen func(Tag) Mapper
}

func (suite *MapperSuite) TestMapNonCompoundRoots() {
	suite.Run("tag=byte", func() {
		tag := NewByteTag("myName", -17)
		var value int8
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapByte("", &value))
		suite.Equal(int8(-17), value)
	})
	suite.Run("tag=int", func() {
		tag := NewIntTag("myName", -17)
		var value int
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapInt("", &value))
		suite.Equal(-17, value)
	})
	suite.Run("tag=int32", func() {
		tag := NewIntTag("myName", -17)
		var value int32
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapInt32("", &value))
		suite.Equal(int32(-17), value)
	})
	suite.Run("tag=long", func() {
		tag := NewLongTag("myName", -17)
		var value int64
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapLong("", &value))
		suite.Equal(int64(-17), value)
	})
	suite.Run("tag=string", func() {
		tag := NewStringTag("myName", "myValue")
		var value string
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapString("", &value))
		suite.Equal("myValue", value)
	})
	suite.Run("tag=bytearray", func() {
		tag := NewByteArrayTag("myName", []int8{-2, -1, 0, 1, 2, 3})
		var value []int8
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapByteArray("", &value))
		suite.Equal([]int8{-2, -1, 0, 1, 2, 3}, value)
	})
	suite.Run("tag=intarray", func() {
		tag := NewIntArrayTag("myName", []int32{-2, -1, 0, 1, 2, 3})
		var value []int
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapIntArray("", &value))
		suite.Equal([]int{-2, -1, 0, 1, 2, 3}, value)
	})
	suite.Run("tag=int32array", func() {
		tag := NewIntArrayTag("myName", []int32{-2, -1, 0, 1, 2, 3})
		var value []int32
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapInt32Array("", &value))
		suite.Equal([]int32{-2, -1, 0, 1, 2, 3}, value)
	})
	suite.Run("tag=longarray", func() {
		tag := NewLongArrayTag("myName", []int64{-2, -1, 0, 1, 2, 3})
		var value []int64
		mapper := suite.gen(tag)
		suite.NoError(mapper.MapLongArray("", &value))
		suite.Equal([]int64{-2, -1, 0, 1, 2, 3}, value)
	})
}

func (suite *MapperSuite) TestCompoundRoot() {
	tag := NewCompoundTag("someName", []Tag{
		NewCompoundTag("first", []Tag{
			NewStringTag("sub1", "value1"),
			NewIntTag("sub2", 182763422),
			NewLongTag("sub3", -76543567436213),
			NewByteArrayTag("sub4", []int8{-2, -1, 0, 1, 2, 3}),
		}),
		NewCompoundTag("second", []Tag{
			NewCompoundTag("a", []Tag{
				NewCompoundTag("b", []Tag{
					NewCompoundTag("c", []Tag{
						NewCompoundTag("d", []Tag{
							NewCompoundTag("e", []Tag{
								NewCompoundTag("f", []Tag{
									NewShortTag("foo", -32768),
								}),
							}),
						}),
					}),
				}),
			}),
			NewListTag("myList", []Tag{
				NewIntTag("", 0),
				NewIntTag("", 1),
				NewIntTag("", 2),
				NewIntTag("", 3),
				NewIntTag("", 4),
				NewIntTag("", 5),
				NewIntTag("", 6),
				NewIntTag("", 7),
				NewIntTag("", 8),
				NewIntTag("", 9),
			}, IDTagInt),
		}),
	})
	mapper := suite.gen(tag)
	var (
		firstSub1 string
		firstSub2 int
		firstSub3 int64
		firstSub4 []int8
		foo       int16
		arr       []int
	)
	suite.NoError(mapper.MapString("first.sub1", &firstSub1))
	suite.NoError(mapper.MapInt("first.sub2", &firstSub2))
	suite.NoError(mapper.MapLong("first.sub3", &firstSub3))
	suite.NoError(mapper.MapByteArray("first.sub4", &firstSub4))
	suite.NoError(mapper.MapShort("second.a.b.c.d.e.f.foo", &foo))
	suite.NoError(mapper.MapList("second.myList", func(size int) {
		arr = make([]int, size)
	}, func(i int, mapper Mapper) error {
		return mapper.MapInt("", &arr[i])
	}))
	suite.Equal("value1", firstSub1)
	suite.Equal(182763422, firstSub2)
	suite.Equal(int64(-76543567436213), firstSub3)
	suite.Equal([]int8{-2, -1, 0, 1, 2, 3}, firstSub4)
	suite.Equal(int16(-32768), foo)
	suite.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr)
}
