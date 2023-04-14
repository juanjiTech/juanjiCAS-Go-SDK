package juanjiCAS

import (
	"testing"
)

func TestExample(t *testing.T) {
	cas := New("https://cas.juanji.tech")
	_ = cas.GetLoginURL("https://example.com")
	_, _ = cas.Validate("ST-1-1-1-1", "https://example.com")

	cas.SetSendKey("123456")
	_ = cas.Notify(NewNotice("title", "body", "https://example.com").SetFrom("example")) // 个人推送 可设置推送来源
	_ = cas.Notify(NewNotice("title", "body", "https://example.com").SetTo("1"))         // 服务推送 可设置推送目标
}
