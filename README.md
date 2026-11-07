<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/rese/release.yml?branch=main&label=BUILD)](https://github.com/yylego/rese/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/rese)](https://pkg.go.dev/github.com/yylego/rese)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/rese/main.svg)](https://coveralls.io/github/yylego/rese?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/rese.svg)](https://github.com/yylego/rese/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/rese)](https://goreportcard.com/report/github.com/yylego/rese)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# rese

**rese** = **res**(ult) + **e**(rror) - Streamline Go error handling with elegant result extraction.

Transform verbose error checking into clean, expressive code.

## Design Intent

**Traditional Go error handling:**

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

**With rese:**

```go
result := fmt.Sprintf("%s: %d", rese.V1(getText()), rese.V1(getNumber()))
```

Clean, concise, and safe!

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

---

## Installation

```bash
go get github.com/yylego/rese
```

---

## Quick Start

### Basic Value Handling

```go
// V series - Check error and return value
result := rese.V1(getValue())  // Returns value and panics when error is present
a, b := rese.V2(getTwoValues()) // Returns two values and panics when error is present
```

### Pointer Validation

```go
// P series - Check error AND ensure pointer is non-nil
ptr := rese.P1(getPointer())  // Returns non-nil pointer and panics when error/nil is present
```

### Non-Zero Value Validation

```go
// C series - Check error AND ensure value is non-zero
count := rese.C1(getCount())  // Returns non-zero value and panics when error/zero is present
```

---

## Function Series Overview

rese provides multiple function series (0-9 parameters each) for different validation needs:

| Series              | Purpose                          | Example Usage           |
| ------------------- | -------------------------------- | ----------------------- |
| **V** (Value)       | Check error, return value(s)     | `rese.V1(getValue())`   |
| **P** (Pointer)     | Check error + pointer non-nil    | `rese.P1(getPointer())` |
| **C** (Comparable)  | Check error + value non-zero     | `rese.C1(getCount())`   |
| **A** (Array/Slice) | Check error + slice has elements | `rese.A1(getItems())`   |
| **M** (Map)         | Check error + map has entries    | `rese.M1(getMapping())` |
| **F** (Function)    | Execute function + V series      | `rese.F1(getValue)`     |
| **R** (Run pointer) | Execute function + P series      | `rese.R1(getPointer)`   |

### resb Subpackage (Boolean Results)

Handle `(value, bool)` return patterns:

| Series     | Purpose                      | Example Usage           |
| ---------- | ---------------------------- | ----------------------- |
| **resb.V** | Check bool, return value(s)  | `resb.V1(getValue())`   |
| **resb.P** | Check bool + pointer non-nil | `resb.P1(getPointer())` |
| **resb.C** | Check bool + value non-zero  | `resb.C1(getCount())`   |

### Panic Outcome

When a check fails, rese writes a PANIC log that points at the line of business code that calls rese, then panics with an `error`. Given an error, the V, P, C, A, M, F and R series panic with that same error, so `errors.Is` and `errors.As` work on the value `recover()` gets back.

---

## Usage

### Example 1: Value Series (V)

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// getValue simulates a function that returns a value with potential error
// Demonstrates error handling with V1 function
func getValue() (string, error) {
	return "success-result", nil
}

// getNumber simulates a function that returns a numeric value with potential error
// Used to demonstrate V1 with different data types
func getNumber() (int, error) {
	return 42, nil
}

// getTwoValues demonstrates function returning two values with potential error
// Shows V2 usage for multiple value extraction
func getTwoValues() (string, int, error) {
	return "data", 100, nil
}

func main() {
	// Example 1: Basic V1 usage - extracts value when no error
	// When error is present, V1 panics
	result := rese.V1(getValue())
	fmt.Println("Result:", result)

	// Example 2: V1 with numeric type
	// Demonstrates V1 works with different data types
	number := rese.V1(getNumber())
	fmt.Println("Number:", number)

	// Example 3: V2 usage - extracts two values when no error
	// Shows handling of multiple return values
	text, count := rese.V2(getTwoValues())
	fmt.Println("Text:", text, "Count:", count)

	// Example 4: Compact inline usage
	// Demonstrates clean code without intermediate variables
	message := fmt.Sprintf("Got %s with count %d", rese.V1(getValue()), rese.V1(getNumber()))
	fmt.Println(message)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Example 2: Pointer Series (P)

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
)

// Account represents a simple account data structure
// Used to demonstrate pointer validation with P1
type Account struct {
	Name string
	Age  int
}

// getAccount simulates fetching account data
// Returns pointer to Account with potential error
func getAccount() (*Account, error) {
	return &Account{Name: "Alice", Age: 30}, nil
}

// Config represents application configuration
// Used to demonstrate P series with different types
type Config struct {
	Host string
	Port int
}

// getConfig simulates fetching configuration
// Returns pointer to Config with potential error
func getConfig() (*Config, error) {
	return &Config{Host: "localhost", Port: 8080}, nil
}

// getTwoPointers demonstrates function returning two pointers with potential error
// Shows P2 usage for multiple pointer validation
func getTwoPointers() (*Account, *Config, error) {
	return &Account{Name: "Bob", Age: 25}, &Config{Host: "0.0.0.0", Port: 9090}, nil
}

func main() {
	// Example 1: Basic P1 usage - validates pointer is non-nil
	// Panics when error is present and when pointer is nil
	account := rese.P1(getAccount())
	fmt.Printf("Account: %s, Age: %d\n", account.Name, account.Age)

	// Example 2: P1 with different type
	// Demonstrates P1 works with various pointer types
	config := rese.P1(getConfig())
	fmt.Printf("Config: %s:%d\n", config.Host, config.Port)

	// Example 3: P2 usage - validates two pointers
	// Shows handling of multiple pointer return values
	acc, cfg := rese.P2(getTwoPointers())
	fmt.Printf("Account: %s, Config: %s:%d\n", acc.Name, cfg.Host, cfg.Port)

	// Example 4: Compact inline usage with pointer access
	// Demonstrates clean code without intermediate variables
	message := fmt.Sprintf("Account %s is %d years old", rese.P1(getAccount()).Name, rese.P1(getAccount()).Age)
	fmt.Println(message)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Example 3: Multiple Series (resb, A, M)

```go
package main

import (
	"fmt"

	"github.com/yylego/rese"
	"github.com/yylego/rese/resb"
)

// lookupValue simulates a map lookup operation
// Returns value with bool flag indicating success
func lookupValue(key string) (string, bool) {
	data := map[string]string{
		"name": "Alice",
		"team": "Platform",
		"role": "Engineer",
	}
	value, ok := data[key]
	return value, ok
}

// lookupNumber simulates retrieving a numeric value
// Returns number with bool flag indicating presence
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
func getAccountPointer() (*struct{ Name string }, bool) {
	return &struct{ Name string }{Name: "Bob"}, true
}

// getItems demonstrates slice return with error
// Shows A1 usage for slice validation
func getItems() ([]string, error) {
	return []string{"apple", "banana", "orange"}, nil
}

// getScores demonstrates map return with error
// Shows M1 usage for map validation
func getScores() (map[string]int, error) {
	return map[string]int{"math": 95, "english": 88}, nil
}

func main() {
	// Example 1: resb.V1 usage - checks bool and returns value
	// Panics when bool is false
	name := resb.V1(lookupValue("name"))
	fmt.Println("Name:", name)

	// Example 2: resb.V1 with numeric type
	// Demonstrates resb works with different data types
	age := resb.V1(lookupNumber("age"))
	fmt.Println("Age:", age)

	// Example 3: resb.P1 usage - validates pointer and bool
	// Combines pointer check with bool validation
	account := resb.P1(getAccountPointer())
	fmt.Println("Account:", account.Name)

	// Example 4: A1 usage - validates slice has elements
	// Panics when error is present and when slice is nil/has no elements
	items := rese.A1(getItems())
	fmt.Println("Items:", items)

	// Example 5: M1 usage - validates map has entries
	// Panics when error is present and when map is nil/has no entries
	scores := rese.M1(getScores())
	fmt.Println("Scores:", scores)

	// Example 6: Compact inline usage with resb
	// Demonstrates clean map lookup without intermediate variables
	message := fmt.Sprintf("%s in team %s works as %s",
		resb.V1(lookupValue("name")),
		resb.V1(lookupValue("team")),
		resb.V1(lookupValue("role")))
	fmt.Println(message)
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

---

## Examples

### C Series - Comparable Value Validation

**Check non-zero integer:**

```go
userID := rese.C1(getUserID())  // Returns 12345, panics if error or zero
fmt.Println("User ID:", userID)
```

**Validate multiple comparable values:**

```go
id, status := rese.C2(getIDAndStatus())  // Both must be non-zero
```

### A Series - Array/Slice Validation

**Ensure slice has elements:**

```go
items := rese.A1(getItems())  // Returns []string{"apple", "banana", "orange"}
fmt.Println("Items:", items)
```

### M Series - Map Validation

**Ensure map has entries:**

```go
scores := rese.M1(getScores())  // Returns map[string]int{"math": 95, "english": 88}
fmt.Println("Scores:", scores)
```

### F Series - Function Execution

**Execute function and validate:**

```go
data := rese.F1(fetchData)  // Executes fetchData(), returns "data loaded"
fmt.Println(data)
```

**Execute with multiple returns:**

```go
user, config := rese.F2(initialize)  // Executes and validates two values
```

### Boolean Result Handling (resb)

**Handle map lookup results:**

```go
value := resb.V1(lookupValue())  // Returns "found", panics if not found
fmt.Println("Value:", value)
```

**Read an environment variable that must be set:**

```go
home := resb.V1(os.LookupEnv("HOME"))  // Panics if HOME is not set
```

---

## Complete Function Reference

### V Series - Value Validation

| Function                               | Description                    |
| -------------------------------------- | ------------------------------ |
| `V1[T1](v1 T1, err) T1`                | Check error, return 1 value    |
| `V2[T1,T2](v1 T1, v2 T2, err) (T1,T2)` | Check error, return 2 values   |
| `V3` ~ `V9`                            | Check error, return 3-9 values |

> **Note:** For `V0(err error)` with no return value, consider using [must](https://github.com/yylego/must) package instead.

### P Series - Pointer Validation

| Function                                   | Description                                         |
| ------------------------------------------ | --------------------------------------------------- |
| `P1[T1](v1 *T1, err) *T1`                  | Check error + pointer non-nil, return 1 pointer     |
| `P2[T1,T2](v1 *T1, v2 *T2, err) (*T1,*T2)` | Check error + pointers non-nil, return 2 pointers   |
| `P3` ~ `P9`                                | Check error + pointers non-nil, return 3-9 pointers |

> **Note:** `P0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

### C Series - Comparable Validation

| Function                                          | Description                                      |
| ------------------------------------------------- | ------------------------------------------------ |
| `C1[T1 comparable](v1 T1, err) T1`                | Check error + value non-zero, return 1 value     |
| `C2[T1,T2 comparable](v1 T1, v2 T2, err) (T1,T2)` | Check error + values non-zero, return 2 values   |
| `C3` ~ `C9`                                       | Check error + values non-zero, return 3-9 values |

> **Note:** `C0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

### A Series - Array/Slice Validation

| Function                                       | Description                                           |
| ---------------------------------------------- | ----------------------------------------------------- |
| `A1[T1](v1 []T1, err) []T1`                    | Check error + slice has elements, return 1 slice      |
| `A2[T1,T2](v1 []T1, v2 []T2, err) ([]T1,[]T2)` | Check error + slices have elements, return 2 slices   |
| `A3` ~ `A9`                                    | Check error + slices have elements, return 3-9 slices |

> **Note:** `A0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

### M Series - Map Validation

| Function                                                                  | Description                                      |
| ------------------------------------------------------------------------- | ------------------------------------------------ |
| `M1[K1,V1](m1 map[K1]V1, err) map[K1]V1`                                  | Check error + map has entries, return 1 map      |
| `M2[K1,K2,V1,V2](m1 map[K1]V1, m2 map[K2]V2, err) (map[K1]V1, map[K2]V2)` | Check error + maps have entries, return 2 maps   |
| `M3` ~ `M9`                                                               | Check error + maps have entries, return 3-9 maps |

> **Note:** `M0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

### F Series - Function Execution (Value)

| Function                                        | Description                              |
| ----------------------------------------------- | ---------------------------------------- |
| `F1[T1](run func() (T1, error)) T1`             | Execute function, apply V1 validation    |
| `F2[T1,T2](run func() (T1, T2, error)) (T1,T2)` | Execute function, apply V2 validation    |
| `F3` ~ `F9`                                     | Execute function, apply V3-V9 validation |

> **Note:** `F0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

### R Series - Function Execution (Pointer)

| Function                                            | Description                              |
| --------------------------------------------------- | ---------------------------------------- |
| `R1[T1](run func() (*T1, error)) *T1`               | Execute function, apply P1 validation    |
| `R2[T1,T2](run func() (*T1, *T2, error)) (*T1,*T2)` | Execute function, apply P2 validation    |
| `R3` ~ `R9`                                         | Execute function, apply P3-P9 validation |

> **Note:** `R0` is same as `V0`, use [must](https://github.com/yylego/must) package for simple error checking.

---

## resb Package - Boolean Result Handling

The `resb` subpackage handles functions that return `(value, bool)` instead of `(value, error)`.

### Common Use Cases

```go
// Map lookup - returns (value, bool)
func lookupConfig(key string) (string, bool) {
    cfg := map[string]string{"host": "localhost", "port": "8080"}
    return cfg[key]
}


// Using resb.V1 to handle (value, bool) results
config := resb.V1(lookupConfig("host"))      // Panics if key not found
```

---

## Related Projects

Explore more error handling packages in this ecosystem:

### Advanced Packages

- **[must](https://github.com/yylego/must)** - Must-style assertions with rich type support and detailed error context
- **[rese](https://github.com/yylego/rese)** - Result extraction with panic, focused on safe value unwrapping (this project)

### Foundation Packages

- **[done](https://github.com/yylego/done)** - Simple, focused error handling with method chaining
- **[sure](https://github.com/yylego/sure)** - Generates code that creates custom validation methods

Each package targets different use cases, from quick prototyping to production systems with comprehensive error handling.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
