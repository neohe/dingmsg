package main

import (
	"fmt"
	"os"

	"github.com/hugozhu/godingtalk"
)

func main() {
	c := godingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	c.RefreshAccessToken()
	err := c.SendAppMessage(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
	}
}
