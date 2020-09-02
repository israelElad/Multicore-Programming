//Exercise 5.11: The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend the topoSort function to report cycles.

package main

import (
	"fmt"
	"sort"
	"strings"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, prev []string)
	visitAll = func(items []string, prev []string) {
		for _, item := range items {

			for _, prev_item := range prev {
				if item == prev_item { //cycle found
					prev_as_str := strings.Join(prev, "-->")
					fmt.Println("Cycle found! previous courses: " + prev_as_str + ". current: " + item)
				}
			}

			if !seen[item] {
				seen[item] = true
				visitAll(m[item], append(prev, item))
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys, nil)
	return order
}
