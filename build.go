//go:generate swig -go -cgo -c++ -intgosize 64 -module z3go z3++.h
package z3go

// #cgo CPPFLAGS: -I/usr/local/Cellar/z3/4.5.0/include
// #cgo LDFLAGS: -L/usr/local/Cellar/z3/4.5.0/lib -lz3
import "C"
