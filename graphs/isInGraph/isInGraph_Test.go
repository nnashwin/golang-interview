package main

import (
	"fmt"
	"testing"
)

func TestBuildLL(t *testing.T) {
	input := []string{string('t')}
	cases := []struct {
		in  []string
		out *List
	}{
		{
			input,
			&List{},
		},
	}

	for _, c := range cases {
		var out List
		el := Element{}
		el.Val = string('t')
		out.root = Element{&el, &out, ""}
		got := BuildLL(input)

		if out.root.next.Val != got.root.next.Val {
			t.Errorf("BuildLL(%g) != %g", c.in, got)
		}
	}
}

func TestIsInGraph(t *testing.T) {
	input := []string{"t", "X", "T", "x", "y", "s", "P", "x"}
	l := BuildLL(input)
	cases := []struct {
		graph *List
		str   string
		out   bool
	}{
		{
			l,
			"tyler",
			false,
		},

		{
			l,
			"cookies",
			false,
		},

		{
			l,
			"tx",
			true,
		},

		{
			l,
			"texas",
			false,
		},

		{
			l,
			"",
			false,
		},
	}

	for _, c := range cases {
		expected := c.out
		got := IsInGraph(c.str, c.graph)
		fmt.Println(expected)
		fmt.Println(got)

		if expected != got {
			t.Error("string is not contained within the graph")
		}
	}
}
