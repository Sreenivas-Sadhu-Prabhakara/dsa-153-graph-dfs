package main

import "fmt"

// DFS returns the depth-first visit order over an undirected graph with n
// vertices (0-based), covering every component in vertex order. Neighbors
// are visited in the order edges listed them.
func DFS(n int, edges [][2]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	out := make([]int, 0, n)
	var walk func(v int)
	walk = func(v int) {
		visited[v] = true
		out = append(out, v)
		for _, w := range adj[v] {
			if !visited[w] {
				walk(w)
			}
		}
	}
	for v := 0; v < n; v++ {
		if !visited[v] {
			walk(v) // next unvisited vertex starts a new component
		}
	}
	return out
}

func main() {
	fmt.Println(DFS(5, [][2]int{{0, 1}, {0, 2}, {1, 3}}))
}
