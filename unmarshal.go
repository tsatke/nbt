package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

func UnmarshalReader(rd io.Reader, order binary.ByteOrder, v interface{}) error {
	dec := NewDecoder(rd, order)
	tag, err := dec.ReadTag()
	if err != nil {
		return fmt.Errorf("read tag: %w", err)
	}
	return unmarshalInto(tag, reflect.ValueOf(v).Elem())
}

func unmarshalInto(tag Tag, target reflect.Value) error {
	switch tag.ID() {
	case IDTagByte:
		target.SetUint(uint64(tag.(*Byte).Value))
	case IDTagByteArray:
		source := tag.(*ByteArray).Value
		newTarget := reflect.MakeSlice(target.Type(), len(source), len(source))
		for i := 0; i < newTarget.Len(); i++ {
			newTarget.Index(i).SetUint(uint64(source[i]))
		}
		target.Set(newTarget)
	case IDTagShort:
		target.SetInt(int64(tag.(*Short).Value))
	case IDTagInt:
		target.SetInt(int64(tag.(*Int).Value))
	case IDTagIntArray:
		source := tag.(*IntArray).Value
		newTarget := reflect.MakeSlice(target.Type(), len(source), len(source))
		for i := 0; i < newTarget.Len(); i++ {
			newTarget.Index(i).SetInt(int64(source[i]))
		}
		target.Set(newTarget)
	case IDTagLong:
		target.SetInt(tag.(*Long).Value)
	case IDTagLongArray:
		source := tag.(*LongArray).Value
		newTarget := reflect.MakeSlice(target.Type(), len(source), len(source))
		for i := 0; i < newTarget.Len(); i++ {
			newTarget.Index(i).SetInt(source[i])
		}
		target.Set(newTarget)
	case IDTagFloat:
		target.SetFloat(float64(tag.(*Float).Value))
	case IDTagDouble:
		target.SetFloat(tag.(*Double).Value)
	case IDTagString:
		target.SetString(tag.(*String).Value)
	case IDTagCompound:
		values := tag.(*Compound).Value
		targetType := target.Type()
		for i := 0; i < targetType.NumField(); i++ {
			typeField := targetType.Field(i)
			field := target.Field(i)
			actualName := typeField.Name
			tagValue := parseStructTag(typeField.Tag.Get(structTag))
			if tagValue.ignore {
				continue
			} else if tagValue.name != "" {
				actualName = tagValue.name
			}

			actualField := field
			if field.Kind() == reflect.Ptr {
				actualField.Set(reflect.New(field.Type().Elem()))
				actualField = actualField.Elem()
			}
			if err := unmarshalInto(values[actualName], actualField); err != nil {
				return err
			}
		}
	case IDTagList:
		source := tag.(*List).Value
		newTarget := reflect.MakeSlice(target.Type(), len(source), len(source))
		for i := 0; i < newTarget.Len(); i++ {
			if err := unmarshalInto(source[i], newTarget.Index(i)); err != nil {
				return err
			}
		}
		target.Set(newTarget)
	default:
		return fmt.Errorf("unsupported type %s", tag.ID())
	}
	return nil
}
