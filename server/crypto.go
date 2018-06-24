package server

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	r "math/rand"
	"sort"
	"strings"
	"time"
)

type DingTalkCrypto struct {
	Token          string
	EncodingAESKey string
	SuiteKey       string
	BKey           []byte
	Block          cipher.Block
}

func NewDingTalkCrypto(token, encodingAESKey, suiteKey string) *DingTalkCrypto {
	if len(encodingAESKey) != aesEncodeKeyLen {
		panic("不合法的EncodingAESKey")
	}
	bkey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher(bkey)
	if err != nil {
		panic(err.Error())
	}
	c := &DingTalkCrypto{
		Token:          token,
		EncodingAESKey: encodingAESKey,
		SuiteKey:       suiteKey,
		BKey:           bkey,
		Block:          block,
	}
	return c
}

func (c *DingTalkCrypto) GetDecryptMsg(signture, timestamp, nonce, secretMsg string) (string, error) {
	if !c.VerificationSignature(c.Token, timestamp, nonce, secretMsg, signture) {
		return "", errors.New("ERROR: 签名不匹配")
	}
	decode, err := base64.StdEncoding.DecodeString(secretMsg)
	if err != nil {
		return "", err
	}
	if len(decode) < aes.BlockSize {
		return "", errors.New("ERROR: 密文太短")
	}
	blockMode := cipher.NewCBCDecrypter(c.Block, c.BKey[:c.Block.BlockSize()])
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = pkCS7UnPadding(plantText)
	size := binary.BigEndian.Uint32(plantText[16:20])
	plantText = plantText[20:]
	corpID := plantText[size:]
	if string(corpID) != c.SuiteKey {
		return "", errors.New("ERROR: CorpID匹配不正确")
	}
	return string(plantText[:size]), nil
}

func (c *DingTalkCrypto) GetEncryptMsg(msg, timestamp, nonce string) (string, string, error) {
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(msg)))
	msg = randomString(16) + string(size) + msg + c.SuiteKey
	plantText := pkCS7Padding([]byte(msg), c.Block.BlockSize())
	if len(plantText)%aes.BlockSize != 0 {
		return "", "", errors.New("ERROR: 消息体size不为16倍数")
	}
	blockMode := cipher.NewCBCEncrypter(c.Block, c.BKey[:c.Block.BlockSize()])
	cipherText := make([]byte, len(plantText))
	blockMode.CryptBlocks(cipherText, plantText)
	outMsg := base64.StdEncoding.EncodeToString(cipherText)
	signture := c.CreateSignture(c.Token, timestamp, nonce, string(outMsg))
	return string(outMsg), signture, nil
}

// 数据签名
func (c *DingTalkCrypto) CreateSignture(token, timestamp, nonce, msg string) string {
	params := make([]string, 0)
	params = append(params, token)
	params = append(params, timestamp)
	params = append(params, nonce)
	params = append(params, msg)
	sort.Strings(params)
	return sha1Sign(strings.Join(params, ""))
}

// 验证数据签名
func (c *DingTalkCrypto) VerificationSignature(token, timestamp, nonce, msg, signture string) bool {
	return c.CreateSignture(token, timestamp, nonce, msg) == signture
}

// 解密补位
func pkCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// 加密补位
func pkCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 随机字符串
func randomString(n int, alphabets ...byte) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	data := make([]byte, n)
	randby := false
	if num, err := rand.Read(data); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range data {
		if len(alphabets) == 0 {
			if randby {
				data[i] = alphanum[r.Intn(len(alphanum))]
			} else {
				data[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				data[i] = alphabets[r.Intn(len(alphabets))]
			} else {
				data[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return string(data)
}
