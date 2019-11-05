package modules

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestURLEncode(t *testing.T) {
	res, err := url.PathUnescape("https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res, err = url.QueryUnescape("https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func TestAppPage_String(t *testing.T) {
	decode, err := NewAppPage("1234", "https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		panic(err)
	}

	notDecode, err := NewAppPage("1234", "https://www.58pic.com/index.php?m=login&a=callback&type=weixin")
	if err != nil {
		panic(err)
	}

	if decode.AppPageURL() == notDecode.AppPageURL() {
		fmt.Println("equal")
	}
}

func TestAppPage_CallbackURL(t *testing.T) {
	appPage, err := NewAppPage("123", "https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		fmt.Println(err)
		return
	}

	callbackURL, err := appPage.CallbackURL("456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("callbackURL:", callbackURL)
}

//func TestPoller_String(t *testing.T) {
//	poller := NewPolling("123")
//	fmt.Println(poller)
//}

func TestReadClosedChan(t *testing.T) {
	msg := make(chan string)
	go func() {
		msg <- "123"
		close(msg)
	}()

	time.Sleep(2 * time.Second)
	if val, ok := <-msg; ok {
		fmt.Println(val)
	} else {
		fmt.Println("read val fail")
	}
}

func TestSaveCookies(t *testing.T) {
	http.Response{}.Cookies()
}
