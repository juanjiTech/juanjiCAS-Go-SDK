package juanjiCAS

import (
	"github.com/guonaihong/gout"
)

func (c *CAS) GetLoginURL(redirectURL string) string {
	u := c.domain.JoinPath("/login")
	q := u.Query()
	q.Set("service", redirectURL)
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *CAS) GetLogoutURL(redirectURL string) string {
	u := c.domain.JoinPath("/logout")
	q := u.Query()
	q.Set("service", redirectURL)
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *CAS) GetValidateURL(ticket string, redirectURL string) string {
	u := c.domain.JoinPath("/serviceValidate")
	q := u.Query()
	q.Set("ticket", ticket)
	q.Set("service", redirectURL)
	u.RawQuery = q.Encode()
	return u.String()
}

// getProxyValidateURL
// NOT SUPPORT YET
// deprecated
func (c *CAS) getProxyValidateURL(ticket string, redirectURL string) string {
	u := c.domain.JoinPath("/proxyValidate")
	q := u.Query()
	q.Set("ticket", ticket)
	q.Set("service", redirectURL)
	u.RawQuery = q.Encode()
	return u.String()
}

// Validate redirectURL is the service url
func (c *CAS) Validate(ticket string, redirectURL string) (CasServiceResponse, error) {
	casResponse := CasServiceResponse{}
	err := gout.GET(c.GetValidateURL(ticket, redirectURL) + "&format=JSON").BindJSON(&casResponse).Do()
	return casResponse, err
}
