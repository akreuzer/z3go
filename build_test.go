package z3go_test

import (
	"fmt"
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

func TestZ3goExceptionHandline(t *testing.T) {
	// This is basically the error example from examples/
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	x := c.Bool_const("x")

	defer func() {
		if r := recover(); r == nil {
			t.Error("The C++ exception was not translated into a panic")
		}
	}()

	// The next call fails because x is a Boolean.
	expr := z3.Add(x, 1)
	_ = expr
}

func TestZ3goUnsatCore1Example(t *testing.T) {
	c := z3.NewContext()
	defer z3.DeleteContext(c)

	p1 := c.Bool_const("p1")
	p2 := c.Bool_const("p2")
	p3 := c.Bool_const("p3")
	x := c.Int_const("x")
	y := c.Int_const("y")

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)

	s.Add(z3.Implies(p1, z3.Greater(x, 10))) // p1 => x > 10
	s.Add(z3.Implies(p1, z3.Greater(y, x)))  // p1 => y > x
	s.Add(z3.Implies(p2, z3.Less(y, 5)))     // p2 => y < 5
	s.Add(z3.Implies(p3, z3.Greater(y, 0)))  // p3 => y > 0
	assumptions1 := z3.NewExprVector(c)
	defer z3.DeleteExprVector(assumptions1)
	assumptions1.Push_back(p1)
	assumptions1.Push_back(p2)
	assumptions1.Push_back(p3)
	if s.Check(assumptions1) != z3.Unsat {
		t.Error("p1, p2, p3 together should be unsat.")
	}
	core := s.Unsat_core()
	if core.Size() != 2 {
		t.Error("Core should only contain 2 elements.")
	}
	assumptions2 := z3.NewExprVector(c)
	defer z3.DeleteExprVector(assumptions2)
	assumptions2.Push_back(p1)
	assumptions2.Push_back(p3)
	if s.Check(assumptions2) != z3.Sat {
		t.Error("p1 and p3 should be sat together.")
	}
}

func TestZ3goTypeCanary(t *testing.T) {
	// Check that classes implement Stringer interface
	var e z3.ExprVector
	var _ fmt.Stringer = e

	var s z3.Solver
	var _ fmt.Stringer = s

	var m z3.Model
	var _ fmt.Stringer = m

	var g z3.Goal
	var _ fmt.Stringer = g

	var r z3.Apply_result
	var _ fmt.Stringer = r

	// Check that expr are return of operator functions
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	var x, y z3.Expr
	x = c.Int_const("x")
	y = c.Int_const("y")
	var _ z3.Expr = z3.Add(x, y)
	var _ z3.Expr = z3.Add(x, 1)
	var _ z3.Expr = z3.Add(2, y)
}
