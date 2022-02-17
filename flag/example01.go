package main

import (
	"flag"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
	floatFlag  float64
)

func init() {
	flag.IntVar(&intFlag, "intFlag", 0, "int flag value")
	flag.BoolVar(&boolFlag, "boolFlag", true, "bool flag value")
	flag.StringVar(&stringFlag, "stringFlag", "default", "string flag value")
	flag.Float64Var(&floatFlag, "floatFlag", 0.0, "float flag value")
}
