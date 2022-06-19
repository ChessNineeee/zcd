package client

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
)

type TLSClient struct {
	certPath  string
	keyPath   string
	caPath    string
	err       error
	cer       tls.Certificate
	certPool  *x509.CertPool
	tlsConfig *tls.Config
}

func NewTLSClient(certPath, keyPath, caPath string) *TLSClient {
	return &TLSClient{
		certPath:  certPath,
		keyPath:   keyPath,
		caPath:    caPath,
		certPool:  x509.NewCertPool(),
		tlsConfig: &tls.Config{},
	}
}

func (t *TLSClient) loadKeyPair() *TLSClient {
	if t.err == nil {
		t.cer, t.err = tls.LoadX509KeyPair(t.certPath, t.keyPath)
	}
	return t
}

func (t *TLSClient) loadCA() *TLSClient {
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

func (t *TLSClient) initTLSConfig() *TLSClient {
	if t.err == nil {
		t.tlsConfig.Certificates = []tls.Certificate{t.cer}
		t.tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		t.tlsConfig.ClientCAs = t.certPool
	}
	return t
}

func (t *TLSClient) Dial(address string) (*tls.Conn, error) {
	return tls.Dial("tcp", address, t.tlsConfig)
}
