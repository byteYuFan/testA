package main

import "testing"

// 1 2 4 5 3 6 7 --- PASS: TestFirstNotRecurrence (0.00s)
func TestFirstNotRecurrence(t *testing.T) {
	tree := createTest()
	FirstRecurrence(tree.Root)
}

// 4 2 5 1 6 3 7 --- PASS: TestMiddleNotRecurrence (0.00s)
func TestMiddleNotRecurrence(t *testing.T) {
	tree := createTest()
	MiddleNotRecurrence(tree.Root)
}

// 4 5 2 6 7 3 1 --- PASS: TestLastNotRecurrence (0.00s)
func TestLastNotRecurrence(t *testing.T) {
	tree := createTest()
	LastNotRecurrence(tree.Root)
}
