package rese

import (
	"github.com/yylego/done"
	"github.com/yylego/must/mustmap"
)

// M0 checks error and panics if present
// Use github.com/yylego/must package for M0 scenario
//
// M0 检查错误，如果存在则 panic
// M0 场景建议使用 github.com/yylego/must 包
func M0(err error) {
	done.Done(err)
}

// M1 checks error and map has entries, returns single map
// Panics when error is present and when map is nil/has no entries
//
// M1 检查错误和映射有条目，返回单个映射
// 当有错误时 panic，当映射为 nil/无条目时 panic
func M1[K1 comparable, V1 any](m1 map[K1]V1, err error) map[K1]V1 {
	done.Done(err)
	mustmap.Have(m1)
	return m1
}

// M2 checks error and maps have entries, returns two maps
// Panics when error is present and when some map is nil/has no entries
//
// M2 检查错误和映射有条目，返回两个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M2[K1, K2 comparable, V1, V2 any](m1 map[K1]V1, m2 map[K2]V2, err error) (map[K1]V1, map[K2]V2) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	return m1, m2
}

// M3 checks error and maps have entries, returns three maps
// Panics when error is present and when some map is nil/has no entries
//
// M3 检查错误和映射有条目，返回三个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M3[K1, K2, K3 comparable, V1, V2, V3 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, err error) (map[K1]V1, map[K2]V2, map[K3]V3) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	return m1, m2, m3
}

// M4 checks error and maps have entries, returns 4 maps
// Panics when error is present and when some map is nil/has no entries
//
// M4 检查错误和映射有条目，返回 4 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M4[K1, K2, K3, K4 comparable, V1, V2, V3, V4 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	return m1, m2, m3, m4
}

// M5 checks error and maps have entries, returns 5 maps
// Panics when error is present and when some map is nil/has no entries
//
// M5 检查错误和映射有条目，返回 5 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M5[K1, K2, K3, K4, K5 comparable, V1, V2, V3, V4, V5 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	return m1, m2, m3, m4, m5
}

// M6 checks error and maps have entries, returns 6 maps
// Panics when error is present and when some map is nil/has no entries
//
// M6 检查错误和映射有条目，返回 6 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M6[K1, K2, K3, K4, K5, K6 comparable, V1, V2, V3, V4, V5, V6 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	return m1, m2, m3, m4, m5, m6
}

// M7 checks error and maps have entries, returns 7 maps
// Panics when error is present and when some map is nil/has no entries
//
// M7 检查错误和映射有条目，返回 7 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M7[K1, K2, K3, K4, K5, K6, K7 comparable, V1, V2, V3, V4, V5, V6, V7 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	return m1, m2, m3, m4, m5, m6, m7
}

// M8 checks error and maps have entries, returns 8 maps
// Panics when error is present and when some map is nil/has no entries
//
// M8 检查错误和映射有条目，返回 8 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M8[K1, K2, K3, K4, K5, K6, K7, K8 comparable, V1, V2, V3, V4, V5, V6, V7, V8 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, m8 map[K8]V8, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7, map[K8]V8) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	mustmap.Have(m8)
	return m1, m2, m3, m4, m5, m6, m7, m8
}

// M9 checks error and maps have entries, returns 9 maps
// Panics when error is present and when some map is nil/has no entries
//
// M9 检查错误和映射有条目，返回 9 个映射
// 当有错误时 panic，当某个映射为 nil/无条目时 panic
func M9[K1, K2, K3, K4, K5, K6, K7, K8, K9 comparable, V1, V2, V3, V4, V5, V6, V7, V8, V9 any](m1 map[K1]V1, m2 map[K2]V2, m3 map[K3]V3, m4 map[K4]V4, m5 map[K5]V5, m6 map[K6]V6, m7 map[K7]V7, m8 map[K8]V8, m9 map[K9]V9, err error) (map[K1]V1, map[K2]V2, map[K3]V3, map[K4]V4, map[K5]V5, map[K6]V6, map[K7]V7, map[K8]V8, map[K9]V9) {
	done.Done(err)
	mustmap.Have(m1)
	mustmap.Have(m2)
	mustmap.Have(m3)
	mustmap.Have(m4)
	mustmap.Have(m5)
	mustmap.Have(m6)
	mustmap.Have(m7)
	mustmap.Have(m8)
	mustmap.Have(m9)
	return m1, m2, m3, m4, m5, m6, m7, m8, m9
}
