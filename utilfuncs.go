package z3go

// This file contains utility functions

import (
	"log"
)

// String results a string representation of a value of the
// enum check_result.
func (c Z3Check_result) String() string {
	switch c {
	case Unsat:
		return "unsat"
	case Sat:
		return "sat"
	case Unknown:
		return "unknown"
	default:
		log.Printf("z3go.CheckResultString got called with an invalid check result")
		return "Invalid check result"
	}
}
