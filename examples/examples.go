// This package contains translation of examples in
// https://github.com/Z3Prover/z3/blob/master/examples/c%2B%2B/example.cpp
package main

import (
	"fmt"

	z3 "github.com/akreuzer/z3go"
)

func deMorgan() {
	fmt.Println("de-Morgan example")

	c := z3.NewContext()
	defer z3.DeleteContext(c)

	x := c.Bool_const("x")
	y := c.Bool_const("y")

	conjecture := z3.Equals(z3.Not(z3.And(x, y)), z3.Or(z3.Not(x), z3.Not(y)))
	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)
	s.Add(z3.Not(conjecture))

	fmt.Println(s)
	fmt.Println(s.To_smt2())

	switch s.Check() {
	case z3.Unsat:
		fmt.Println("de-Morgan is valid")
	case z3.Sat:
		fmt.Println("de-Morgan is not valid")
	case z3.Unknown:
		fmt.Println("Unknown")
	}
}

func findModelExample1() {
	fmt.Println("find_model_example1")

	c := z3.NewContext()
	defer z3.DeleteContext(c)

	x := c.Int_const("x")
	y := c.Int_const("y")

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)

	s.Add(z3.GreaterEq(x, 1))
	s.Add(z3.Less(y, z3.Add(x, 3)))
	fmt.Println(s.Check())

	m := s.Get_model()

	fmt.Println(m)
	for i := 0; i < int(m.Size()); i++ {
		v := m.Get(i)
		if v.Arity() != 0 {
			fmt.Println("This should not happened. This problems does only contain constants.")
		}
		fmt.Printf("%v = %v\n", v.Name().Str(), m.Get_const_interp(v).Get_numeral_int())
	}
}

func proveExample1() {
	fmt.Println("prove_example1")
	c := z3.NewContext()
	defer z3.DeleteContext(c)

	x := c.Int_const("x")
	y := c.Int_const("y")
	intSort := c.Int_sort()
	g := z3.Function("g", intSort, intSort)

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)

	conjecture1 := z3.Implies(z3.Equals(x, y), z3.Equals(g.ApplyFct(x), g.ApplyFct(y)))
	fmt.Printf("conjecture1\n%v\n", conjecture1)
	s.Add(z3.Not(conjecture1))
	if s.Check() == z3.Unsat {
		fmt.Println("proved")
	} else {
		fmt.Println("failed to prove")
	}
	s.Reset()

	conjecture2 := z3.Implies(z3.Equals(x, y), z3.Equals(g.ApplyFct(g.ApplyFct(x)), g.ApplyFct(y)))
	fmt.Printf("conjecture2\n%v\n", conjecture2)
	s.Add(z3.Not(conjecture2))
	if s.Check() == z3.Unsat {
		fmt.Println("proved")
	} else {
		fmt.Println("failed to prove")
		m := s.Get_model()
		fmt.Printf("counterexample: %v\n", m)
		fmt.Printf("g(g(x)) = %v\n", m.Eval(g.ApplyFct(g.ApplyFct(x))))
		fmt.Printf("g(y)    = %v\n", m.Eval(g.ApplyFct(y)))
	}
}

func proveExample2() {
	fmt.Println("prove_example2")
	c := z3.NewContext()
	defer z3.DeleteContext(c)

	x := c.Int_const("x")
	y := c.Int_const("y")
	z := c.Int_const("z")
	intSort := c.Int_sort()
	g := z3.Function("g", intSort, intSort)

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)

	conjecture1 := z3.Implies(
		z3.And(
			z3.NotEquals(g.ApplyFct(z3.Subtract(g.ApplyFct(x), g.ApplyFct(y))), g.ApplyFct(z)),
			z3.And(z3.LessEq(z3.Add(x, z), y), z3.LessEq(y, x))),
		z3.Less(z, 0))

	fmt.Printf("conjecture1\n%v\n", conjecture1)
	s.Add(z3.Not(conjecture1))
	if s.Check() == z3.Unsat {
		fmt.Println("proved")
	} else {
		fmt.Println("failed to prove")
	}
	s.Reset()

	conjecture2 := z3.Implies(
		z3.And(
			z3.NotEquals(g.ApplyFct(z3.Subtract(g.ApplyFct(x), g.ApplyFct(y))), g.ApplyFct(z)),
			z3.And(z3.LessEq(z3.Add(x, z), y), z3.LessEq(y, x))),
		z3.Less(z, -1))

	fmt.Printf("conjecture1\n%v\n", conjecture1)
	s.Add(z3.Not(conjecture2))
	if s.Check() == z3.Unsat {
		fmt.Println("proved")
	} else {
		fmt.Println("failed to prove")
		fmt.Printf("counterexample: %v\n", s.Get_model())
	}
}

func nonlinearExample1() {
	fmt.Println("nonlinear example 1")
	cfg := z3.NewConfig()
	defer z3.DeleteConfig(cfg)
	cfg.Set("auto_config", true)
	c := z3.NewContext(cfg)
	defer z3.DeleteContext(c)

	x := c.Real_const("x")
	y := c.Real_const("y")
	z := c.Real_const("z")

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)

	// x^2 + y^2 == 1
	s.Add(z3.Equals(z3.Add(z3.Mult(x, x), z3.Mult(y, y)), 1))
	// x^3 + z^3 < 1/2
	s.Add(z3.Less(z3.Add(z3.Mult(x, z3.Mult(x, x)), z3.Mult(z, z3.Mult(z, z))), c.Real_val("1/2")))
	s.Add(z3.NotEquals(z, 0))

	fmt.Println(s.Check())
	m := s.Get_model()
	fmt.Println(m)
	z3.Set_param("pp.decimal", true)
	fmt.Println("model in decimal notation")
	fmt.Println(m)
	z3.Set_param("pp.decimal-precision", 50)
	fmt.Println("model using 50 decimal places")
	fmt.Println(m)
}

func prove(conjecture z3.Expr) {
	c := conjecture.Ctx() // Get the context
	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)
	s.Add(z3.Not(conjecture))
	fmt.Printf("Conjecture: %v\n", conjecture)
	if s.Check() == z3.Unsat {
		fmt.Println("proved")
		return
	}
	fmt.Println("failed to prove")
	fmt.Printf("counterexample:\n%v\n", s.Get_model())
}
func proveWithCast(i interface{}) {
	if expr, ok := i.(z3.Expr); ok {
		prove(expr)
		return
	}
	fmt.Println("error: can only try to prove expressions")
}

/* bitvector_example1
 * Simple bit-vector example. This example disproves that x - 10 <= 0 IFF x <= 10 for (32-bit) machine integers
 */
func bitvectorExample1() {
	fmt.Println("bitvector example 1")
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	x := c.Bv_const("x", 32)

	// using signed <=
	proveWithCast(z3.Equals(z3.LessEq(z3.Subtract(x, 10), 0), z3.LessEq(x, 10)))

	// using unsigned <=
	proveWithCast(z3.Equals(z3.Ule(z3.Subtract(x, 10), 0), z3.Ule(x, 10)))

	y := c.Bv_const("y", 32)
	prove(z3.Implies(z3.Equals(z3.Concat(x, y), z3.Concat(y, x)), z3.Equals(x, y)))
}

/* bitvector_examples2
 * Find x and y such that: x ^ y - 103 == x * y
 */
func bitvectorExample2() {
	fmt.Println("bitvector example 2")
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	x := c.Bv_const("x", 32)
	y := c.Bv_const("y", 32)

	s := z3.NewSolver(c)
	defer z3.DeleteSolver(s)
	conj := z3.Equals(z3.Subtract(z3.BXor(x, y), 103), z3.Mult(x, y))
	s.Add(conj)
	fmt.Println(s)
	fmt.Println(s.Check())
	fmt.Println(s.Get_model())
}

// capi_example skipped

func errorExample() {
	fmt.Println("error example")
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	x := c.Bool_const("x")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v\n", r)
		}
	}()

	// The next call fails because x is a Boolean.
	expr := z3.Add(x, 1)
	fmt.Println(expr)

	// skiped other parts of the example since we do not have C api available
}

// skipped ite_example1 since it is c api
func iteExample2() {
	fmt.Println("if-then-else example2")
	c := z3.NewContext()
	defer z3.DeleteContext(c)
	b := c.Bool_const("b")
	x := c.Int_const("x")
	y := c.Int_const("y")
	fmt.Println(z3.Greater(z3.Ite(b, x, y), 0))
}

func unsatCoreExample1() {
	fmt.Println("unsat core example1")
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
	fmt.Println(s.Check(assumptions1))
	core := s.Unsat_core()
	fmt.Println(core)
	fmt.Printf("size: %v\n", core.Size())
	for i := 0; uint(i) < core.Size(); i++ {
		core.Get(i)
	}
	assumptions2 := z3.NewExprVector(c)
	defer z3.DeleteExprVector(assumptions2)
	assumptions2.Push_back(p1)
	assumptions2.Push_back(p3)
	fmt.Println(s.Check(assumptions2))
}

func main() {
	deMorgan()
	findModelExample1()
	proveExample1()
	proveExample2()
	nonlinearExample1()
	bitvectorExample1()
	bitvectorExample2()
	errorExample()
	iteExample2()
	unsatCoreExample1()
}
