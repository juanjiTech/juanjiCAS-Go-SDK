package juanjiCAS

import (
	"encoding/xml"
)

type CASResponse struct {
	CASServiceResponse `json:"serviceResponse"`
}

type CASServiceResponse struct {
	XMLName      xml.Name                  `xml:"cas:serviceResponse" json:"-"`
	Xmlns        string                    `xml:"xmlns:cas,attr" json:"-"`
	Success      *CASAuthenticationSuccess `json:"authenticationSuccess,omitempty"`
	Failure      *CASAuthenticationFailure `json:"authenticationFailure,omitempty"`
	ProxySuccess string                    `xml:"cas:proxyTicket,omitempty" json:"proxyTicket,omitempty"`
	ProxyFailure *CASProxyFailure          `json:"proxyFailure,omitempty"`
}

type CASAuthenticationSuccess struct {
	XMLName    xml.Name      `xml:"cas:authenticationSuccess" json:"-"`
	User       string        `xml:"cas:user" json:"user"`
	Attributes CASAttributes `xml:"cas:attributes" json:"attributes"`
}

type CASAuthenticationFailure struct {
	XMLName xml.Name `xml:"cas:authenticationFailure" json:"-"`
	Code    string   `xml:"code,attr" json:"code"`
	Desc    string   `xml:",chardata" json:"description"`
}

type CASPgtIou struct {
	XMLName xml.Name `xml:"cas:proxyGrantingTicket" json:"-"`
	Ticket  string   `xml:",chardata" json:"ticket"`
}

type CASAttributes struct {
	XMLName xml.Name `xml:"cas:attributes" json:"-"`
	Email   string   `xml:"email" json:"email"`
}

type CASProxyFailure struct {
	XMLName xml.Name `xml:"cas:proxyFailure" json:"-"`
	Code    string   `xml:"string" json:"code"`
	Desc    string   `xml:",chardata" json:"description"`
}
