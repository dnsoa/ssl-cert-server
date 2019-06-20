package gocert

import (
	"crypto/tls"
	"errors"
	"strings"
)

type provider struct {
	serverURL string
}

func NewProvider(serverURL string) *provider {
	return &provider{
		serverURL: serverURL,
	}
}

func (p *provider) TLSConfig() *tls.Config {

}

func (p *provider) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	name := hello.ServerName
	if name == "" {
		return nil, errors.New("gocert: missing server name")
	}
	if !strings.Contains(strings.Trim(name, "."), ".") {
		return nil, errors.New("gocert: server name component count invalid")
	}
	if strings.ContainsAny(name, `+/\`) {
		return nil, errors.New("gocert: server name contains invalid character")
	}
	name = strings.TrimSuffix(name, ".")

}

func (p *provider) getCertificate(domain string) {

}
