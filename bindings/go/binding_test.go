package tree_sitter_fixed_form_fortran_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_fixed_form_fortran "github.com/tree-sitter/tree-sitter-fixed_form_fortran/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_fixed_form_fortran.Language())
	if language == nil {
		t.Errorf("Error loading FixedFormFortran grammar")
	}
}
