package main

import "testing"

// 1 2 4 2 5 1 3 6 3 7 --- PASS: TestMorris (0.00s)
func TestMorris(t *testing.T) {
	tree := createTest()
	Morris(tree.Root)
}

func TestFirstMorris(t *testing.T) {
	tree := createTest()
	FirstMorris(tree.Root)
}

func TestMiddleMorris(t *testing.T) {
	tree := createTest()
	MiddleMorris(tree.Root)
}

func TestLastMorris(t *testing.T) {
	tree := createTest()
	LastMorris(tree.Root)
}
