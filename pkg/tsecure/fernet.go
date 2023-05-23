package tsecure

import (
	"gb-auth-gate/config"
	"github.com/fernet/fernet-go"
	"github.com/sarulabs/di"
)

type FernetCrypto struct {
	EncryptionKey *fernet.Key
}

func BuildFernetEncryptor(ctn di.Container) (interface{}, error) {
	rawKey := ctn.Get("config").(*config.Config).SecureConfig.Fernet.Key
	key, err := fernet.DecodeKey(rawKey)
	if err != nil {
		return nil, err
	}
	return &FernetCrypto{
		EncryptionKey: key,
	}, nil

}

func (fe *FernetCrypto) Encrypt(message string) (cipher string, err error) {
	return fe.Encrypt(message)
}

func (fe *FernetCrypto) Decrypt(cipher string) (message string, err error) {
	return fe.Decrypt(cipher)
}
