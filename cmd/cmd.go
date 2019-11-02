package cmd

import "net/http"

// Cookies 是 wxlogin 的核心,
// 其目的就是模拟浏览器行为, 获取第三方应用发来的 cookies.
// todo: 实现
func Cookies() []*http.Cookie {
	return nil
}
