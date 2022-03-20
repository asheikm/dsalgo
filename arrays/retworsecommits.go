/*
You are working on a project and you noticed that there has been a performance decrease between the two releases. You have a function:
boolean worseCommit(int commit1, int commit2)
that runs performance tests and returns true if commit2 is worse than commit1 and false otherwise.
Find all of the bad commits that have decreased the performance between releases. Assume no improvement in performance.
Commit Id: 1, 2, 3, 4, 5, 6, 7, 8, 9
Performance: 10, 10, 10, 8, 8, 8, 5, 5, 5
Output 4, 7
*/
package main

import (
	"fmt"
)

func main() {
	commitIds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	badcommits := allBadCommits(commitIds)
	fmt.Println(badcommits)
}

/* o(n) - complexity */
func allBadCommits(commitIds []int) []int {
	badCommits := make([]int, 0, 0)
	for i := 0; i < len(commitIds); i++ {
		if worseCommit(commitIds[i], commitIds[i+i]) {
			badCommits = append(badCommits, commitIds[i+1])
		}
	}
	return badCommits
}
