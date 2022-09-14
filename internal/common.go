package internal

import (
	"caesar/global"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
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
	if err != nil {
		return false
	}
	return true
}
