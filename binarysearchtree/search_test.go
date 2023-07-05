package binarysearchtree

import "testing"

func TestSearch(t *testing.T) {
	tree := createTree()

	if !tree.Search(6) {
		t.Error("should have 6")
	}

	if tree.Search(37) {
		t.Error("should not have 37")
	}
}
