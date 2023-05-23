package tsecure

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type HashGroup uint16

const (
	SHA256 HashGroup = iota
	SHA512
	SHA1
	MD5
)

var hashesMap = map[HashGroup]func() hash.Hash{
	SHA1:   sha1.New,
	SHA256: sha256.New,
	SHA512: sha512.New,
	MD5:    md5.New,
}

// CalcSignature Method that build signature w/ algorithm from constants which put in hashesMap
func CalcSignature(secret, message string, hashType HashGroup) string {
	mac := hmac.New(hashesMap[hashType], []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
