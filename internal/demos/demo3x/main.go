package main

import (
	"fmt"

	"github.com/yylego/rese"
	"github.com/yylego/rese/resb"
)

// lookupValue simulates a map lookup operation
// Returns value with bool flag indicating success
//
// lookupValue 模拟映射查找操作
// 返回值和布尔标志表示成功与否
func lookupValue(key string) (string, bool) {
	data := map[string]string{
		"name": "Alice",
		"city": "Shanghai",
		"role": "Engineer",
	}
	value, ok := data[key]
	return value, ok
}

// lookupNumber simulates retrieving a numeric value
// Returns number with bool flag indicating presence
//
// lookupNumber 模拟获取数字值
// 返回数字和布尔标志表示是否存在
func lookupNumber(key string) (int, bool) {
	data := map[string]int{
		"age":   30,
		"score": 95,
		"count": 100,
	}
	value, ok := data[key]
	return value, ok
}

// getAccountPointer simulates fetching account data with bool flag
// Returns pointer with bool flag instead of error
//
// getAccountPointer 模拟使用布尔标志获取账号数据
// 返回指针和布尔标志而不是错误
func getAccountPointer() (*struct{ Name string }, bool) {
	return &struct{ Name string }{Name: "Bob"}, true
}

// getItems demonstrates slice return with error
// Shows A1 usage for slice validation
//
// getItems 演示返回切片及错误
// 展示 A1 用于切片验证
func getItems() ([]string, error) {
	return []string{"apple", "banana", "orange"}, nil
}

// getScores demonstrates map return with error
// Shows M1 usage for map validation
//
// getScores 演示返回映射及错误
// 展示 M1 用于映射验证
func getScores() (map[string]int, error) {
	return map[string]int{"math": 95, "english": 88}, nil
}

func main() {
	// Example 1: resb.V1 usage - checks bool and returns value
	// Panics when bool is false
	//
	// 示例 1: resb.V1 用法 - 检查布尔值并返回值
	// 当布尔值为 false 时 panic
	name := resb.V1(lookupValue("name"))
	fmt.Println("Name:", name)

	// Example 2: resb.V1 with numeric type
	// Demonstrates resb works with different data types
	//
	// 示例 2: resb.V1 处理数字类型
	// 演示 resb 可以处理不同数据类型
	age := resb.V1(lookupNumber("age"))
	fmt.Println("Age:", age)

	// Example 3: resb.P1 usage - validates pointer and bool
	// Combines pointer check with bool validation
	//
	// 示例 3: resb.P1 用法 - 验证指针和布尔值
	// 结合指针检查和布尔验证
	account := resb.P1(getAccountPointer())
	fmt.Println("Account:", account.Name)

	// Example 4: A1 usage - validates slice has elements
	// Panics when error is present and when slice is nil/has no elements
	//
	// 示例 4: A1 用法 - 验证切片有元素
	// 当有错误时 panic，当切片为 nil/无元素时 panic
	items := rese.A1(getItems())
	fmt.Println("Items:", items)

	// Example 5: M1 usage - validates map has entries
	// Panics when error is present and when map is nil/has no entries
	//
	// 示例 5: M1 用法 - 验证映射有条目
	// 当有错误时 panic，当映射为 nil/无条目时 panic
	scores := rese.M1(getScores())
	fmt.Println("Scores:", scores)

	// Example 6: Compact inline usage with resb
	// Demonstrates clean map lookup without intermediate variables
	//
	// 示例 6: 使用 resb 的简洁内联用法
	// 演示不使用中间变量的简洁映射查找
	message := fmt.Sprintf("%s from %s works as %s",
		resb.V1(lookupValue("name")),
		resb.V1(lookupValue("city")),
		resb.V1(lookupValue("role")))
	fmt.Println(message)
}
