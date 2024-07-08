package example

import (
	"testing"

	"github.com/gpbPiazza/notebook/my_go_playground/playground/archp"
)

func TestDependenciesTest(t *testing.T) {

	// importPath of packages parents
	verifier := archp.NewVerifier("github.com/gpbPiazza/notebook/algo_data/graphs")

	err := verifier.Verify()

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
