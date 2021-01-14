package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

func MarshalWriter(w io.Writer, order binary.ByteOrder, val interface{}) error {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		return MarshalWriter(w, order, value.Elem())
	}

	return marshalFrom(w, order, value)
}

func marshalFrom(w io.Writer, order binary.ByteOrder, value reflect.Value) error {
	enc := NewEncoder(w, order)

	tag, err := createTag(value)
	if err != nil {
		return err
	}
	if err := enc.WriteTag(tag); err != nil {
		return err
	}
	return nil
}

func createTag(value reflect.Value) (Tag, error) {
	var tag Tag
	switch value.Kind() {
	case reflect.String:
		tag = NewStringTag("", value.String())
	case reflect.Int8:
		tag = NewByteTag("", int8(value.Int()))
	case reflect.Uint8:
		tag = NewByteTag("", int8(value.Uint()))
	case reflect.Int16:
		tag = NewShortTag("", int16(value.Int()))
	case reflect.Uint16:
		tag = NewShortTag("", int16(value.Uint()))
	case reflect.Int32:
		tag = NewIntTag("", int32(value.Int()))
	case reflect.Uint32:
		tag = NewIntTag("", int32(value.Uint()))
	case reflect.Int64:
		tag = NewLongTag("", value.Int())
	case reflect.Uint64:
		tag = NewLongTag("", int64(value.Uint()))
	case reflect.Ptr:
		return createTag(value.Elem())
	case reflect.Struct:
		tag = NewCompoundTag("", []Tag{})

		typ := value.Type()
		for i := 0; i < typ.NumField(); i++ {
			typeField := typ.Field(i)
			field := value.Field(i)
			name := typeField.Name
			tagValue := parseStructTag(typeField.Tag.Get(structTag))
			if tagValue.ignore {
				continue
			} else if field.IsZero() && tagValue.omitempty {
				continue
			} else if tagValue.name != "" {
				name = tagValue.name
			}
			created, err := createTag(field)
			if err != nil {
				return nil, err
			}
			created.SetName(name)
			tag.(*Compound).Value[name] = created
		}
	case reflect.Slice:
		switch value.Type().Elem().Kind() {
		case reflect.Int32:
			slice := value.Interface().([]int32)
			return NewIntArrayTag("", slice), nil
		case reflect.Int64:
			slice := value.Interface().([]int64)
			return NewLongArrayTag("", slice), nil
		case reflect.Uint32:
			slice := value.Interface().([]uint32)
			conv := make([]int32, len(slice))
			for i := 0; i < len(conv); i++ {
				conv[i] = int32(slice[i])
			}
			return NewIntArrayTag("", conv), nil
		case reflect.Uint64:
			slice := value.Interface().([]uint64)
			conv := make([]int64, len(slice))
			for i := 0; i < len(conv); i++ {
				conv[i] = int64(slice[i])
			}
			return NewLongArrayTag("", conv), nil
		default:
			var tags []Tag
			for i := 0; i < value.Len(); i++ {
				created, err := createTag(value.Index(i))
				if err != nil {
					return nil, err
				}
				tags = append(tags, created)
			}
			if len(tags) == 0 {
				return nil, fmt.Errorf("empty slice unhandled")
			}
			return NewListTag("", tags, tags[0].ID()), nil
		}
	default:
		return nil, fmt.Errorf("unhandled type %s", value.Type().String())
	}
	return tag, nil
}
