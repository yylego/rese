// Package rese: Result extraction with error handling - Streamline Go error handling patterns
// Transform verbose error checking into clean, expressive code with type-safe generic functions
// Provides multiple function series (V, P, C, A) for different validation scenarios
// Each series supports 0-9 parameters with error checking and value validation
//
// rese: 结果提取与错误处理 - 简化 Go 错误处理模式
// 将冗长的错误检查转换为简洁、富有表现力的类型安全泛型函数代码
// 提供多个函数系列（V、P、C、A）用于不同的验证场景
// 每个系列支持 0-9 个参数，带有错误检查和值验证
package rese

import (
	"github.com/yylego/rese/internal/resecheck"
)

// V0 checks error and panics if present
// Use github.com/yylego/must package for V0 scenario
//
// V0 检查错误，如果存在则 panic
// V0 场景建议使用 github.com/yylego/must 包
func V0(err error) {
	resecheck.Done(1, err)
}

// V1 checks error and returns single value
// Panics when error is present and returns the value when not
//
// V1 检查错误并返回单个值
// 当有错误时 panic，没有错误时返回值
func V1[T1 any](v1 T1, err error) T1 {
	resecheck.Done(1, err)
	return v1
}

// V2 checks error and returns two values
// Panics when error is present and returns the values when not
//
// V2 检查错误并返回两个值
// 当有错误时 panic，没有错误时返回值
func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	resecheck.Done(1, err)
	return v1, v2
}

// V3 checks error and returns three values
// Panics when error is present and returns the values when not
//
// V3 检查错误并返回三个值
// 当有错误时 panic，没有错误时返回值
func V3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	resecheck.Done(1, err)
	return v1, v2, v3
}

// V4 checks error and returns 4 values
// Panics when error is present and returns the values when not
//
// V4 检查错误并返回 4 个值
// 当有错误时 panic，没有错误时返回值
func V4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4
}

// V5 checks error and returns 5 values
// Panics when error is present and returns the values when not
//
// V5 检查错误并返回 5 个值
// 当有错误时 panic，没有错误时返回值
func V5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4, v5
}

// V6 checks error and returns 6 values
// Panics when error is present and returns the values when not
//
// V6 检查错误并返回 6 个值
// 当有错误时 panic，没有错误时返回值
func V6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4, v5, v6
}

// V7 checks error and returns 7 values
// Panics when error is present and returns the values when not
//
// V7 检查错误并返回 7 个值
// 当有错误时 panic，没有错误时返回值
func V7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4, v5, v6, v7
}

// V8 checks error and returns 8 values
// Panics when error is present and returns the values when not
//
// V8 检查错误并返回 8 个值
// 当有错误时 panic，没有错误时返回值
func V8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// V9 checks error and returns 9 values
// Panics when error is present and returns the values when not
//
// V9 检查错误并返回 9 个值
// 当有错误时 panic，没有错误时返回值
func V9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	resecheck.Done(1, err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

// P0 checks error and panics if present
// Use github.com/yylego/must package for P0 scenario
//
// P0 检查错误，如果存在则 panic
// P0 场景建议使用 github.com/yylego/must 包
func P0(err error) {
	resecheck.Done(1, err)
}

// P1 checks error and validates pointer is non-nil, returns single pointer
// Panics when error is present and when pointer is nil
//
// P1 检查错误并验证指针非 nil，返回单个指针
// 当有错误时 panic，当指针为 nil 时 panic
func P1[T1 any](v1 *T1, err error) *T1 {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	return v1
}

// P2 checks error and validates pointers non-nil, returns two pointers
// Panics when error is present and when some pointer is nil
//
// P2 检查错误并验证指针非 nil，返回两个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	return v1, v2
}

// P3 checks error and validates pointers non-nil, returns three pointers
// Panics when error is present and when some pointer is nil
//
// P3 检查错误并验证指针非 nil，返回三个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P3[T1, T2, T3 any](v1 *T1, v2 *T2, v3 *T3, err error) (*T1, *T2, *T3) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	return v1, v2, v3
}

// P4 checks error and validates pointers non-nil, returns 4 pointers
// Panics when error is present and when some pointer is nil
//
// P4 检查错误并验证指针非 nil，返回 4 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P4[T1, T2, T3, T4 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, err error) (*T1, *T2, *T3, *T4) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	return v1, v2, v3, v4
}

// P5 checks error and validates pointers non-nil, returns 5 pointers
// Panics when error is present and when some pointer is nil
//
// P5 检查错误并验证指针非 nil，返回 5 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P5[T1, T2, T3, T4, T5 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, err error) (*T1, *T2, *T3, *T4, *T5) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	resecheck.Full(1, v5)
	return v1, v2, v3, v4, v5
}

// P6 checks error and validates pointers non-nil, returns 6 pointers
// Panics when error is present and when some pointer is nil
//
// P6 检查错误并验证指针非 nil，返回 6 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P6[T1, T2, T3, T4, T5, T6 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, err error) (*T1, *T2, *T3, *T4, *T5, *T6) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	resecheck.Full(1, v5)
	resecheck.Full(1, v6)
	return v1, v2, v3, v4, v5, v6
}

// P7 checks error and validates pointers non-nil, returns 7 pointers
// Panics when error is present and when some pointer is nil
//
// P7 检查错误并验证指针非 nil，返回 7 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P7[T1, T2, T3, T4, T5, T6, T7 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	resecheck.Full(1, v5)
	resecheck.Full(1, v6)
	resecheck.Full(1, v7)
	return v1, v2, v3, v4, v5, v6, v7
}

// P8 checks error and validates pointers non-nil, returns 8 pointers
// Panics when error is present and when some pointer is nil
//
// P8 检查错误并验证指针非 nil，返回 8 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	resecheck.Full(1, v5)
	resecheck.Full(1, v6)
	resecheck.Full(1, v7)
	resecheck.Full(1, v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// P9 checks error and validates pointers non-nil, returns 9 pointers
// Panics when error is present and when some pointer is nil
//
// P9 检查错误并验证指针非 nil，返回 9 个指针
// 当有错误时 panic，当某个指针为 nil 时 panic
func P9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, v9 *T9, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9) {
	resecheck.Done(1, err)
	resecheck.Full(1, v1)
	resecheck.Full(1, v2)
	resecheck.Full(1, v3)
	resecheck.Full(1, v4)
	resecheck.Full(1, v5)
	resecheck.Full(1, v6)
	resecheck.Full(1, v7)
	resecheck.Full(1, v8)
	resecheck.Full(1, v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

// C0 checks error and panics if present
// Use github.com/yylego/must package for C0 scenario
//
// C0 检查错误，如果存在则 panic
// C0 场景建议使用 github.com/yylego/must 包
func C0(err error) {
	resecheck.Done(1, err)
}

// C1 checks error and value non-zero, returns single comparable value
// Panics on error and panics when value is zero
//
// C1 检查错误和值非零，返回单个可比较值
// 当有错误时 panic，当值为零时 panic
func C1[T1 comparable](v1 T1, err error) T1 {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	return v1
}

// C2 checks error and values non-zero, returns two comparable values
// Panics on error and panics when some value is zero
//
// C2 检查错误和值非零，返回两个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	return v1, v2
}

// C3 checks error and values non-zero, returns three comparable values
// Panics on error and panics when some value is zero
//
// C3 检查错误和值非零，返回三个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C3[T1, T2, T3 comparable](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	return v1, v2, v3
}

// C4 checks error and values non-zero, returns 4 comparable values
// Panics on error and panics when some value is zero
//
// C4 检查错误和值非零，返回 4 个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C4[T1, T2, T3, T4 comparable](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	return v1, v2, v3, v4
}

// C5 checks error and values non-zero, returns five comparable values
// Panics on error and panics when some value is zero
//
// C5 检查错误和值非零，返回五个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C5[T1, T2, T3, T4, T5 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	resecheck.Nice(1, v5)
	return v1, v2, v3, v4, v5
}

// C6 checks error and values non-zero, returns six comparable values
// Panics on error and panics when some value is zero
//
// C6 检查错误和值非零，返回六个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C6[T1, T2, T3, T4, T5, T6 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	resecheck.Nice(1, v5)
	resecheck.Nice(1, v6)
	return v1, v2, v3, v4, v5, v6
}

// C7 checks error and values non-zero, returns seven comparable values
// Panics on error and panics when some value is zero
//
// C7 检查错误和值非零，返回七个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C7[T1, T2, T3, T4, T5, T6, T7 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	resecheck.Nice(1, v5)
	resecheck.Nice(1, v6)
	resecheck.Nice(1, v7)
	return v1, v2, v3, v4, v5, v6, v7
}

// C8 checks error and values non-zero, returns eight comparable values
// Panics on error and panics when some value is zero
//
// C8 检查错误和值非零，返回八个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C8[T1, T2, T3, T4, T5, T6, T7, T8 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	resecheck.Nice(1, v5)
	resecheck.Nice(1, v6)
	resecheck.Nice(1, v7)
	resecheck.Nice(1, v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// C9 checks error and values non-zero, returns nine comparable values
// Panics on error and panics when some value is zero
//
// C9 检查错误和值非零，返回九个可比较值
// 当有错误时 panic，当某个值为零时 panic
func C9[T1, T2, T3, T4, T5, T6, T7, T8, T9 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	resecheck.Done(1, err)
	resecheck.Nice(1, v1)
	resecheck.Nice(1, v2)
	resecheck.Nice(1, v3)
	resecheck.Nice(1, v4)
	resecheck.Nice(1, v5)
	resecheck.Nice(1, v6)
	resecheck.Nice(1, v7)
	resecheck.Nice(1, v8)
	resecheck.Nice(1, v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

// A0 checks error and panics if present
// Use github.com/yylego/must package for A0 scenario
//
// A0 检查错误，如果存在则 panic
// A0 场景建议使用 github.com/yylego/must 包
func A0(err error) {
	resecheck.Done(1, err)
}

// A1 checks error and validates slice has elements, returns single slice
// Panics on error and panics when slice is nil/has no elements
//
// A1 检查错误并验证切片有元素，返回单个切片
// 当有错误时 panic，当切片为 nil/无元素时 panic
func A1[T1 any](v1 []T1, err error) []T1 {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	return v1
}

// A2 checks error and slices have elements, returns two slices
// Panics on error and panics when some slice is nil/has no elements
//
// A2 检查错误和切片非空，返回两个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A2[T1, T2 any](v1 []T1, v2 []T2, err error) ([]T1, []T2) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	return v1, v2
}

// A3 checks error and slices have elements, returns three slices
// Panics on error and panics when some slice is nil/has no elements
//
// A3 检查错误和切片非空，返回三个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A3[T1, T2, T3 any](v1 []T1, v2 []T2, v3 []T3, err error) ([]T1, []T2, []T3) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	return v1, v2, v3
}

// A4 checks error and slices have elements, returns 4 slices
// Panics on error and panics when some slice is nil/has no elements
//
// A4 检查错误和切片非空，返回 4 个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A4[T1, T2, T3, T4 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, err error) ([]T1, []T2, []T3, []T4) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	return v1, v2, v3, v4
}

// A5 checks error and slices have elements, returns five slices
// Panics on error and panics when some slice is nil/has no elements
//
// A5 检查错误和切片非空，返回五个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A5[T1, T2, T3, T4, T5 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, v5 []T5, err error) ([]T1, []T2, []T3, []T4, []T5) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	resecheck.Have(1, v5)
	return v1, v2, v3, v4, v5
}

// A6 checks error and slices have elements, returns six slices
// Panics on error and panics when some slice is nil/has no elements
//
// A6 检查错误和切片非空，返回六个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A6[T1, T2, T3, T4, T5, T6 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, v5 []T5, v6 []T6, err error) ([]T1, []T2, []T3, []T4, []T5, []T6) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	resecheck.Have(1, v5)
	resecheck.Have(1, v6)
	return v1, v2, v3, v4, v5, v6
}

// A7 checks error and slices have elements, returns seven slices
// Panics on error and panics when some slice is nil/has no elements
//
// A7 检查错误和切片非空，返回七个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A7[T1, T2, T3, T4, T5, T6, T7 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, v5 []T5, v6 []T6, v7 []T7, err error) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	resecheck.Have(1, v5)
	resecheck.Have(1, v6)
	resecheck.Have(1, v7)
	return v1, v2, v3, v4, v5, v6, v7
}

// A8 checks error and slices have elements, returns eight slices
// Panics on error and panics when some slice is nil/has no elements
//
// A8 检查错误和切片非空，返回八个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, v5 []T5, v6 []T6, v7 []T7, v8 []T8, err error) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	resecheck.Have(1, v5)
	resecheck.Have(1, v6)
	resecheck.Have(1, v7)
	resecheck.Have(1, v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// A9 checks error and slices have elements, returns nine slices
// Panics on error and panics when some slice is nil/has no elements
//
// A9 检查错误和切片非空，返回九个切片
// 当有错误时 panic，当某个切片为 nil/无元素时 panic
func A9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 []T1, v2 []T2, v3 []T3, v4 []T4, v5 []T5, v6 []T6, v7 []T7, v8 []T8, v9 []T9, err error) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9) {
	resecheck.Done(1, err)
	resecheck.Have(1, v1)
	resecheck.Have(1, v2)
	resecheck.Have(1, v3)
	resecheck.Have(1, v4)
	resecheck.Have(1, v5)
	resecheck.Have(1, v6)
	resecheck.Have(1, v7)
	resecheck.Have(1, v8)
	resecheck.Have(1, v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}
