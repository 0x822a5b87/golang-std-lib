package main

import (
	"flag"
)

var (
	pIntFlag    *int
	pBoolFlag   *bool
	pStringFlag *string
)

func init() {
	pIntFlag = flag.Int("pIntFlag", 0, "int flag value")
	pBoolFlag = flag.Bool("pBoolFlag", false, "bool flag value")
	pStringFlag = flag.String("pStringFlag", "default", "string flag value")
}
