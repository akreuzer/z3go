# z3go: A Z3 wrapper for Golang

[![Build Status](https://travis-ci.org/akreuzer/z3go.svg?branch=master)](https://travis-ci.org/akreuzer/z3go)

z3go is a wrapper library for the [Z3 SMT-solver](https://github.com/Z3Prover/z3).
It uses [SWIG](http://www.swig.org/).

## Installation

Make sure that you have Z3 installed. (On mac `brew install z3`)

```bash
# Set the include and library path if needed
# For macOS and Z3 installed using homebrew
export CGO_CPPFLAGS="-I/usr/local/Cellar/z3/4.5.0/include"
export CGO_LDFLAGS="-L/usr/local/Cellar/z3/4.5.0/lib"

go install github.com/akreuzer/z3go
```

## Documentation

z3go is a wrapper of the C++ interface of Z3.
We started translation the examples for the Z3 C++ interface to to Go.
The can be found in the `examples/` folder.

We renamed operator that clashed with the Go-Syntax.

| C++ Name | z3go   |
|----------|--------|
| !        | Not    |
| \|\|     | Or     |
| &&       | And    |
| ==       | Equals |
| !=       | NotEquals |
| <        | Less   |
| <=       | LessEq |
| >        | Greater |
| >=       | GreaterEq |
| +        | Add    |
| -        | Subtract |
| *        | Mult   |
| /        | Div    |
| model[i] | model.Get(i) |
| f(x)     | f.ApplyFct(x) |
| ^        | BXor (Bitwise xor)  |
| \|       | BOr    |
| &        | BAnd   |
| ~        | BComp (Bitwise complement) |
| & (Tactics) | TacticAnd |
| \| (Tactics) | TacticOr |
| &&, \|\|, <,... (Probe) | ProbeAnd, ProbeOr, ProbeLess, ... |


## Hacking

Edit `z3++.h` and then use `go generate` to call SWIG.
