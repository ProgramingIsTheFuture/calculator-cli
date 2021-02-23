package main

import (
	"github.com/ProgramingIsTheFuture/Cli/args"
	"github.com/ProgramingIsTheFuture/Cli/loggers"
	"github.com/alexflint/go-arg"
)

func main() {
	var args args.Args
	arg.MustParse(&args)

	switch {
	case args.SumInt != nil:
		loggers.LogResult(args.SumIntNums())
	case args.SubInt != nil:
		loggers.LogResult(args.SubIntNums())

	case args.SumFloat != nil:
		loggers.LogResult(args.SumFloatNums())
	case args.SubFloat != nil:
		loggers.LogResult(args.SubFloatNums())
	}
}
