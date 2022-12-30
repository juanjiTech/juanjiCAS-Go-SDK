package juanjiCAS

type Notice struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`  // 详情链接
	To    string `json:"to"`   // UID 仅当推送为服务推送时该字段生效
	From  string `json:"from"` // 服务名 仅当推送为个人推送时该字段生效
}

func NewNotice(title, body, url string) *Notice {
	return &Notice{Title: title, Body: body, Url: url}
}

// SetTo 设置推送目标 仅当推送为服务推送时该设置有用
func (n *Notice) SetTo(to string) *Notice {
	n.To = to
	return n
}

// SetFrom 设置推送来源 仅当推送为个人推送时该设置有用
func (n *Notice) SetFrom(from string) *Notice {
	n.From = from
	return n
}
