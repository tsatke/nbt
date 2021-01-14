package nbt

import (
	"reflect"
	"testing"
)

func Test_parseStructTag(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantTag sTag
	}{
		{
			"empty",
			"",
			sTag{},
		},
		{
			"name",
			"foobar",
			sTag{
				name: "foobar",
			},
		},
		{
			"ignore",
			"-",
			sTag{
				ignore: true,
			},
		},
		{
			"name ignore",
			"foobar,-",
			sTag{
				name:   "foobar",
				ignore: true,
			},
		},
		{
			"ignore name",
			"-,foobar",
			sTag{
				name:   "foobar",
				ignore: true,
			},
		},
		{
			"name ignore omitempty",
			"foobar,-,omitempty",
			sTag{
				name:      "foobar",
				ignore:    true,
				omitempty: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTag := parseStructTag(tt.in); !reflect.DeepEqual(gotTag, tt.wantTag) {
				t.Errorf("parseStructTag() = %v, want %v", gotTag, tt.wantTag)
			}
		})
	}
}
