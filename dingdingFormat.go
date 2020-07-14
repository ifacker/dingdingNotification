package dingdingNotification

import (
	"fmt"
	"strconv"
)

//钉钉消息之 text 格式化，
//传入要发送的内容（content）
//还有是否@所有人（isAtAll）
func DingMSG_text(content string, isAtAll bool) string {
	sendText := `{
    "msgtype": "text", 
    "text": {
        "content": "%s"
    }, 
    "at": {
        "isAtAll": %s
    }
}`
	sendText = fmt.Sprintf(sendText, content, strconv.FormatBool(isAtAll))
	return sendText
}

//钉钉消息之 markdown 格式化
//传入标题（title）
//传入要发送的内容（content）
//还有是否@所有人（isAtAll）
func DingMSG_markdown(title string, content string, isAtAll bool) string {
	sendMarkdown := `{"msgtype":"markdown",
     "markdown": {
         "title":"%s",
         "text": "%s"
     },
	"at":{
          "isAtAll": %s
      }
	}`
	sendMarkdown = fmt.Sprintf(sendMarkdown, title, content, strconv.FormatBool(isAtAll))
	return sendMarkdown
}

//开发者们还可以更具自己的需要往里面添加自己需要的格式