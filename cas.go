package juanjiCAS

type CAS struct {
	domain  string
	sendKey string
}

func New(domain ...string) *CAS {
	if len(domain) == 0 {
		return &CAS{domain: "https://cas.juanji.tech"}
	}
	return &CAS{domain: domain[0]}
}

func (c *CAS) SetDomain(domain string) *CAS {
	c.domain = domain
	return c
}

func (c *CAS) SetSendKey(sendKey string) *CAS {
	c.sendKey = sendKey
	return c
}
