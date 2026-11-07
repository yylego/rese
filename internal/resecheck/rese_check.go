// Package resecheck holds the checks behind rese and resb, each one taking the skip of the frame the log reports
// Each exported function passes its own depth, so the PANIC log points at the code that calls rese, not at a line inside rese
// Each check panics with an error, so recover() gets an error back and errors.Is keeps working
//
// resecheck 存放 rese 和 resb 背后的检查，每个检查都带着日志要报告的那一层的 skip
// 导出函数传自己的深度，PANIC 日志就指向调用 rese 的代码，而不是 rese 内部的某一行
// 每个检查都用 error panic，recover() 拿到的是 error，errors.Is 照常可用
package resecheck

import (
	"github.com/pkg/errors"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// Done panics with the error itself when the error is not nil
// skip counts the frames above the function that calls Done to the code the log reports: 1 is the code that calls that function
//
// Done 在错误非 nil 时用这个错误本身 panic
// skip 从调用 Done 的函数往上数到日志要报告的那一层：1 就是调用那个函数的代码
func Done(skip int, err error) {
	if err != nil {
		panicWithCause(skip+1, err, "EXPECTED NO ERROR(BUT HAS ERROR)", zap.Error(err))
	}
}

// True panics when the value is false, skip counts the frames as Done does
// True 在值为 false 时 panic，skip 的数法跟 Done 一样
func True(skip int, v bool) {
	if !v {
		panicWithMessage(skip+1, "VALUE IS FALSE(SHOULD BE TRUE)", zap.Bool("v", v))
	}
}

// Full panics when the value is nil, skip counts the frames as Done does
// Full 在指针为 nil 时 panic，skip 的数法跟 Done 一样
func Full[T any](skip int, v *T) {
	if v == nil {
		panicWithMessage(skip+1, "VALUE ABSENT(SHOULD BE PRESENT)")
	}
}

// Nice panics when the value is zero, skip counts the frames as Done does
// Nice 在值为零值时 panic，skip 的数法跟 Done 一样
func Nice[V comparable](skip int, a V) {
	var zero V
	if a == zero {
		panicWithMessage(skip+1, "VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
}

// Have panics when the slice has no elements, skip counts the frames as Done does
// Have 在切片没有元素时 panic，skip 的数法跟 Done 一样
func Have[T any](skip int, a []T) {
	if len(a) == 0 {
		panicWithMessage(skip+1, "SLICE IS EMPTY(SHOULD HAVE ITEMS)")
	}
}

// HaveMap panics when the map has no entries, skip counts the frames as Done does
// HaveMap 在映射没有条目时 panic，skip 的数法跟 Done 一样
func HaveMap[K comparable, V any](skip int, m map[K]V) {
	if len(m) == 0 {
		panicWithMessage(skip+1, "MAP IS EMPTY(SHOULD HAVE ITEMS)")
	}
}

// panicWithMessage panics with an error made from the message
// panicWithMessage 用由消息生成的 error panic
func panicWithMessage(skip int, message string, fields ...zap.Field) {
	panicWithCause(skip+1, errors.New(message), message, fields...)
}

// panicWithCause writes the PANIC log, then panics with the cause itself
// panicWithCause 写 PANIC 日志，然后用 cause 本身 panic
func panicWithCause(skip int, cause error, message string, fields ...zap.Field) {
	zaplog.LOGS.Skip(skip+1).WithOptions(zaplog.WithPanicCause(cause)).Panic(message, fields...)
}
