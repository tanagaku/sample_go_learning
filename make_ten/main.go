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
	multiply = "*"
	division = "/"
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

func validation(args []string) ([]string, error) {
	if len(args) != 4 {
		return nil, errors.New("引数は4つ指定してください")
	}

	for _, arg := range args {
		if i, err := strconv.ParseInt(arg, 10, 64); err != nil {
			return nil, errors.New("引数は1桁の数値を指定してください " + err.Error())
		} else if i > 9 {
			return nil, errors.New("引数は1桁の数値を指定してください")
		}
	}
	return args, nil
}

func make_ten(nums []string) (string, error) {

	nums, err := convertFloatArray(nums)
	if err != nil {
		return "", err
	}
	opes := []string{plus, minus, multiply, division}
	for _, v := range PermuteWithRepetition(opes, 3) {
		ope := nums[0] + v[0] + nums[1] + v[1] + nums[2] + v[2] + nums[3]

		tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, ope)
		if err != nil {
			return "", errors.New(err.Error())
		}
		if tv.Value.String() == "10" {
			return ope + "=10", nil
		}
	}
	return "", nil
}

//重複順列を返す
func PermuteWithRepetition(opes []string, getn int) [][]string {
	if getn == 0 {
		return [][]string{{}}
	}

	results := [][]string{}
	for i := 0; i < len(opes); i++ {
		childPatterns := PermuteWithRepetition(opes, getn-1)
		for _, chilchildPattern := range childPatterns {
			pattern := append([]string{opes[i]}, chilchildPattern...)
			results = append(results, pattern)
		}
	}

	return results
}

func convertFloatArray(nums []string) ([]string, error) {
	var ss []string
	for _, n := range nums {
		s, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return nil, err
		}

		ss = append(ss, strconv.FormatFloat(s, 'f', -1, 64))
	}
	return ss, nil
}
