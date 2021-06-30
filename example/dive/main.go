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
	var test mode.Struct2
	// 不能为空
	err := checker.Struct(test)
	fmt.Println(err)

	// 长度大于0
	test.Array = []string{}
	err = checker.Struct(test)
	fmt.Println(err)
	// 元素不能为空
	test.Array = []string{"11", ""}
	err = checker.Struct(test)
	fmt.Println(err)

	var test1 mode.Struct3
	checker.RegisterAlias("keymax", "max=5")

	// 不能为空
	err = checker.Struct(test1)
	fmt.Println(err)
	// 长度大于0
	test1.Map = map[string]string{}
	err = checker.Struct(test1)
	fmt.Println(err)

	// key值长度最大为5，value值不能为空
	test1.Map = map[string]string{"123456": "1", "12345": ""}
	err = checker.Struct(test1)
	fmt.Println(err)

	test1.Map = map[string]string{"123": "1", "1234": "2", "12345": "333"}
	err = checker.Struct(test1)
	fmt.Println(err)

	var test2 mode.Struct4
	err = checker.Struct(test2)
	fmt.Println(err)

	test2.Column1 = []*mode.Struct5{{Column2: ""}, nil}
	err = checker.Struct(test2)
	fmt.Println(err)
}
