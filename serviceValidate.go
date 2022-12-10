package juanjiCAS

import (
	"github.com/guonaihong/gout"
)

type CAS struct {
	domain string
}

func New(domain string) *CAS {
	return &CAS{
		domain: domain,
	}
}

func (c *CAS) SetDomain(domain string) {
	c.domain = domain
}

func (c *CAS) GetLoginURL(redirectURL string) string {
	return c.domain + "/login?service=" + redirectURL
}

func (c *CAS) GetLogoutURL(redirectURL string) string {
	return c.domain + "/logout?service=" + redirectURL
}

func (c *CAS) GetValidateURL(ticket string, redirectURL string) string {
	return c.domain + "/serviceValidate?ticket=" + ticket + "&service=" + redirectURL
}

// getProxyValidateURL
// NOT SUPPORT YET
// deprecated
func (c *CAS) getProxyValidateURL(ticket string, redirectURL string) string {
	return c.domain + "/proxyValidate?ticket=" + ticket + "&service=" + redirectURL
}

// Validate redirectURL is the service url
func (c *CAS) Validate(ticket string, redirectURL string) (CASResponse, error) {
	casResponse := CASResponse{}
	err := gout.GET(c.GetValidateURL(ticket, redirectURL) + "&format=JSON").BindJSON(&casResponse).Do()
	return casResponse, err
}
