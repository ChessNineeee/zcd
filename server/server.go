package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"net"
)

type TLSRouter struct {
	certPath  string
	keyPath   string
	caPath    string
	address   string
	err       error
	cer       tls.Certificate
	certPool  *x509.CertPool
	tlsConfig *tls.Config
	listener  net.Listener
}

func NewTLSRouter(certPath, keyPath, caPath, address string) *TLSRouter {
	return &TLSRouter{
		certPath:  certPath,
		keyPath:   keyPath,
		caPath:    caPath,
		address:   address,
		certPool:  x509.NewCertPool(),
		tlsConfig: &tls.Config{},
	}
}

func (t *TLSRouter) loadKeyPair() *TLSRouter {
	if t.err == nil {
		t.cer, t.err = tls.LoadX509KeyPair(t.certPath, t.keyPath)
	}
	return t
}

func (t *TLSRouter) loadCA() *TLSRouter {
	if t.err == nil {
		ca, err := ioutil.ReadFile(t.caPath)
		if err != nil {
			t.err = err
			log.Println("加载CA证书出错", err.Error())
			return t
		}
		if ok := t.certPool.AppendCertsFromPEM(ca); !ok {
			t.err = errors.New("添加根证书出错")
		}
	}
	return t
}

func (t *TLSRouter) initTLSConfig() *TLSRouter {
	if t.err == nil {
		t.tlsConfig.Certificates = []tls.Certificate{t.cer}
		t.tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		t.tlsConfig.ClientCAs = t.certPool
	}
	return t
}

func (t *TLSRouter) listen() *TLSRouter {
	if t.err == nil {
		t.listener, t.err = tls.Listen("tcp", t.address, t.tlsConfig)
	}
	return t
}
