package juanjiCAS

import "net/url"

type CAS struct {
	domain  *url.URL
	sendKey string
}

func New(domain ...string) *CAS {
	c := &CAS{}
	if len(domain) == 0 {
		return c.SetDomain("https://cas.juanji.tech")
	}
	return c.SetDomain(domain[0])
}

func (c *CAS) SetDomain(domain string) *CAS {
	u, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}
	c.domain = u
	return c
}

func (c *CAS) SetSendKey(sendKey string) *CAS {
	c.sendKey = sendKey
	return c
}
