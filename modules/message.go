package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	myTime "wxlogin/utils/time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// 微信二维码页链接
	AppPageFormat = "https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=%s&scope=%s"
	// 二维码链接前缀
	QrcodeURLPref = "https://open.weixin.qq.com"
	// 轮询链接
	PollingFormat = "https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=%s&_=%d"
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
	// 中间数据
	uuid string
}

func (ap *AppPage) AppPageURL() string {
	return fmt.Sprintf(AppPageFormat, ap.appID, ap.redirectURL, ap.responseType, ap.scope)
}

// CallbackURL 用于生成获取 cookie 的链接,
// code 是必须的
func (ap *AppPage) CallbackURL(code string) (string, error) {
	if code == "" {
		return "", fmt.Errorf("code can not empty")
	}

	urlPref, err := url.QueryUnescape(ap.redirectURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s&code=%s&state=", urlPref, code), nil
}

// VisitAppPage 用于访问微信二维码页,
// 其返回值为二维码 id (也用作之后的 uuid),
// 同时保存该二维码
func (ap *AppPage) VisitAppPage() error {
	resp, err := http.Get(ap.AppPageURL())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	var qrcodeURLSuf string
	doc.Find(".lightBorder").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("src"); ok {
			qrcodeURLSuf = val
		}
	})
	if qrcodeURLSuf == "" {
		return fmt.Errorf("qrcode url suffix not found")
	}

	// todo: 可能要把图片转储到其他地方, 目前写到磁盘上
	if err = downloadQrcode(QrcodeURLPref + qrcodeURLSuf); err != nil {
		return err
	}

	ap.uuid = parseQrcodeID(qrcodeURLSuf)
	return nil
}

func (ap *AppPage) PollingURL() string {
	return fmt.Sprintf(PollingFormat, ap.uuid, myTime.NowUnixMilli())
}

func (ap *AppPage) Poll() <-chan string {
	code := make(chan string)
	deadline := time.NewTimer(5 * time.Minute)

	go func() {
		for {
			se
			pollingURL := ap.PollingURL()
			http.Get(pollingURL)
		}
	}()

	return code
}

//func NewPolling(uuid string) *Poller {
//	poller := &Poller{
//		uuid: uuid,
//	}
//	return poller
//}
//
//type Poller struct {
//	uuid string
//}
//
//func (pol *Poller) String() string {
//	return fmt.Sprintf(PollingFormat, pol.uuid, time.NowUnixMilli())
//}

func downloadQrcode(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	file, err := os.Create("qrcode.jpeg")
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()

	file.Write(data)
	return nil
}

// todo: 更安全的解析
func parseQrcodeID(suffix string) string {
	res := strings.Split(suffix, "/")
	if len(res) < 4 {
		return ""
	}

	return res[3]
}

// todo: 更安全的解析
func parseScanResp(data string) (res []string) {
	parts := strings.Split(data, ";")
	for _, p := range parts {
		if p == "" {
			continue
		}
		pp1 := strings.Split(p, "=")[1]
		pp1 = strings.Trim(pp1, "'")
		res = append(res, pp1)
	}
	return
}
