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
		var out *List
		el := Element{}
		el.Val = 't'
		out.root = Element{&el, out, ""}
		got := BuildLL(input)
		if out.root.next.Val != got.root.next.Val {
			t.Errorf("BuildLL(%g) != %g", c.in, got)
		}
	}

	fmt.Printf("%+v", cases)
}

// func TestIsInGraph(t *testing.T) {
//   cases := []struct {
//     graph 	LLNode
//     str	string
//   }{
//     {
//       LLNode{"}
//       "tyler"
//     }
//   }
// }
