package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcabb"
	expected := 3
	actual := lengthOfLongestSubstring(str)
	if expected != actual {
		t.Errorf("The expected length of the longest Substring didn't match the actual.  Expected: %d\nActual: %d", expected, actual)
	}
}
