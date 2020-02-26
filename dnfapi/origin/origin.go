package origin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//allowOrigins 保存允许跨域的地址
var allowOrigins = []string{}

func init() {

	allowOrigins = []string{}
	f, err := os.Open("origin.conf")
	if err != nil {
		defer fmt.Println("打开配置文件origin.conf失败")
		return
	}
	defer f.Close()
	defer fmt.Println("origin模块加载成功")
	freader := bufio.NewReader(f)
	for {
		line, err := freader.ReadBytes('\n')
		newOrigin := strings.Trim(string(line), "\n")
		newOrigin = strings.Trim(newOrigin, "\r")
		allowOrigins = append(allowOrigins, newOrigin)
		if err != nil {
			return
		}
	}

}

//IsAllowOrigin 判断是否在允许的域内访问
func IsAllowOrigin(ip string) bool {
	for _, v := range allowOrigins {
		if v == ip {

			return true
		}

	}
	return false
}
