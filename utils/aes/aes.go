package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/astaxie/beego"

	"github.com/lzxz1234/AdminBoot/utils/base58"
)

var aeskey = []byte(beego.AppConfig.String("boot.aesKey"))

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// EncryptString with default key
func EncryptString(origData string) string {

	result, err := Encrypt([]byte(origData))
	if err != nil {
		return ""
	}
	return base58.Encode(result)
}

// DecryptString with default key
func DecryptString(crypted string) (string, error) {

	cryptBytes := base58.Decode(crypted)
	result, err := Decrypt(cryptBytes)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// Encrypt with default key
func Encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = _PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aeskey[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// Decrypt with default key
func Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, aeskey[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = _PKCS5UnPadding(origData)
	return origData, nil
}
