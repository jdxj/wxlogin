package logs

import (
	"fmt"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	res := parseResp("wx_errcode=408;window.wx_code='asdfas';")
	fmt.Println(res)

}
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

func TestE(t *testing.T) {
	fmt.Printf("%f", 1e12)
}

func TestLc(t *testing.T) {
	fmt.Println(1000 / 10 / 10)
}
