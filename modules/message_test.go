package modules

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURLEncode(t *testing.T)  {
	res, err := url.PathUnescape("https%3A%2F%2Fwww.58pic.com%2Findex.php%3Fm%3Dlogin%26a%3Dcallback%26type%3Dweixin")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
