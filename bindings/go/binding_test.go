package tree_sitter_javascript_orchard_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_javascript_orchard "codeberg.org/grammar-orchard/tree-sitter-javascript-orchard/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_javascript_orchard.Language())
	if language == nil {
		t.Errorf("Error loading JavaScript grammar")
	}
}
