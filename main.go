package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// 58pic.com 微信登录二维码
// https://open.weixin.qq.com/connect/qrconnect?appid=wx47e0850600dd30fc&redirect_uri=https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin&response_type=code&scope=snsapi_login

// 目标: 取出验证码
// 流程:
//     1. 访问 58pic.com 微信扫码页
//     2. 获取二维码
//     3. 轮询手机扫码状态
//     4. 扫码后的链接, 可不访问
//     5. 千图网回调
//     6. 重定向

// 其中 3, 4步骤还在访问, 不确定票在哪个响应里

var AppID58pic = "https://open.weixin.qq.com/connect/qrconnect?appid=wx47e0850600dd30fc&redirect_uri=https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin&response_type=code&scope=snsapi_login"
var QrcodeURLPref = "https://open.weixin.qq.com"

// https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=011KWtIYwdjR0lRi&_=1572590279242
var LoopURLNotScan = "https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=%s&_=%d"

// https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=001hjUNi2kHhBwC4&last=404&_=1572593319534
// todo: 404 可能需要动态获取
var LoopURLScan = "https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=%s&last=404&_=%d"

// https://www.58pic.com/index.php?m=login&a=callback&type=weixin&code=061ia8Fp12swNi0jJqFp1Ol5Fp1ia8Fh&state=
var CallBackURLL = "https://www.58pic.com/index.php?m=login&a=callback&type=weixin&code=%s&state="

// 重定向
// <script>window.top.document.location.href='https://accounts.58pic.com/index.php?m=Ssl_login&ticket=071KxpKn0f1gfl1jcHLn0lApKn0KxpKJ&callback_url=Ly93d3cuNThwaWMuY29tLw==';</script>

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", AppID58pic, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var qrcodeURLSuf string

	doc.Find(".lightBorder").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("src"); ok {
			qrcodeURLSuf = val
			fmt.Println(val)
		}
	})

	// todo: 验证正确性
	downloadQrcode(QrcodeURLPref + qrcodeURLSuf)

	id := qrcodeID(qrcodeURLSuf)

	ticket := ""
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		loopURL := fmt.Sprintf(LoopURLNotScan, id, time.Now().UnixNano()/1e6)
		req, err = http.NewRequest("GET", loopURL, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Cookies()

		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp.Body.Close()

		dataStr := string(data)
		res := parseResp(dataStr)
		// 之前一直是408
		if res[0] == "405" {
			ticket = res[1]
			fmt.Println("票: ", res[1])
			break
		}
	}

	if ticket == "" {
		fmt.Println("get ticket fail")
		return
	}

	callBackURL := fmt.Sprintf(CallBackURLL, ticket)
	req, err = http.NewRequest("GET", callBackURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()

	fmt.Printf("%s\n", data)
	// cookie 在这
	fmt.Println("cookie2:", resp.Cookies())

	realURL := parseRealURL(string(data))

	resp, err = http.Get(realURL)

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()

	fmt.Println("cookie1:", resp.Cookies())
}

func qrcodeID(suffix string) string {
	res := strings.Split(suffix, "/")
	if len(res) < 4 {
		return ""
	}

	return res[3]
}

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

// wx_errcode=408;window.wx_code='';
func parseResp(data string) (res []string) {
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

func parseRealURL(data string) string {
	res := strings.Split(data, "'")
	if len(res) < 3 {
		return ""
	}
	return res[1]
}

// todo: 解析 cookies
func parseCookies() {

}
