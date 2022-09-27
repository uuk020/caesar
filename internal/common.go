package internal

import (
	"bytes"
	"caesar/global"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetNowFormatTodayTime() string {
	cstSh, _ := time.LoadLocation(global.Setting.Timezone)
	now := time.Now().In(cstSh)
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func GetNowTimestamp() int64 {
	cstSh, _ := time.LoadLocation(global.Setting.Timezone)
	now := time.Now().In(cstSh)
	return now.Unix()
}

func RandomCode() string {
	var s string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		r := rand.Intn(9)
		s += strconv.Itoa(r)
	}
	return s
}

func SendMail(to []string, subject string, msg string) bool {
	emailConfig := global.Setting.Email
	auth := smtp.PlainAuth("", emailConfig.Name, emailConfig.Password, emailConfig.Host)
	smtpAddr := fmt.Sprintf("%s:%v", emailConfig.Host, emailConfig.Port)
	err := smtp.SendMail(smtpAddr, auth, emailConfig.Name, to, []byte("From: 凯撒密码网管理员<jie893609357@163.com>\r\n"+
		"Subject:"+subject+"\r\n"+
		"\r\n"+msg))
	return err != nil
}

func AesEncrypt(orig string, key string, keySize int) (string, error) {
	origData := []byte(orig)

	switch keySize {
	default:
		return "", errors.New("keySize 参数错误")
	case 16, 24, 32:
		break
	}

	k := keyPadding([]byte(key), keySize)

	block, err := aes.NewCipher(k)

	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted), nil
}

func AesDecrypt(cryted string, key string, keySize int) (string, error) {
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", err
	}

	switch keySize {
	default:
		return "", errors.New("keySize 参数错误")
	case 16, 24, 32:
		break
	}

	k := keyPadding([]byte(key), keySize)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

func keyPadding(key []byte, keySize int) []byte {
	padding := keySize - len(key)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(key, padText...)
}

// 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// logger 配置
func NewLoggerConfig() middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogMethod:    true,
		LogUserAgent: true,
		LogError:     true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status != http.StatusOK {
				file := fmt.Sprintf("%s%s.log", global.Setting.LogAddr, GetNowFormatTodayTime())
				fd, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return err
				}
				s := "Request: time:" + v.StartTime.Local().String() + ", uri:" + v.URI + ", method:" + v.Method + ", user-agent:" + v.UserAgent + ", remote-ip:" + v.RemoteIP + ", error:" + v.Error.Error() + "\n"
				if _, err := fd.WriteString(s); err != nil {
					return err
				}
				defer fd.Close()
			}
			fmt.Printf("REQUEST: time: %v, uri: %v, status: %v, method: %v, user-agent: %v\n", v.StartTime.Local().String(), v.URI, v.Status, v.Method, v.UserAgent)
			return nil
		},
	}

}
