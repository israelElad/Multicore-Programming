// Rewrite topoSort to use maps instead of slices and eliminate the initial sort. Verify that the results, though nondeterministic, are valid topological orderings
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": false},
	"calculus":   {"linear algebra": false},
	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},
	"data structures":       {"discrete math": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

//!-table

//!+main
func main() {
	var order = topoSort(prereqs)
	for i := 0; i < len(order); i++ {
		fmt.Printf("%d:\t%s\n", i+1, order[i])
	}
}

func topoSort(m map[string]map[string]bool) map[int]string {
	//course num as key, course name as value.
	var order = make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	var i = 0
	visitAll = func(items map[string]bool) {
		for item := range items {

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[i] = item
				i++
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)

	return order
}

//!-main
