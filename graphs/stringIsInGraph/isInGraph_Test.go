package main

import (
	"fmt"
)

func TestBuildLL(t *testing.T) {
	cases := []struct {
		in  []string
		out LLNode
	}{
		[]string{"tyler", "eats", "cabbage", "daily"},
		LLNode{"tyler", LLNode{"eats"}},
	}
}

// func TestIsInGraph(t *testing.T) {
//   cases := []struct {
//     graph 	LLNode
//     str	string
//   }{
//     {
//       LLNode{}
//       "tyler"
//     }
//   }
// }
