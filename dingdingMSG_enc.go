package dingdingNotification

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//机器人加签
func dingENC(content string, webhook string, secret string) *http.Response{
	//时间戳
	timestamp := time.Now().UnixNano() / 1e6
	//fmt.Println(timestamp)

	stringToSign := fmt.Sprintf("%d\n%s",timestamp, secret)

	//使用HmacSHA256算法计算签名
	hmacCode := hmac.New(sha256.New, []byte(secret))
	hmacCode.Write([]byte(stringToSign))

	//然后进行Base64 encode
	base64Code := base64.StdEncoding.EncodeToString(hmacCode.Sum(nil))

	//最后再把签名参数再进行urlEncode
	urlCode := url.QueryEscape(base64Code)

	//构造新的webhook
	webhook_new := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, urlCode)

	//格式化消息并发送
	format := []byte(content)
	resp, _ := http.Post(webhook_new, "application/json", bytes.NewBuffer(format))
	//fmt.Println(resp.Status)
	return resp
}

//应当传入对应的 webhook 和密钥（secret）
//传入消息内容即可发送至钉钉机器人（必须符合传输格式）
func DingdingMSG_enc(webhook string, secret string, content string)*http.Response{
	return dingENC(content, webhook, secret)
}

//