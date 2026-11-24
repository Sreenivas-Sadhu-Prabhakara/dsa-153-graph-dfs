package main

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	// 0-1, 0-2, 1-3; vertex 4 isolated.
	got := DFS(5, [][2]int{{0, 1}, {0, 2}, {1, 3}})
	want := []int{0, 1, 3, 2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DFS = %v, want %v", got, want)
	}

	// Cycle must not revisit.
	got = DFS(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	if !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("cycle DFS = %v", got)
	}

	if got := DFS(0, nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("empty graph = %v", got)
	}
	// Edgeless graph: every vertex its own component.
	if got := DFS(3, nil); !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("edgeless = %v", got)
	}
}
