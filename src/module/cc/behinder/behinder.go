package behinder

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type GenerateMode uint32

const (
	ModeFile GenerateMode = iota
	ModeCode
	ModeCommand
)

type IBehinder interface {
	Generate() []byte
	Execute(argv, envp string)
	Upload(src, dst string)
	Download(src, dst string)
}

type Behinder struct {
	SessionKey []byte
	RandomIV   []byte
}

// pkcs7 padding
func (cc *Behinder) pad(data []byte) []byte {
	padSize := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, padding...)
}

func (cc *Behinder) unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

func (cc *Behinder) encrypt(plainData []byte) (cipherData []byte) {
	block, err := aes.NewCipher(cc.SessionKey)
	if err != nil {
		return nil
	}
	plainPadding := cc.pad(plainData)
	cipherData = make([]byte, len(plainPadding))
	blockMode := cipher.NewCBCEncrypter(block, cc.RandomIV)
	blockMode.CryptBlocks(cipherData, plainPadding)
	base64.StdEncoding.Encode(cipherData, cipherData)
	return cipherData
}

func (cc *Behinder) decrypt(cipherData []byte) (plainData []byte) {
	cipherData, err := base64.StdEncoding.DecodeString(string(cipherData))
	if err != nil {
		return nil
	}
	block, err := aes.NewCipher(cc.SessionKey)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, cc.RandomIV)
	plainData = make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	plainData = cc.unpad(plainData)
	return plainData
}
