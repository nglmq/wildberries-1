package main

import "fmt"

func filterStrings(strings []string) []string {
	var filtered []string

	for _, str := range strings {
		var found bool // flag to check if value already exists

		for _, interStr := range filtered {
			if str == interStr { // if value already in the filtered list, skip it
				found = true
			}
		}

		if !found { // if value not found in the filtered list, add it to the filtered list
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func main() {
	strings := []string{"cat", "dog", "cat", "tree", "cat"} // slice of strings

	filtered := filterStrings(strings)
	fmt.Println(filtered) // [cat dog tree]
}
