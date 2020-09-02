//Use the breadthFirst function to explore a different structure.
//For example, you could use the course dependencies from the topoSort example (a directed graph).

package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

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

//!+breadthFirst from Findlinks3 crawler
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main() {
	for course, course_prereqs := range prereqs {
		fmt.Println(course + " prerequisites:")
		breadthFirst(appendDeps, course_prereqs)
		fmt.Println("*********************")
	}
}

func appendDeps(course string) []string {
	var order []string
	fmt.Println(course)
	for _, item := range prereqs[course] {
		order = append(order, item)
	}
	return order
}

/*
first run:

formal languages prerequisites:
discrete math
intro to programming
*********************
networks prerequisites:
operating systems
data structures
computer organization
discrete math
intro to programming
*********************
operating systems prerequisites:
data structures
computer organization
discrete math
intro to programming
*********************
algorithms prerequisites:
data structures
discrete math
intro to programming
*********************
compilers prerequisites:
data structures
formal languages
computer organization
discrete math
intro to programming
*********************
data structures prerequisites:
discrete math
intro to programming
*********************
databases prerequisites:
data structures
discrete math
intro to programming
*********************
discrete math prerequisites:
intro to programming
*********************
programming languages prerequisites:
data structures
computer organization
discrete math
intro to programming
*********************
calculus prerequisites:
linear algebra
*********************


second run:

operating systems prerequisites:
data structures
computer organization
discrete math
intro to programming
*********************
algorithms prerequisites:
data structures
discrete math
intro to programming
*********************
data structures prerequisites:
discrete math
intro to programming
*********************
databases prerequisites:
data structures
discrete math
intro to programming
*********************
discrete math prerequisites:
intro to programming
*********************
programming languages prerequisites:
data structures
computer organization
discrete math
intro to programming
*********************
calculus prerequisites:
linear algebra
*********************
compilers prerequisites:
data structures
formal languages
computer organization
discrete math
intro to programming
*********************
formal languages prerequisites:
discrete math
intro to programming
*********************
networks prerequisites:
operating systems
data structures
computer organization
discrete math
intro to programming
*********************

as you can see, the prerequisites of each course are the same,
but the order of which the courses appear isn't.
that's because map elements aren't being stored in a fixed order, so iterating over a map return it's elements in a different order each time.
that's why we get different outputs- depend on which courses we pulled first from the map.
*/
