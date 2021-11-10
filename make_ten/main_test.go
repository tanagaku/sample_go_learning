package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validation(t *testing.T) {
	//引数が4つではない
	args := []string{"1", "2", "3"}
	result, err := validation(args)
	assert.Nil(t, result)
	assert.Equal(t, "引数は4つ指定してください", err.Error())

	//数値以外を指定
	args = []string{"1", "2", "3", "TEST"}
	result, err = validation(args)
	assert.Nil(t, result)
	assert.Equal(t, "引数は1桁の数値を指定してください strconv.ParseInt: parsing \"TEST\": invalid syntax", err.Error())

	//2桁以上の数値を指定
	args = []string{"1", "222", "3", "4"}
	result, err = validation(args)
	assert.Nil(t, result)
	assert.Equal(t, "引数は1桁の数値を指定してください", err.Error())

	//正常
	args = []string{"1", "2", "3", "4"}
	result, err = validation(args)
	assert.Nil(t, err)
	assert.Equal(t, []int64{1, 2, 3, 4}, result)

}
