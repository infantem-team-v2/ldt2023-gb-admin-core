package tsecure

// IEncryptor Basic interface for structures which encrypt and decrypt messages
type IEncryptor interface {
	Encrypt(message string) (cipher string, err error)
	Decrypt(cipher string) (message string, err error)
}
