package message

type MessageRequest struct {
	Msgtype    string             `json:"msgtype"`
	Text       *TextMessage       `json:"text,omitempty"`
	Image      *FileMessage       `json:"image,omitempty"`
	Voice      *VoiceMessage      `json:"voice,omitempty"`
	File       *FileMessage       `json:"file,omitempty"`
	Link       *LinkMessage       `json:"link,omitempty"`
	Oa         *OaMessage         `json:"oa,omitempty"`
	Markdown   *MarkdownMessage   `json:"markdown,omitempty"`
	ActionCard *ActionCardMessage `json:"action_card,omitempty"`
}

func (m *MessageRequest) Clone() *MessageRequest {
	result := &MessageRequest{
		Msgtype: m.Msgtype,
	}
	switch m.Msgtype {
	case "text":
		result.Text = m.Text
	case "image":
		result.Image = m.Image
	case "voice":
		result.Voice = m.Voice
	case "file":
		result.File = m.File
	case "link":
		result.Link = m.Link
	case "oa":
		result.Oa = m.Oa
	case "markdown":
		result.Markdown = m.Markdown
	case "action_card":
		result.ActionCard = m.ActionCard
	}
	return result
}

type TextMessage struct {
	Content string `json:"content"`
}

type FileMessage struct {
	MediaId string `json:"media_id"`
}

type VoiceMessage struct {
	FileMessage
	Duration string `json:"duration"`
}

type LinkMessage struct {
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

type OaMessage struct {
	MessageUrl   string              `json:"message_url"`
	PcMessageUrl *string             `json:"pc_message_url,omitempty"`
	Head         OaMessageHead       `json:"head"`
	StatusBar    *OaMessageStatusBar `json:"status_bar,omitempty"`
	Body         OaMessageBody       `json:"body"`
}

type OaMessageHead struct {
	Bgcolor string `json:"bgcolor"`
	Text    string `json:"text"`
}

type OaMessageStatusBar struct {
	StatusValue *string `json:"status_value,omitempty"`
	StatusBg    *string `json:"status_bg,omitempty"`
}

type OaMessageBody struct {
	Title     *string          `json:"title,omitempty"`
	Form      []*OaMessageForm `json:"form,omitempty"`
	Rich      *OaMessageRich   `json:"rich,omitempty"`
	Content   *string          `json:"content,omitempty"`
	Image     *string          `json:"image,omitempty"`
	FileCount *string          `json:"file_count,omitempty"`
	Author    *string          `json:"author,omitempty"`
}

type OaMessageForm struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type OaMessageRich struct {
	Num  *string `json:"num,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

type MarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCardMessage struct {
	Markdown       string                  `json:"markdown"`
	Title          *string                 `json:"title,omitempty"`
	SingleTitle    *string                 `json:"single_title,omitempty"`
	SingleUrl      *string                 `json:"single_url,omitempty"`
	BtnOrientation *string                 `json:"btn_orientation,omitempty"`
	BtnJsonList    []*ActionCardMessageBtn `json:"btn_json_list"`
}
type ActionCardMessageBtn struct {
	Title     *string `json:"title,omitempty"`
	ActionUrl *string `json:"action_url,omitempty"`
}
