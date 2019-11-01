package modules

const (
	AppURLFormat = "https://open.weixin.qq.com/connect/qrconnect?appid=wx47e0850600dd30fc&redirect_uri=https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin&response_type=code&scope=snsapi_login"
)

func NewAppPage(appID, redURL string) *AppPage {
	ap := &AppPage{
		appID:        appID,
		redirectURL:  redURL,
		responseType: "code",
		scope:        "snsapi_login",
		state:        "",
	}
	return ap
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
	return ""
}
