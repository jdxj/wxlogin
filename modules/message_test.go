package modules

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURLEncode(t *testing.T) {
	res, err := url.PathUnescape("https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res, err = url.QueryUnescape("https://www.58pic.com/index.php?m=login&a=callback&type=weixin")
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

	if decode.String() == notDecode.String() {
		fmt.Println("equal")
	}
}

func TestPoller_String(t *testing.T) {
	poller := NewPolling("123")
	fmt.Println(poller)
}
