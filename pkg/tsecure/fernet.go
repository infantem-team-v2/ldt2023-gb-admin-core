package tsecure

import (
	"gb-admin-core/config"
	"github.com/fernet/fernet-go"
	"github.com/sarulabs/di"
	"time"
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
	tok, err := fernet.EncryptAndSign([]byte(message), fe.EncryptionKey)
	return string(tok), err
}

func (fe *FernetCrypto) Decrypt(cipher string) (message string, err error) {
	message = string(fernet.VerifyAndDecrypt([]byte(cipher), 0*time.Second, []*fernet.Key{fe.EncryptionKey}))

	return message, nil
}
