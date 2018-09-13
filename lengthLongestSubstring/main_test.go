package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		expected     int
		stringToTest string
	}{
		{
			3,
			"abcabc",
		},
		{
			4,
			"pwwkewppooo",
		},
		{
			4,
			"pokeppp",
		},
		{
			1,
			"aaaaaa",
		},
		{
			3,
			"dvdk",
		},
	}

	for _, c := range cases {
		actual := lengthOfLongestSubstring(c.stringToTest)
		if c.expected != actual {
			t.Errorf("The expected length of the longest Substring didn't match the actual.  Expected: %d\nActual: %d", c.expected, actual)
		}
	}
}
