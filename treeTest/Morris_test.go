package main

import "testing"

// 1 2 4 2 5 1 3 6 3 7 --- PASS: TestMorris (0.00s)
func TestMorris(t *testing.T) {
	tree := createTest()
	Morris(tree.Root)
}

/*
=== RUN   TestFirstMorris
1 2 4 5 3 6 7 --- PASS: TestFirstMorris (0.00s)
PASS
*/
func TestFirstMorris(t *testing.T) {
	tree := createTest()
	FirstMorris(tree.Root)
}

/*
=== RUN   TestMiddleMorris
4 2 5 1 6 3 7 --- PASS: TestMiddleMorris (0.00s)
PASS
*/
func TestMiddleMorris(t *testing.T) {
	tree := createTest()
	MiddleMorris(tree.Root)
}

/*
=== RUN   TestLastMorris
4 5 2 6 7 3 1
--- PASS: TestLastMorris (0.00s)
*/
func TestLastMorris(t *testing.T) {
	tree := createTest()
	LastMorris(tree.Root)
}
