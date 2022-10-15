package internal

import (
	"bytes"
	"caesar/global"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SendMail(to []string, subject string, msg string) bool {
	emailConfig := global.Setting.Email
	auth := smtp.PlainAuth("", emailConfig.Name, emailConfig.Password, emailConfig.Host)
	smtpAddr := fmt.Sprintf("%s:%v", emailConfig.Host, emailConfig.Port)
	err := smtp.SendMail(smtpAddr, auth, emailConfig.Name, to, []byte("From: 凯撒密码网管理员<jie893609357@163.com>\r\n"+
		"Subject:"+subject+"\r\n"+
		"\r\n"+msg))
	return err != nil
}

func AesEncrypt(data, key string) (string, error) {
	keyByte, err := keyPadding([]byte(key))
	if err != nil {
		return "", err
	}
	encrypted := cryptor.AesCbcEncrypt([]byte(data), keyByte)
	if len(encrypted) == 0 {
		return "", errors.New("加密失败")
	}
	return cryptor.Base64StdEncode(string(encrypted)), nil
}

func AesDecrypt(data, key string) (string, error) {
	keyByte, err := keyPadding([]byte(key))
	if err != nil {
		return "", err
	}
	decodeData := cryptor.Base64StdDecode(data)
	if decodeData == "" {
		return "", errors.New("base64 解码失败")
	}
	decrypted := cryptor.AesCbcDecrypt([]byte(decodeData), keyByte)
	if len(decrypted) == 0 {
		return "", errors.New("解密失败")
	}
	return string(decrypted), nil
}

func keyPadding(key []byte) ([]byte, error) {
	keySize := global.Setting.AesKeyLength
	switch keySize {
	default:
		return nil, errors.New("AesKeyLength 配置有误")
	case 16, 24, 32:
		break
	}

	padding := keySize - len(key)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(key, padText...), nil
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
				file := fmt.Sprintf("%s%s.log", global.Setting.LogAddr, datetime.GetNowDate())
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
