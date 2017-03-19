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

func main() {
	deMorgan()
	findModelExample1()
}
