package modules

import (
	"fmt"
	"net/url"
)

const (
	AppPageFormat = "https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=%s&scope=%s"
)

// NewAppPage 用于创建微信二维页的 url.
// 扫码成功后所要重定向的 url 已做查询编码,
// redURL 参数传入编码或未编码的都可.
func NewAppPage(appID, redURL string) (*AppPage, error) {
	decode, err := url.QueryUnescape(redURL)
	if err != nil {
		return nil, err
	}
	redURL = url.QueryEscape(decode)

	ap := &AppPage{
		appID:        appID,
		redirectURL:  redURL,
		responseType: "code",
		scope:        "snsapi_login",
		state:        "",
	}
	return ap, nil
}

// AppPage 用于描述某个第三方网站在使用微信扫码登录时的跳转连接
type AppPage struct {
	appID        string
	redirectURL  string // 需要经过 url 编码
	responseType string // 固定为 "code"
	scope        string // 固定为 "snsapi_login"
	state        string // 不是必须的, 暂时固定为空串 ""
}

// todo: 实现
func (ap *AppPage) String() string {
	return fmt.Sprintf(AppPageFormat, ap.appID, ap.redirectURL, ap.responseType, ap.scope)
}
