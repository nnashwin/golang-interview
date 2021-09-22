package main

import ( 
	"fmt" 
	"sort"
)

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less (i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap (i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams (ss []string) [][]string {
	m := map[string][]string{}
	for _, s := range ss {
		sorted := SortString(s)
		m[sorted] = append(m[sorted], s)
	}

	var res [][]string

	for _, strSl := range m {
		res = append(res, strSl)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	// should return 3 groups because the words are not anagrams of one another.
	fmt.Println(groupAnagrams([]string{"a", "an", "ant"}))
}