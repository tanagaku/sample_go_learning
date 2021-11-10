package main

import (
	"errors"
	"flag"
	"fmt"
	"go/token"
	"go/types"
	"strconv"
)

const (
	plus     = "+"
	minus    = "-"
	multiply = "×"
	division = "÷"
)

/*条件:コマンドライン引数は4つであること
コマンドライン引数はすべて1桁の数値であること*/
func main() {
	flag.Parse()
	args := flag.Args()

	nums, err := validation(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := make_ten(nums)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func validation(args []string) ([]int64, error) {
	if len(args) != 4 {
		return nil, errors.New("引数は4つ指定してください")
	}
	var nums []int64

	for _, arg := range args {
		if i, err := strconv.ParseInt(arg, 10, 64); err != nil {
			return nil, errors.New("引数は1桁の数値を指定してください " + err.Error())
		} else if i > 9 {
			return nil, errors.New("引数は1桁の数値を指定してください")
		} else {
			nums = append(nums, i)
		}
	}
	return nums, nil
}

func make_ten(nums []int64) (string, error) {

	test := "1+2+3+4"
	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, test)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if tv.Value.String() == "10" {
		return test, nil
	}
	return "", nil
}
