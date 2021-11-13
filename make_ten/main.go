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

	opeNums := []int{1, 2, 3, 4}
	for _, v := range Permute(opeNums) {
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

func Permute(nums []int) [][]string {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)

	strs := make([][]string, len(ret), n)
	for i, vals := range ret {
		for _, v := range vals {
			strs[i] = append(strs[i], makeOpe(v))
		}
	}
	return strs
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

func makeOpe(i int) string {
	switch i {
	case 1:
		return plus
	case 2:
		return minus
	case 3:
		return multiply
	case 4:
		return division
	}
	return "" //ココには来ない
}
