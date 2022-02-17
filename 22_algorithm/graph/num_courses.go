package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	hasCycle := false
	graph := buildGraph(numCourses, prerequisites)

	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	var traverse func(graph [][]int, s int)
	traverse = func(graph [][]int, s int) {
		if onPath[s] {
			hasCycle = true
		}
		if visited[s] || hasCycle {
			return
		}
		visited[s] = true
		onPath[s] = true
		for _, v := range graph[s] {
			traverse(graph, v)
		}
		onPath[s] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	return !hasCycle
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)

	for _, v := range prerequisites {
		from := v[1]
		to := v[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}

func main() {
	pre := []int{1, 0}
	fmt.Println(canFinish(2, [][]int{pre}))
}
