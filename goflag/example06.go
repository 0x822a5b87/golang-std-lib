package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
)

type MathCommand struct {
	Op     string `long:"op" description:"operation to execute" subcommands-optional:"+"`
	Args   []string
	Result int64
}

func (mc *MathCommand) Execute(args []string) error {
	fmt.Println("mc.Op: ", mc.Op)
	for _, arg := range args {
		num, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}

		mc.Result += num
	}

	mc.Args = args
	return nil
}

func example06Main() {
	type Option struct {
		Math MathCommand `command:"math"`
	}

	var opt Option
	_, err := flags.Parse(&opt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result of %s is %d", strings.Join(opt.Math.Args, opt.Math.Op), opt.Math.Result)
}
