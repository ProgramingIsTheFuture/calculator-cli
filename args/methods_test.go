package args_test

import (
	"os"
	"strings"
	"testing"

	"github.com/ProgramingIsTheFuture/Cli/args"
	"github.com/alexflint/go-arg"
	"github.com/stretchr/testify/require"
)

func CreateArgs() args.Args {
	var args args.Args
	arg.MustParse(&args)

	return args
}

func TestSumIntNums(t *testing.T) {
	os.Args = strings.Split("./main intsum --num1 2 --num2 2", " ")

	args := CreateArgs()

	result := args.SumIntNums()

	require.Equal(t, int32(4), result)
}

func TestSumFloatNums(t *testing.T) {
	os.Args = strings.Split("./main floatsum --num1 2.5 --num2 2.5", " ")

	args := CreateArgs()

	result := args.SumFloatNums()

	require.Equal(t, float32(5), result)
}
