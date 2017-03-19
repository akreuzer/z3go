package z3go_test

import (
	"testing"

	z3 "github.com/akreuzer/z3go"
)

func TestZ3go(t *testing.T) {
	// This is basically the deMorgan example from examples/
	c := z3.NewContext()
	defer z3.DeleteContext(c)

	x := c.Bool_const("x")
	y := c.Bool_const("y")

	conjecture := z3.Equals(z3.Not(z3.And(x, y)), z3.Or(z3.Not(x), z3.Not(y)))
	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)
	s.Add(z3.Not(conjecture))

	if s.Check() != z3.Unsat {
		t.Error("Could not validate de Morgan's rule.")
	}
}
