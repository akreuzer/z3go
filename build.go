//Package z3go contains a wrapper for the Z3 SMT-solver library.
package z3go

//go:generate swig -go -cgo -c++ -intgosize 64 -module z3go z3++.h

// #cgo CPPFLAGS: -I/usr/local/Cellar/z3/4.5.0/include
// #cgo LDFLAGS: -L/usr/local/Cellar/z3/4.5.0/lib -lz3
import "C"
