package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// Account represents a simple account data structure
// Used to demonstrate pointer validation with P1
//
// Account 表示简单的账号数据结构
// 用于演示 P1 进行指针验证
type Account struct {
	Name string
	Age  int
}

// getAccount simulates fetching account data
// Returns pointer to Account with potential error
//
// getAccount 模拟获取账号数据
// 返回指向 Account 的指针及可能的错误
func getAccount() (*Account, error) {
	return &Account{Name: "Alice", Age: 30}, nil
}

// Config represents application configuration
// Used to demonstrate P series with different types
//
// Config 表示应用程序配置
// 用于演示 P 系列处理不同类型
type Config struct {
	Host string
	Port int
}

// getConfig simulates fetching configuration
// Returns pointer to Config with potential error
//
// getConfig 模拟获取配置
// 返回指向 Config 的指针及可能的错误
func getConfig() (*Config, error) {
	return &Config{Host: "localhost", Port: 8080}, nil
}

// getTwoPointers demonstrates function returning two pointers with potential error
// Shows P2 usage for multiple pointer validation
//
// getTwoPointers 演示返回两个指针的函数可能返回错误
// 展示 P2 用于验证多个指针
func getTwoPointers() (*Account, *Config, error) {
	return &Account{Name: "Bob", Age: 25}, &Config{Host: "0.0.0.0", Port: 9090}, nil
}

func main() {
	// Example 1: Basic P1 usage - validates pointer is non-nil
	// Panics when error is present and when pointer is nil
	//
	// 示例 1: 基础 P1 用法 - 验证指针非 nil
	// 当有错误时 panic，当指针为 nil 时 panic
	account := rese.P1(getAccount())
	fmt.Printf("Account: %s, Age: %d\n", account.Name, account.Age)

	// Example 2: P1 with different type
	// Demonstrates P1 works with various pointer types
	//
	// 示例 2: P1 处理不同类型
	// 演示 P1 可以处理各种指针类型
	config := rese.P1(getConfig())
	fmt.Printf("Config: %s:%d\n", config.Host, config.Port)

	// Example 3: P2 usage - validates two pointers
	// Shows handling of multiple pointer return values
	//
	// 示例 3: P2 用法 - 验证两个指针
	// 展示处理多个指针返回值
	acc, cfg := rese.P2(getTwoPointers())
	fmt.Printf("Account: %s, Config: %s:%d\n", acc.Name, cfg.Host, cfg.Port)

	// Example 4: Compact inline usage with pointer access
	// Demonstrates clean code without intermediate variables
	//
	// 示例 4: 简洁的内联用法访问指针
	// 演示不使用中间变量的简洁代码
	message := fmt.Sprintf("Account %s is %d years old", rese.P1(getAccount()).Name, rese.P1(getAccount()).Age)
	fmt.Println(message)
}
