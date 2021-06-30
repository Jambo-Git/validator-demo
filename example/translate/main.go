package main

import (
	"fmt"
	"validateDemo/check"
	"validateDemo/mode"
)

var checker *check.Check

func init() {
	checker = check.NewCheck()
}

func main() {
	var test mode.Struct7
	test.ULen = "张"
	err := checker.Struct(test)
	fmt.Println(err)

}
