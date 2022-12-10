package juanjiCAS

import "testing"

func TestExample(t *testing.T) {
	cas := New("https://cas.juanji.tech")
	cas.Validate("ST-1-1-1-1", "https://example.com")
}
