package assert

import (
	"fmt"
	"testing"

	"github.com/Appleby43/blakescript/ast"
)

//returns true if the test fails
func IntEquals(actual int, expected int, t *testing.T) bool {
	if (actual != expected) {
		s := fmt.Sprintf("Actual: %d, Expected: %d", actual, expected)
		t.Error(s)
		return true
	}
	return false
}

func StringEquals(actual string, expected string, t *testing.T) bool {
	if (actual != expected) {
		s := fmt.Sprintf("Actual: %s, Expected: %s", actual, expected)
		t.Error(s)
		return true
	}
	return false
}

func ScriptEquals(actual ast.BlakeScript, expected ast.BlakeScript, t *testing.T) {
	
} 