package test

import (
	"testing"

	"github.com/assac453/some-testing/abs"
)

func TestAbs(t *testing.T) {
	got := abs.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; wand 1", got)
	}
}
