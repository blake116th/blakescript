package assert

import (
	"fmt"
	"testing"
)

func IntEquals(actual int, expected int, t *testing.T) {
	if (actual != expected) {
		s := fmt.Sprintf("Actual: %d, Expected: %d", actual, expected)
		t.Error(s)
	}
}