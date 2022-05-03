package dingtalk_test

import (
	"fmt"
	"testing"

	"github.com/kevin2027/easy-dingtalk/message"
)

func TestMessageCorpconversationaSyncsendV2(t *testing.T) {
	var err error
	defer deferErr(&err)
	msg := &message.MessageRequest{
		Msgtype: "text",
		Text: &message.TextMessage{
			Content: "这是一段文本消息",
		},
	}
	taskId, err := client.Message().CorpconversationaSyncsendV2([]string{"user0"}, nil, false, msg)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", taskId)
}

func TestMessageCorpconversationaSyncsendV2Markdown(t *testing.T) {
	var err error
	defer deferErr(&err)
	msg := &message.MessageRequest{
		Msgtype: "markdown",
		Markdown: &message.MarkdownMessage{
			Text:  "这是一段文本消息",
			Title: "标题",
		},
	}
	taskId, err := client.Message().CorpconversationaSyncsendV2([]string{"user0"}, nil, false, msg)
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	fmt.Printf("%v\n", taskId)
}
