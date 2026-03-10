// Package resb: Boolean result handling - Streamline functions that return (value, bool) patterns
// Transform map lookups and value checks with ok-pattern into clean, panic-on-false validation
// Provides V, P, C series for handling (value, bool) results with type-safe generic functions
// Each series supports 0-9 parameters with bool checking and value validation
//
// resb: 布尔结果处理 - 简化返回 (value, bool) 模式的函数
// 将映射查找和可选值检查转换为简洁的布尔值 false 时 panic 验证
// 提供 V、P、C 系列，通过类型安全泛型函数处理 (value, bool) 结果
// 每个系列支持 0-9 个参数，带有布尔值检查和值验证
package resb

import (
	"github.com/yylego/must"
)

// V0 checks bool flag and panics if false
// Use for functions that return just a bool
//
// V0 检查布尔标志，如果为 false 则 panic
// 用于返回单个布尔值的函数
func V0(ok bool) {
	must.True(ok)
}

// V1 checks bool flag and returns single value
// Panics if bool is false, otherwise returns the value
//
// V1 检查布尔标志并返回单个值
// 如果布尔值为 false 则 panic，否则返回值
func V1[T1 any](v1 T1, ok bool) T1 {
	must.True(ok)
	return v1
}

// V2 checks bool flag and returns two values
// Panics if bool is false, otherwise returns both values
//
// V2 检查布尔标志并返回两个值
// 如果布尔值为 false 则 panic，否则返回两个值
func V2[T1, T2 any](v1 T1, v2 T2, ok bool) (T1, T2) {
	must.True(ok)
	return v1, v2
}

// V3 checks bool flag and returns three values
// Panics when bool is false and returns the values when true
//
// V3 检查布尔标志并返回三个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, ok bool) (T1, T2, T3) {
	must.True(ok)
	return v1, v2, v3
}

// V4 checks bool flag and returns 4 values
// Panics when bool is false and returns the values when true
//
// V4 检查布尔标志并返回 4 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, ok bool) (T1, T2, T3, T4) {
	must.True(ok)
	return v1, v2, v3, v4
}

// V5 checks bool flag and returns 5 values
// Panics when bool is false and returns the values when true
//
// V5 检查布尔标志并返回 5 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, ok bool) (T1, T2, T3, T4, T5) {
	must.True(ok)
	return v1, v2, v3, v4, v5
}

// V6 checks bool flag and returns 6 values
// Panics when bool is false and returns the values when true
//
// V6 检查布尔标志并返回 6 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, ok bool) (T1, T2, T3, T4, T5, T6) {
	must.True(ok)
	return v1, v2, v3, v4, v5, v6
}

// V7 checks bool flag and returns 7 values
// Panics when bool is false and returns the values when true
//
// V7 检查布尔标志并返回 7 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, ok bool) (T1, T2, T3, T4, T5, T6, T7) {
	must.True(ok)
	return v1, v2, v3, v4, v5, v6, v7
}

// V8 checks bool flag and returns 8 values
// Panics when bool is false and returns the values when true
//
// V8 检查布尔标志并返回 8 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, ok bool) (T1, T2, T3, T4, T5, T6, T7, T8) {
	must.True(ok)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// V9 checks bool flag and returns 9 values
// Panics when bool is false and returns the values when true
//
// V9 检查布尔标志并返回 9 个值
// 当布尔值为 false 时 panic，为 true 时返回值
func V9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, ok bool) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	must.True(ok)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

// P0 checks bool flag and panics if false
// Same as V0 - validates just a bool
//
// P0 检查布尔标志，如果为 false 则 panic
// 与 V0 相同 - 验证单个布尔值
func P0(ok bool) {
	must.True(ok)
}

// P1 checks bool flag and pointer non-nil, returns single pointer
// Panics when bool is false and when pointer is nil
//
// P1 检查布尔标志和指针非 nil，返回单个指针
// 当布尔值为 false 时 panic，当指针为 nil 时 panic
func P1[T1 any](v1 *T1, ok bool) *T1 {
	must.True(ok)
	must.Full(v1)
	return v1
}

// P2 checks bool flag and pointers non-nil, returns two pointers
// Panics when bool is false and when some pointer is nil
//
// P2 检查布尔标志和指针非 nil，返回两个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P2[T1, T2 any](v1 *T1, v2 *T2, ok bool) (*T1, *T2) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	return v1, v2
}

// P3 checks bool flag and pointers non-nil, returns three pointers
// Panics when bool is false and when some pointer is nil
//
// P3 检查布尔标志和指针非 nil，返回三个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P3[T1, T2, T3 any](v1 *T1, v2 *T2, v3 *T3, ok bool) (*T1, *T2, *T3) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	return v1, v2, v3
}

// P4 checks bool flag and pointers non-nil, returns 4 pointers
// Panics when bool is false and when some pointer is nil
//
// P4 检查布尔标志和指针非 nil，返回 4 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P4[T1, T2, T3, T4 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, ok bool) (*T1, *T2, *T3, *T4) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	return v1, v2, v3, v4
}

// P5 checks bool flag and pointers non-nil, returns 5 pointers
// Panics when bool is false and when some pointer is nil
//
// P5 检查布尔标志和指针非 nil，返回 5 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P5[T1, T2, T3, T4, T5 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, ok bool) (*T1, *T2, *T3, *T4, *T5) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	must.Full(v5)
	return v1, v2, v3, v4, v5
}

// P6 checks bool flag and pointers non-nil, returns 6 pointers
// Panics when bool is false and when some pointer is nil
//
// P6 检查布尔标志和指针非 nil，返回 6 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P6[T1, T2, T3, T4, T5, T6 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, ok bool) (*T1, *T2, *T3, *T4, *T5, *T6) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	must.Full(v5)
	must.Full(v6)
	return v1, v2, v3, v4, v5, v6
}

// P7 checks bool flag and pointers non-nil, returns 7 pointers
// Panics when bool is false and when some pointer is nil
//
// P7 检查布尔标志和指针非 nil，返回 7 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P7[T1, T2, T3, T4, T5, T6, T7 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, ok bool) (*T1, *T2, *T3, *T4, *T5, *T6, *T7) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	must.Full(v5)
	must.Full(v6)
	must.Full(v7)
	return v1, v2, v3, v4, v5, v6, v7
}

// P8 checks bool flag and pointers non-nil, returns 8 pointers
// Panics when bool is false and when some pointer is nil
//
// P8 检查布尔标志和指针非 nil，返回 8 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, ok bool) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	must.Full(v5)
	must.Full(v6)
	must.Full(v7)
	must.Full(v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// P9 checks bool flag and pointers non-nil, returns 9 pointers
// Panics when bool is false and when some pointer is nil
//
// P9 检查布尔标志和指针非 nil，返回 9 个指针
// 当布尔值为 false 时 panic，当某个指针为 nil 时 panic
func P9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, v9 *T9, ok bool) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9) {
	must.True(ok)
	must.Full(v1)
	must.Full(v2)
	must.Full(v3)
	must.Full(v4)
	must.Full(v5)
	must.Full(v6)
	must.Full(v7)
	must.Full(v8)
	must.Full(v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

// C0 checks bool flag and panics if false
// Same as V0 - validates just a bool
//
// C0 检查布尔标志，如果为 false 则 panic
// 与 V0 相同 - 验证单个布尔值
func C0(ok bool) {
	must.True(ok)
}

// C1 checks bool flag and value non-zero, returns single comparable value
// Panics when bool is false and when value is zero
//
// C1 检查布尔标志和值非零，返回单个可比较值
// 当布尔值为 false 时 panic，当值为零时 panic
func C1[T1 comparable](v1 T1, ok bool) T1 {
	must.True(ok)
	must.Nice(v1)
	return v1
}

// C2 checks bool flag and values non-zero, returns two comparable values
// Panics when bool is false and when some value is zero
//
// C2 检查布尔标志和值非零，返回两个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C2[T1, T2 comparable](v1 T1, v2 T2, ok bool) (T1, T2) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	return v1, v2
}

// C3 checks bool flag and values non-zero, returns three comparable values
// Panics when bool is false and when some value is zero
//
// C3 检查布尔标志和值非零，返回三个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C3[T1, T2, T3 comparable](v1 T1, v2 T2, v3 T3, ok bool) (T1, T2, T3) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	return v1, v2, v3
}

// C4 checks bool flag and values non-zero, returns 4 comparable values
// Panics when bool is false and when some value is zero
//
// C4 检查布尔标志和值非零，返回 4 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C4[T1, T2, T3, T4 comparable](v1 T1, v2 T2, v3 T3, v4 T4, ok bool) (T1, T2, T3, T4) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	return v1, v2, v3, v4
}

// C5 checks bool flag and values non-zero, returns 5 comparable values
// Panics when bool is false and when some value is zero
//
// C5 检查布尔标志和值非零，返回 5 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C5[T1, T2, T3, T4, T5 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, ok bool) (T1, T2, T3, T4, T5) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	must.Nice(v5)
	return v1, v2, v3, v4, v5
}

// C6 checks bool flag and values non-zero, returns 6 comparable values
// Panics when bool is false and when some value is zero
//
// C6 检查布尔标志和值非零，返回 6 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C6[T1, T2, T3, T4, T5, T6 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, ok bool) (T1, T2, T3, T4, T5, T6) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	must.Nice(v5)
	must.Nice(v6)
	return v1, v2, v3, v4, v5, v6
}

// C7 checks bool flag and values non-zero, returns 7 comparable values
// Panics when bool is false and when some value is zero
//
// C7 检查布尔标志和值非零，返回 7 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C7[T1, T2, T3, T4, T5, T6, T7 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, ok bool) (T1, T2, T3, T4, T5, T6, T7) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	must.Nice(v5)
	must.Nice(v6)
	must.Nice(v7)
	return v1, v2, v3, v4, v5, v6, v7
}

// C8 checks bool flag and values non-zero, returns 8 comparable values
// Panics when bool is false and when some value is zero
//
// C8 检查布尔标志和值非零，返回 8 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C8[T1, T2, T3, T4, T5, T6, T7, T8 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, ok bool) (T1, T2, T3, T4, T5, T6, T7, T8) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	must.Nice(v5)
	must.Nice(v6)
	must.Nice(v7)
	must.Nice(v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

// C9 checks bool flag and values non-zero, returns 9 comparable values
// Panics when bool is false and when some value is zero
//
// C9 检查布尔标志和值非零，返回 9 个可比较值
// 当布尔值为 false 时 panic，当某个值为零时 panic
func C9[T1, T2, T3, T4, T5, T6, T7, T8, T9 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, ok bool) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	must.True(ok)
	must.Nice(v1)
	must.Nice(v2)
	must.Nice(v3)
	must.Nice(v4)
	must.Nice(v5)
	must.Nice(v6)
	must.Nice(v7)
	must.Nice(v8)
	must.Nice(v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}
