package main

import (
	"reflect"
	"testing"
)

func TestCreateMap(t *testing.T) {
	cases := []struct {
		in  string
		out map[string]float64
	}{
		{
			"tyler",
			map[string]float64{
				"t": 1,
				"y": 1,
				"l": 1,
				"e": 1,
				"r": 1,
			},
		},
	}

	for _, c := range cases {
		got := CreateMap(c.in)
		if reflect.DeepEqual(got, c.out) != true {
			t.Errorf("CreateMap(%q) == %q, not %q", c.in, got, c.out)
		}
	}
}

func TestStrDiff(t *testing.T) {
	cases := []struct {
		str1, str2 string
		diff       float64
	}{
		{"tyler", "tyl", 2},
		{"", "", 0},
	}

	for _, c := range cases {
		got := StrDiff(c.str1, c.str2)
		if got != c.diff {
			t.Errorf("StrDiff(%q, %q) == %q, not %q", c.str1, c.str2, got, c.diff)
		}
	}
}
