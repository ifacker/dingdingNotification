package dingdingNotification

import (
	"fmt"
	"testing"
)

var WEB_HOOK = "xxxx"
var SECRET = "xxx"


func TestDingdingMSG_enc(t *testing.T) {

	//text 格式发送
	send_text := DingMSG_text("hello world", false)
	resp := DingdingMSG_enc(WEB_HOOK, SECRET, send_text)
	fmt.Println(resp.Status)

	//markdown 格式发送
	content := fmt.Sprintf("# hello \n> #### hello world")
	send_markdown := DingMSG_markdown("hello", content, true)
	resp1 := DingdingMSG_enc(WEB_HOOK, SECRET, send_markdown)
	fmt.Println(resp1.Status)
}

func TestDingdingMSG(t *testing.T) {

	//text 格式发送，注意关键词
	send_text := DingMSG_text("关键词：hello world", true)
	resp := DingdingMSG(WEB_HOOK, send_text)
	fmt.Println(resp.Status)

	//markdown 格式发送，注意关键词，关键词放到标题或内容里都可以
	content := fmt.Sprintf("# hello \n> #### 关键词：hello world")
	send_markdown := DingMSG_markdown("hello", content, false)
	resp1 := DingdingMSG(WEB_HOOK, send_markdown)
	fmt.Println(resp1.Status)
}