package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// getValue simulates a function that returns a value with potential error
// Demonstrates error handling with V1 function
//
// getValue 模拟一个可能返回错误的函数
// 演示使用 V1 函数进行错误处理
func getValue() (string, error) {
	return "success-result", nil
}

// getNumber simulates a function that returns a numeric value with potential error
// Used to demonstrate V1 with different data types
//
// getNumber 模拟一个可能返回错误的数字函数
// 用于演示 V1 处理不同数据类型
func getNumber() (int, error) {
	return 42, nil
}

// getTwoValues demonstrates function returning two values with potential error
// Shows V2 usage for multiple value extraction
//
// getTwoValues 演示返回两个值的函数可能返回错误
// 展示 V2 用于提取多个值
func getTwoValues() (string, int, error) {
	return "data", 100, nil
}

func main() {
	// Example 1: Basic V1 usage - extracts value when no error
	// When error is present, V1 panics
	//
	// 示例 1: 基础 V1 用法 - 没有错误时提取值
	// 当有错误时，V1 会 panic
	result := rese.V1(getValue())
	fmt.Println("Result:", result)

	// Example 2: V1 with numeric type
	// Demonstrates V1 works with different data types
	//
	// 示例 2: V1 处理数字类型
	// 演示 V1 可以处理不同数据类型
	number := rese.V1(getNumber())
	fmt.Println("Number:", number)

	// Example 3: V2 usage - extracts two values when no error
	// Shows handling of multiple return values
	//
	// 示例 3: V2 用法 - 没有错误时提取两个值
	// 展示处理多个返回值
	text, count := rese.V2(getTwoValues())
	fmt.Println("Text:", text, "Count:", count)

	// Example 4: Compact inline usage
	// Demonstrates clean code without intermediate variables
	//
	// 示例 4: 简洁的内联用法
	// 演示不使用中间变量的简洁代码
	message := fmt.Sprintf("Got %s with count %d", rese.V1(getValue()), rese.V1(getNumber()))
	fmt.Println(message)
}
