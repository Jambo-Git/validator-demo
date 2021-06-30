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
	var test mode.Struct1
	// 必填
	test.RequiredString = ""
	err := checker.Struct(test)
	fmt.Println(fmt.Sprintf("RequiredString########################################：\n%v", err))
	test.RequiredString = "1"
	// 必填,数字类型不能为0
	test.RequiredNumber = 0
	err = checker.Struct(test)
	fmt.Println(fmt.Sprintf("RequiredNumber########################################：\n%v", err))
	test.RequiredNumber = 1

}
