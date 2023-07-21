package juanjiCAS

import "net/url"

type CAS struct {
	frontend *url.URL
	backend  *url.URL
	sendKey  string
}

func New(domain ...string) *CAS {
	c := &CAS{}
	if len(domain) == 0 {
		return c.SetDomain("https://api.cas.juanji.tech").
			SetFrontendDomain("https://cas.juanji.tech")
	}
	return c.SetDomain(domain[0])
}

// SetDomain set backend domain
func (c *CAS) SetDomain(domain string) *CAS {
	u, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}
	c.backend = u
	return c
}

// SetFrontendDomain set frontend domain
func (c *CAS) SetFrontendDomain(domain string) *CAS {
	u, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}
	c.frontend = u
	return c
}

func (c *CAS) SetSendKey(sendKey string) *CAS {
	c.sendKey = sendKey
	return c
}
