package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const otherWord = "otherWord"
const transformYaml = "transform.yaml"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	b, err := ioutil.ReadFile(transformYaml)
	if err != nil {
		panic(err.Error())
	}

	transforms := make([]string, 0) //項目が増えた場合構造体にする
	err = yaml.Unmarshal(b, &transforms)
	if err != nil {
		panic(err.Error())
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
