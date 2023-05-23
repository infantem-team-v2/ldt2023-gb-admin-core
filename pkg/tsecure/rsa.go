package tsecure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"gb-auth-gate/config"
	"github.com/sarulabs/di"
	"os"
)

type RsaCrypto struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func openRsaFile(path string) (key []byte, err error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pemFile, err := os.Open(fmt.Sprintf("%s%s", workDir, path))
	if err != nil {
		return nil, err
	}
	defer pemFile.Close()
	_, err = pemFile.Read(key)
	return key, err
}

func parseRsaPublicKey(key []byte) (publicKey *rsa.PublicKey, err error) {
	return x509.ParsePKCS1PublicKey(key)
}

func parseRsaPrivateKey(key []byte) (privateKey *rsa.PrivateKey, err error) {
	return x509.ParsePKCS1PrivateKey(key)
}

func BuildRsaCrypto(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	publicKeyRaw, err := openRsaFile(cfg.SecureConfig.RSA.PublicKey)
	if err != nil {
		return nil, err
	}
	publicKey, err := parseRsaPublicKey(publicKeyRaw)
	if err != nil {
		return nil, err
	}
	privateKeyRaw, err := openRsaFile(cfg.SecureConfig.RSA.PrivateKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := parseRsaPrivateKey(privateKeyRaw)
	if err != nil {
		return nil, err
	}

	return &RsaCrypto{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

func (r *RsaCrypto) Encrypt(message string) (cipher string, err error) {
	cipherRaw, err := rsa.EncryptOAEP(hashesMap[SHA256](), rand.Reader, r.publicKey, []byte(message), nil)
	return string(cipherRaw), err
}

func (r *RsaCrypto) Decrypt(cipher string) (message string, err error) {
	messageRaw, err := rsa.DecryptOAEP(hashesMap[SHA256](), rand.Reader, r.privateKey, []byte(cipher), nil)
	return string(messageRaw), err
}
