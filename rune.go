package rese

// F0 executes function and checks error, applying V0 validation
// Use github.com/yylego/must package for F0 scenario
//
// F0 执行函数并检查错误，应用 V0 验证
// F0 场景建议使用 github.com/yylego/must 包
func F0(run func() error) {
	V0(run())
}

// F1 executes function and checks the result, applying V1 validation
// Returns single value from function and panics when error is present
//
// F1 执行函数并检查结果，应用 V1 验证
// 返回函数的单个值，当有错误时 panic
func F1[T1 any](run func() (T1, error)) T1 {
	return V1(run())
}

// F2 executes function and checks the result, applying V2 validation
// Returns two values from function and panics when error is present
//
// F2 执行函数并检查结果，应用 V2 验证
// 返回函数的两个值，当有错误时 panic
func F2[T1, T2 any](run func() (T1, T2, error)) (T1, T2) {
	return V2(run())
}

// F3 executes function and checks the result, applying V3 validation
// Returns three values from function and panics when error is present
//
// F3 执行函数并检查结果，应用 V3 验证
// 返回函数的三个值，当有错误时 panic
func F3[T1, T2, T3 any](run func() (T1, T2, T3, error)) (T1, T2, T3) {
	return V3(run())
}

// F4 executes function and checks the result, applying V4 validation
// Returns 4 values from function and panics when error is present
//
// F4 执行函数并检查结果，应用 V4 验证
// 返回函数的 4 个值，当有错误时 panic
func F4[T1, T2, T3, T4 any](run func() (T1, T2, T3, T4, error)) (T1, T2, T3, T4) {
	return V4(run())
}

// F5 executes function and checks the result, applying V5 validation
// Returns 5 values from function and panics when error is present
//
// F5 执行函数并检查结果，应用 V5 验证
// 返回函数的 5 个值，当有错误时 panic
func F5[T1, T2, T3, T4, T5 any](run func() (T1, T2, T3, T4, T5, error)) (T1, T2, T3, T4, T5) {
	return V5(run())
}

// F6 executes function and checks the result, applying V6 validation
// Returns 6 values from function and panics when error is present
//
// F6 执行函数并检查结果，应用 V6 验证
// 返回函数的 6 个值，当有错误时 panic
func F6[T1, T2, T3, T4, T5, T6 any](run func() (T1, T2, T3, T4, T5, T6, error)) (T1, T2, T3, T4, T5, T6) {
	return V6(run())
}

// F7 executes function and checks the result, applying V7 validation
// Returns 7 values from function and panics when error is present
//
// F7 执行函数并检查结果，应用 V7 验证
// 返回函数的 7 个值，当有错误时 panic
func F7[T1, T2, T3, T4, T5, T6, T7 any](run func() (T1, T2, T3, T4, T5, T6, T7, error)) (T1, T2, T3, T4, T5, T6, T7) {
	return V7(run())
}

// F8 executes function and checks the result, applying V8 validation
// Returns 8 values from function and panics when error is present
//
// F8 执行函数并检查结果，应用 V8 验证
// 返回函数的 8 个值，当有错误时 panic
func F8[T1, T2, T3, T4, T5, T6, T7, T8 any](run func() (T1, T2, T3, T4, T5, T6, T7, T8, error)) (T1, T2, T3, T4, T5, T6, T7, T8) {
	return V8(run())
}

// F9 executes function and checks the result, applying V9 validation
// Returns 9 values from function and panics when error is present
//
// F9 执行函数并检查结果，应用 V9 验证
// 返回函数的 9 个值，当有错误时 panic
func F9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](run func() (T1, T2, T3, T4, T5, T6, T7, T8, T9, error)) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return V9(run())
}

// R0 executes function and checks error, applying P0 validation
// Use github.com/yylego/must package for R0 scenario
//
// R0 执行函数并检查错误，应用 P0 验证
// R0 场景建议使用 github.com/yylego/must 包
func R0(run func() error) {
	P0(run())
}

// R1 executes function and checks the result with pointer validation, applying P1 validation
// Returns single non-nil pointer from function and panics when error/nil is present
//
// R1 执行函数并检查结果和指针验证，应用 P1 验证
// 返回函数的单个非 nil 指针，当有错误/nil 时 panic
func R1[T1 any](run func() (*T1, error)) *T1 {
	return P1(run())
}

// R2 executes function and checks the result with pointer validation, applying P2 validation
// Returns two non-nil pointers from function and panics when error/nil is present
//
// R2 执行函数并检查结果和指针验证，应用 P2 验证
// 返回函数的两个非 nil 指针，当有错误/nil 时 panic
func R2[T1, T2 any](run func() (*T1, *T2, error)) (*T1, *T2) {
	return P2(run())
}

// R3 executes function and checks the result with pointer validation, applying P3 validation
// Returns three non-nil pointers from function and panics when error/nil is present
//
// R3 执行函数并检查结果和指针验证，应用 P3 验证
// 返回函数的三个非 nil 指针，当有错误/nil 时 panic
func R3[T1, T2, T3 any](run func() (*T1, *T2, *T3, error)) (*T1, *T2, *T3) {
	return P3(run())
}

// R4 executes function and checks the result with pointer validation, applying P4 validation
// Returns 4 non-nil pointers from function and panics when error/nil is present
//
// R4 执行函数并检查结果和指针验证，应用 P4 验证
// 返回函数的 4 个非 nil 指针，当有错误/nil 时 panic
func R4[T1, T2, T3, T4 any](run func() (*T1, *T2, *T3, *T4, error)) (*T1, *T2, *T3, *T4) {
	return P4(run())
}

// R5 executes function and checks the result with pointer validation, applying P5 validation
// Returns 5 non-nil pointers from function and panics when error/nil is present
//
// R5 执行函数并检查结果和指针验证，应用 P5 验证
// 返回函数的 5 个非 nil 指针，当有错误/nil 时 panic
func R5[T1, T2, T3, T4, T5 any](run func() (*T1, *T2, *T3, *T4, *T5, error)) (*T1, *T2, *T3, *T4, *T5) {
	return P5(run())
}

// R6 executes function and checks the result with pointer validation, applying P6 validation
// Returns 6 non-nil pointers from function and panics when error/nil is present
//
// R6 执行函数并检查结果和指针验证，应用 P6 验证
// 返回函数的 6 个非 nil 指针，当有错误/nil 时 panic
func R6[T1, T2, T3, T4, T5, T6 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, error)) (*T1, *T2, *T3, *T4, *T5, *T6) {
	return P6(run())
}

// R7 executes function and checks the result with pointer validation, applying P7 validation
// Returns 7 non-nil pointers from function and panics when error/nil is present
//
// R7 执行函数并检查结果和指针验证，应用 P7 验证
// 返回函数的 7 个非 nil 指针，当有错误/nil 时 panic
func R7[T1, T2, T3, T4, T5, T6, T7 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7) {
	return P7(run())
}

// R8 executes function and checks the result with pointer validation, applying P8 validation
// Returns 8 non-nil pointers from function and panics when error/nil is present
//
// R8 执行函数并检查结果和指针验证，应用 P8 验证
// 返回函数的 8 个非 nil 指针，当有错误/nil 时 panic
func R8[T1, T2, T3, T4, T5, T6, T7, T8 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8) {
	return P8(run())
}

// R9 executes function and checks the result with pointer validation, applying P9 validation
// Returns 9 non-nil pointers from function and panics when error/nil is present
//
// R9 执行函数并检查结果和指针验证，应用 P9 验证
// 返回函数的 9 个非 nil 指针，当有错误/nil 时 panic
func R9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](run func() (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9, error)) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9) {
	return P9(run())
}
