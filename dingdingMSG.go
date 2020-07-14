package dingdingNotification

import (
	"bytes"
	"fmt"
	"net/http"
)

//发送消息到钉钉，需要加入关键字才可以正常发送
func ding(webhook string, data string)*http.Response{
	format := []byte(data)

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(format))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(*resp)
	return resp
}

//传如 webhook 和 要发送的数据（data）进行发送
//注意：需要加入关键字或者IP地址加入了钉钉机器人
//的白名单之后才可以正常发送哦！
func DingdingMSG(webhook string, data string)*http.Response{
	return ding(webhook,data)
}