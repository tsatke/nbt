package nbt

import (
	"fmt"
	"strings"
)

type simpleMapper struct {
	tag Tag
}

func NewSimpleMapper(source Tag) Mapper {
	return &simpleMapper{
		tag: source,
	}
}

func (m *simpleMapper) Query(query string) (Tag, error) {
	if query == "" {
		return m.tag, nil
	}

	frags := strings.Split(query, ".")
	last := frags[len(frags)-1]
	current, ok := m.tag.(*Compound)
	if !ok {
		return nil, fmt.Errorf("root element is not a compound")
	}
	for i, frag := range frags[:len(frags)-1] {
		temp, ok := current.Get(frag)
		if !ok {
			fmt.Println(ToString(m.tag))
			return nil, fmt.Errorf("can't find %s", strings.Join(frags[:i+1], "."))
		}
		compound, ok := temp.(*Compound)
		if !ok {
			return nil, fmt.Errorf("%s is not a compound", strings.Join(frags[:i], "."))
		}
		current = compound
	}

	res, ok := current.Get(last)
	if !ok {
		return nil, fmt.Errorf("can't find %s", query)
	}
	return res, nil
}

func (m *simpleMapper) MapByte(query string, target *int8) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagByte:
		*target = res.(*Byte).Value
	default:
		return fmt.Errorf("%s is not a byte", query)
	}
	return nil
}

func (m *simpleMapper) MapShort(query string, target *int16) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagShort:
		*target = res.(*Short).Value
	default:
		return fmt.Errorf("%s is not a short", query)
	}
	return nil
}

func (m *simpleMapper) MapInt(query string, target *int) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagInt:
		*target = int(res.(*Int).Value)
	default:
		return fmt.Errorf("%s is not an int", query)
	}
	return nil
}

func (m *simpleMapper) MapInt32(query string, target *int32) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagInt:
		*target = res.(*Int).Value
	default:
		return fmt.Errorf("%s is not an int", query)
	}
	return nil
}

func (m *simpleMapper) MapLong(query string, target *int64) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagLong:
		*target = res.(*Long).Value
	default:
		return fmt.Errorf("%s is not a long", query)
	}
	return nil
}

func (m *simpleMapper) MapString(query string, target *string) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagString:
		*target = res.(*String).Value
	default:
		return fmt.Errorf("%s is not a string", query)
	}
	return nil
}

func (m *simpleMapper) MapByteArray(query string, target *[]int8) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagByteArray:
		*target = res.(*ByteArray).Value
	default:
		return fmt.Errorf("%s is not a byte array", query)
	}
	return nil
}

func (m *simpleMapper) MapIntArray(query string, target *[]int) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagIntArray:
		values := res.(*IntArray).Value
		*target = make([]int, len(values))
		for i := range values {
			(*target)[i] = int(values[i])
		}
	default:
		return fmt.Errorf("%s is not an int array", query)
	}
	return nil
}

func (m *simpleMapper) MapInt32Array(query string, target *[]int32) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagIntArray:
		*target = res.(*IntArray).Value
	default:
		return fmt.Errorf("%s is not an int array", query)
	}
	return nil
}

func (m *simpleMapper) MapLongArray(query string, target *[]int64) error {
	res, err := m.Query(query)
	if err != nil {
		return err
	}
	switch res.ID() {
	case IDTagLongArray:
		*target = res.(*LongArray).Value
	default:
		return fmt.Errorf("%s is not a long array", query)
	}
	return nil
}

func (m *simpleMapper) MapList(query string, initializer func(int), mapping func(int, Mapper) error) error {
	val, err := m.Query(query)
	if err != nil {
		return err
	}
	if val.ID() != IDTagList {
		return fmt.Errorf("%s is not a list", query)
	}
	values := val.(*List).Value
	initializer(len(values))
	for i, value := range values {
		if err := mapping(i, NewSimpleMapper(value)); err != nil {
			return err
		}
	}
	return nil
}

func (m *simpleMapper) MapCustom(query string, mapping func(tag Tag) error) error {
	val, err := m.Query(query)
	if err != nil {
		return err
	}
	if err := mapping(val); err != nil {
		return err
	}
	return nil
}
