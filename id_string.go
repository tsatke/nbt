// Code generated by "stringer -linecomment -type=ID"; DO NOT EDIT.

package nbt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IDTagEnd-0]
	_ = x[IDTagByte-1]
	_ = x[IDTagShort-2]
	_ = x[IDTagInt-3]
	_ = x[IDTagLong-4]
	_ = x[IDTagFloat-5]
	_ = x[IDTagDouble-6]
	_ = x[IDTagByteArray-7]
	_ = x[IDTagString-8]
	_ = x[IDTagList-9]
	_ = x[IDTagCompound-10]
	_ = x[IDTagIntArray-11]
	_ = x[IDTagLongArray-12]
	_ = x[NumIDTags-13]
}

const _ID_name = "TagEndTagByteTagShortTagIntTagLongTagFloatTagDoubleTagByteArrayTagStringTagListTagCompoundTagIntArrayTagLongArrayamount of ID tags"

var _ID_index = [...]uint8{0, 6, 13, 21, 27, 34, 42, 51, 63, 72, 79, 90, 101, 113, 130}

func (i ID) String() string {
	if i >= ID(len(_ID_index)-1) {
		return "ID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ID_name[_ID_index[i]:_ID_index[i+1]]
}
