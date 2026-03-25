package application

import (
	"strconv"
	"theapp/constants"
	"time"
)

// StringToInt は文字列を整数に変換する。解析に失敗した場合は 0 を返す。
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// StringToTime はアプリ共通の日時書式で文字列を時刻に変換する。失敗時はゼロ値を返す。
func StringToTime(s string) time.Time {
	t, _ := time.Parse(constants.DateTimeFormat, s)
	return t
}
