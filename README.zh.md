[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/rese/release.yml?branch=main&label=BUILD)](https://github.com/yylego/rese/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/rese)](https://pkg.go.dev/github.com/yylego/rese)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/rese/master.svg)](https://coveralls.io/github/yylego/rese?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/rese.svg)](https://github.com/yylego/rese/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/rese)](https://goreportcard.com/report/github.com/yylego/rese)

# rese

**rese** = **res**(结果) + **e**(错误) - 用优雅的结果提取简化 Go 错误处理。

将冗长的错误检查转换为简洁、富有表现力的代码。

## 为什么选择 rese？

**传统 Go 错误处理：**
```go
num, err := getNumber()
if err != nil {
    panic(err)
}

text, err := getText()
if err != nil {
    panic(err)
}

result := fmt.Sprintf("%s: %d", text, num)
```

**使用 rese：**
```go
result := fmt.Sprintf("%s: %d", rese.V1(getText()), rese.V1(getNumber()))
```

简洁、优雅、安全！

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## English README

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 安装

```bash
go get github.com/yylego/rese
```

---

## 快速开始

### 基础值处理

```go
// V 系列 - 检查错误并返回值
result := rese.V1(getValue())  // 返回值或在错误时 panic
a, b := rese.V2(getTwoValues()) // 返回两个值或 panic
```

### 指针验证

```go
// P 系列 - 检查错误并确保指针非 nil
ptr := rese.P1(getPointer())  // 返回非 nil 指针或 panic
```

### 非零值验证

```go
// C 系列 - 检查错误并确保值非零
count := rese.C1(getCount())  // 返回非零值或 panic
```

---

## 函数系列概览

rese 提供多个函数系列（每个系列支持 0-9 个参数）满足不同的验证需求：

| 系列 | 用途 | 使用示例 |
|------|------|----------|
| **V** (Value) | 检查错误，返回值 | `rese.V1(getValue())` |
| **P** (Pointer) | 检查错误 + 指针非 nil | `rese.P1(getPointer())` |
| **C** (Comparable) | 检查错误 + 值非零 | `rese.C1(getCount())` |
| **A** (Array/Slice) | 检查错误 + 切片非空 | `rese.A1(getItems())` |
| **M** (Map) | 检查错误 + 映射非空 | `rese.M1(getMapping())` |
| **F** (Function) | 执行函数 + V 系列验证 | `rese.F1(getValue)` |
| **R** (Run pointer) | 执行函数 + P 系列验证 | `rese.R1(getPointer)` |

### resb 子包（布尔结果）

处理 `(value, bool)` 返回模式：

| 系列 | 用途 | 使用示例 |
|------|------|----------|
| **resb.V** | 检查布尔值，返回值 | `resb.V1(getValue())` |
| **resb.P** | 检查布尔值 + 指针非 nil | `resb.P1(getPointer())` |
| **resb.C** | 检查布尔值 + 值非零 | `resb.C1(getCount())` |

---

## 使用方法

### 示例 1: V 系列（值）

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// getValue 模拟一个可能返回错误的函数
// 演示使用 V1 函数进行错误处理
func getValue() (string, error) {
	return "success-result", nil
}

// getNumber 模拟一个可能返回错误的数字函数
// 用于演示 V1 处理不同数据类型
func getNumber() (int, error) {
	return 42, nil
}

// getTwoValues 演示返回两个值的函数可能返回错误
// 展示 V2 用于提取多个值
func getTwoValues() (string, int, error) {
	return "data", 100, nil
}

func main() {
	// 示例 1: 基础 V1 用法 - 没有错误时提取值
	// 当有错误时，V1 会 panic
	result := rese.V1(getValue())
	fmt.Println("Result:", result)

	// 示例 2: V1 处理数字类型
	// 演示 V1 可以处理不同数据类型
	number := rese.V1(getNumber())
	fmt.Println("Number:", number)

	// 示例 3: V2 用法 - 没有错误时提取两个值
	// 展示处理多个返回值
	text, count := rese.V2(getTwoValues())
	fmt.Println("Text:", text, "Count:", count)

	// 示例 4: 简洁的内联用法
	// 演示不使用中间变量的简洁代码
	message := fmt.Sprintf("Got %s with count %d", rese.V1(getValue()), rese.V1(getNumber()))
	fmt.Println(message)
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 示例 2: P 系列（指针）

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// Account 表示简单的账号数据结构
// 用于演示 P1 进行指针验证
type Account struct {
	Name string
	Age  int
}

// getAccount 模拟获取账号数据
// 返回指向 Account 的指针及可能的错误
func getAccount() (*Account, error) {
	return &Account{Name: "Alice", Age: 30}, nil
}

// Config 表示应用程序配置
// 用于演示 P 系列处理不同类型
type Config struct {
	Host string
	Port int
}

// getConfig 模拟获取配置
// 返回指向 Config 的指针及可能的错误
func getConfig() (*Config, error) {
	return &Config{Host: "localhost", Port: 8080}, nil
}

// getTwoPointers 演示返回两个指针的函数可能返回错误
// 展示 P2 用于验证多个指针
func getTwoPointers() (*Account, *Config, error) {
	return &Account{Name: "Bob", Age: 25}, &Config{Host: "0.0.0.0", Port: 9090}, nil
}

func main() {
	// 示例 1: 基础 P1 用法 - 验证指针非 nil
	// 当有错误时 panic，当指针为 nil 时 panic
	account := rese.P1(getAccount())
	fmt.Printf("Account: %s, Age: %d\n", account.Name, account.Age)

	// 示例 2: P1 处理不同类型
	// 演示 P1 可以处理各种指针类型
	config := rese.P1(getConfig())
	fmt.Printf("Config: %s:%d\n", config.Host, config.Port)

	// 示例 3: P2 用法 - 验证两个指针
	// 展示处理多个指针返回值
	acc, cfg := rese.P2(getTwoPointers())
	fmt.Printf("Account: %s, Config: %s:%d\n", acc.Name, cfg.Host, cfg.Port)

	// 示例 4: 简洁的内联用法访问指针
	// 演示不使用中间变量的简洁代码
	message := fmt.Sprintf("Account %s is %d years old", rese.P1(getAccount()).Name, rese.P1(getAccount()).Age)
	fmt.Println(message)
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 示例 3: 多系列功能（resb, A, M）

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
	"github.com/yylego/rese/resb"
)

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

// getAccountPointer 模拟使用布尔标志获取账号数据
// 返回指针和布尔标志而不是错误
func getAccountPointer() (*struct{ Name string }, bool) {
	return &struct{ Name string }{Name: "Bob"}, true
}

// getItems 演示返回切片及错误
// 展示 A1 用于切片验证
func getItems() ([]string, error) {
	return []string{"apple", "banana", "orange"}, nil
}

// getScores 演示返回映射及错误
// 展示 M1 用于映射验证
func getScores() (map[string]int, error) {
	return map[string]int{"math": 95, "english": 88}, nil
}

func main() {
	// 示例 1: resb.V1 用法 - 检查布尔值并返回值
	// 当布尔值为 false 时 panic
	name := resb.V1(lookupValue("name"))
	fmt.Println("Name:", name)

	// 示例 2: resb.V1 处理数字类型
	// 演示 resb 可以处理不同数据类型
	age := resb.V1(lookupNumber("age"))
	fmt.Println("Age:", age)

	// 示例 3: resb.P1 用法 - 验证指针和布尔值
	// 结合指针检查和布尔验证
	account := resb.P1(getAccountPointer())
	fmt.Println("Account:", account.Name)

	// 示例 4: A1 用法 - 验证切片有元素
	// 当有错误时 panic，当切片为 nil/无元素时 panic
	items := rese.A1(getItems())
	fmt.Println("Items:", items)

	// 示例 5: M1 用法 - 验证映射有条目
	// 当有错误时 panic，当映射为 nil/无条目时 panic
	scores := rese.M1(getScores())
	fmt.Println("Scores:", scores)

	// 示例 6: 使用 resb 的简洁内联用法
	// 演示不使用中间变量的简洁映射查找
	message := fmt.Sprintf("%s from %s works as %s",
		resb.V1(lookupValue("name")),
		resb.V1(lookupValue("city")),
		resb.V1(lookupValue("role")))
	fmt.Println(message)
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

---

## 使用示例

### C 系列 - 可比较值验证

**检查非零整数：**
```go
userID := rese.C1(getUserID())  // 返回 12345，如果错误或零值则 panic
fmt.Println("User ID:", userID)
```

**验证多个可比较值：**
```go
id, status := rese.C2(getIDAndStatus())  // 两者都必须非零
```

### A 系列 - 数组/切片验证

**确保切片有元素：**
```go
items := rese.A1(getItems())  // 返回 []string{"apple", "banana", "orange"}
fmt.Println("Items:", items)
```

### M 系列 - 映射验证

**确保映射有条目：**
```go
scores := rese.M1(getScores())  // 返回 map[string]int{"math": 95, "english": 88}
fmt.Println("Scores:", scores)
```

### F 系列 - 函数执行

**执行函数并验证：**
```go
data := rese.F1(fetchData)  // 执行 fetchData()，返回 "data loaded"
fmt.Println(data)
```

**执行并返回多个值：**
```go
user, config := rese.F2(initialize)  // 执行并验证两个值
```

### 布尔结果处理（resb）

**处理映射查找结果：**
```go
value := resb.V1(lookupValue())  // 返回 "found"，如果未找到则 panic
fmt.Println("Value:", value)
```

**检查布尔标志：**
```go
config := resb.V1(m["key"])  // 键不存在时 panic
```

---

## 完整函数参考

### V 系列 - 值验证

| 函数 | 说明 |
|------|------|
| `V1[T1](v1 T1, err) T1` | 检查错误，返回 1 个值 |
| `V2[T1,T2](v1 T1, v2 T2, err) (T1,T2)` | 检查错误，返回 2 个值 |
| `V3` ~ `V9` | 检查错误，返回 3-9 个值 |

> **说明：** 对于无返回值的 `V0(err error)`，建议使用 [must](https://github.com/yylego/must) 包。

### P 系列 - 指针验证

| 函数 | 说明 |
|------|------|
| `P1[T1](v1 *T1, err) *T1` | 检查错误 + 指针非 nil，返回 1 个指针 |
| `P2[T1,T2](v1 *T1, v2 *T2, err) (*T1,*T2)` | 检查错误 + 指针非 nil，返回 2 个指针 |
| `P3` ~ `P9` | 检查错误 + 指针非 nil，返回 3-9 个指针 |

> **说明：** `P0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

### C 系列 - 可比较值验证

| 函数 | 说明 |
|------|------|
| `C1[T1 comparable](v1 T1, err) T1` | 检查错误 + 值非零，返回 1 个值 |
| `C2[T1,T2 comparable](v1 T1, v2 T2, err) (T1,T2)` | 检查错误 + 值非零，返回 2 个值 |
| `C3` ~ `C9` | 检查错误 + 值非零，返回 3-9 个值 |

> **说明：** `C0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

### A 系列 - 数组/切片验证

| 函数 | 说明 |
|------|------|
| `A1[T1](v1 []T1, err) []T1` | 检查错误 + 切片非空，返回 1 个切片 |
| `A2[T1,T2](v1 []T1, v2 []T2, err) ([]T1,[]T2)` | 检查错误 + 切片非空，返回 2 个切片 |
| `A3` ~ `A9` | 检查错误 + 切片非空，返回 3-9 个切片 |

> **说明：** `A0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

### M 系列 - 映射验证

| 函数 | 说明 |
|------|------|
| `M1[K1,V1](m1 map[K1]V1, err) map[K1]V1` | 检查错误 + 映射非空，返回 1 个映射 |
| `M2[K1,K2,V1,V2](m1 map[K1]V1, m2 map[K2]V2, err) (map[K1]V1, map[K2]V2)` | 检查错误 + 映射非空，返回 2 个映射 |
| `M3` ~ `M9` | 检查错误 + 映射非空，返回 3-9 个映射 |

> **说明：** `M0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

### F 系列 - 函数执行（值）

| 函数 | 说明 |
|------|------|
| `F1[T1](run func() (T1, error)) T1` | 执行函数，应用 V1 验证 |
| `F2[T1,T2](run func() (T1, T2, error)) (T1,T2)` | 执行函数，应用 V2 验证 |
| `F3` ~ `F9` | 执行函数，应用 V3-V9 验证 |

> **说明：** `F0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

### R 系列 - 函数执行（指针）

| 函数 | 说明 |
|------|------|
| `R1[T1](run func() (*T1, error)) *T1` | 执行函数，应用 P1 验证 |
| `R2[T1,T2](run func() (*T1, *T2, error)) (*T1,*T2)` | 执行函数，应用 P2 验证 |
| `R3` ~ `R9` | 执行函数，应用 P3-P9 验证 |

> **说明：** `R0` 与 `V0` 相同，建议使用 [must](https://github.com/yylego/must) 包进行简单错误检查。

---

## resb 包 - 布尔结果处理

`resb` 子包处理返回 `(value, bool)` 而非 `(value, error)` 的函数。

### 常见使用场景

```go
// 映射查找 - 返回 (value, bool)
func lookupConfig(key string) (string, bool) {
    cfg := map[string]string{"host": "localhost", "port": "8080"}
    return cfg[key]
}


// 使用 resb.V1 处理这些函数
config := resb.V1(lookupConfig("host"))      // 键不存在时 panic
```

---

## 相关项目

探索此生态系统中的更多错误处理包：

### 高级包

- **[must](https://github.com/yylego/must)** - Must 风格断言，提供丰富的类型支持和详细的错误上下文
- **[rese](https://github.com/yylego/rese)** - 带 panic 的结果提取，专注于安全的值解包（本项目）

### 基础包

- **[done](https://github.com/yylego/done)** - 简单专注的错误处理，支持方法链式调用
- **[sure](https://github.com/yylego/sure)** - 生成代码以创建自定义验证方法

每个包针对不同的使用场景，从快速原型开发到具有全面错误处理的生产系统。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/rese.git`）
3. **导航**：进入克隆的项目（`cd rese`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yylego/rese.svg?variant=adaptive)](https://starchart.cc/yylego/rese)
