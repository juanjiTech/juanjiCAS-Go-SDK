package juanjiCAS_Go_SDK

import "testing"

func TestExample(t *testing.T) {
	cas := New("https://cas.juanji.tech", "secret")
	cas.Validate("ST-1-1-1-1", "https://example.com")
}
