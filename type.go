package juanjiCAS

import (
	"encoding/xml"
	"time"
)

type CasServiceResponse struct {
	XMLName      xml.Name `xml:"cas:serviceResponse" json:"-"`
	Xmlns        string   `xml:"xmlns:cas,attr"`
	Failure      *CasAuthenticationFailure
	Success      *CasAuthenticationSuccess
	ProxySuccess *CasProxySuccess
	ProxyFailure *CasProxyFailure
}

type CasAuthenticationFailure struct {
	XMLName xml.Name `xml:"cas:authenticationFailure" json:"-"`
	Code    string   `xml:"code,attr"`
	Message string   `xml:",innerxml"`
}

type CasAuthenticationSuccess struct {
	XMLName             xml.Name           `xml:"cas:authenticationSuccess" json:"-"`
	User                string             `xml:"cas:user"`
	ProxyGrantingTicket string             `xml:"cas:proxyGrantingTicket,omitempty"`
	Proxies             *CasProxies        `xml:"cas:proxies"`
	Attributes          *CasAttributes     `xml:"cas:attributes"`
	ExtraAttributes     []*CasAnyAttribute `xml:",any"`
}

type CasProxies struct {
	XMLName xml.Name `xml:"cas:proxies" json:"-"`
	Proxies []string `xml:"cas:proxy"`
}

type CasAttributes struct {
	XMLName                                xml.Name  `xml:"cas:attributes" json:"-"`
	AuthenticationDate                     time.Time `xml:"cas:authenticationDate"`
	LongTermAuthenticationRequestTokenUsed bool      `xml:"cas:longTermAuthenticationRequestTokenUsed"`
	IsFromNewLogin                         bool      `xml:"cas:isFromNewLogin"`
	MemberOf                               []string  `xml:"cas:memberOf"`
	UserAttributes                         *CasUserAttributes
	ExtraAttributes                        []*CasAnyAttribute `xml:",any"`
}

type CasUserAttributes struct {
	XMLName       xml.Name             `xml:"cas:userAttributes" json:"-"`
	Attributes    []*CasNamedAttribute `xml:"cas:attribute"`
	AnyAttributes []*CasAnyAttribute   `xml:",any"`
}

type CasNamedAttribute struct {
	XMLName xml.Name `xml:"cas:attribute" json:"-"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",innerxml"`
}

type CasAnyAttribute struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type CasProxySuccess struct {
	XMLName     xml.Name `xml:"cas:proxySuccess" json:"-"`
	ProxyTicket string   `xml:"cas:proxyTicket"`
}
type CasProxyFailure struct {
	XMLName xml.Name `xml:"cas:proxyFailure" json:"-"`
	Code    string   `xml:"code,attr"`
	Message string   `xml:",innerxml"`
}
