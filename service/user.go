package service

import (
	"caesar/controller/forms"
	"caesar/global"
	"caesar/internal"
	"caesar/model"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(r *forms.Register) (int64, error) {
	u := new(model.User)
	code := internal.RandomCode()
	key := "cache:activation:" + r.Email
	val, err := global.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if val != "" {
		return 0, errors.New("验证码已经发送过了, 5分钟后再试")
	}
	id, err := u.Create(r)
	if err != nil {
		return 0, err
	}
	msg := fmt.Sprintf("欢迎注册账号, 验证码为 %s(5 分钟过期), 请勿转发他人", code)
	if internal.SendMail([]string{r.Email}, "凯撒密码网注册验证码", msg) {
		if err := global.Redis.Set(key, code, 5*time.Minute).Err(); err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("邮件发送失败, 请稍后再试")
}

func AgainSendEmail(a *forms.SendMailS) error {
	u := new(model.User)
	code := internal.RandomCode()

	key := "cache:activation:" + a.Email
	if a.Type == "password" {
		key = "cache:password:" + a.Email
	}

	val, err := global.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if val != "" {
		return errors.New("验证码已经发送过了, 5分钟再试")
	}
	if err := u.Find("email", a.Email); err != nil {
		return err
	}

	msg := fmt.Sprintf("欢迎注册账号, 验证码为 %s(5 分钟过期), 请勿转发他人", code)
	subject := "欢迎注册凯撒密码管理-注册激活码"

	if a.Type == "activation" {
		if u.Status == 1 {
			return errors.New("该用户已经激活了")
		}
	} else {
		msg = fmt.Sprintf("重置密码, 验证码为 %s(5 分钟过期), 请勿转发他人", code)
		subject = "凯撒密码管理-重置密码"
	}

	if u.Status == 2 {
		return errors.New("查询不到用户")
	}

	if internal.SendMail([]string{a.Email}, subject, msg) {
		if err := global.Redis.Set(key, code, 5*time.Minute).Err(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("邮件发送失败, 请稍后再试")
}

// Activation 激活用户
func Activation(a *forms.Activation) error {
	key := "cache:activation:" + a.Email
	val, err := global.Redis.Get(key).Result()
	if err != nil {
		return err
	}
	if val != a.Code {
		return errors.New("验证码不一致")
	}
	u := new(model.User)
	err = u.Find("email", a.Email)
	if err != nil {
		return err
	}
	if u.Status == 1 {
		return errors.New("该用户已经激活")
	}
	if u.Status == 2 {
		return errors.New("查询不到该用户")
	}
	_, err = u.UpdateOneField("status", 1, "email", a.Email)
	if err != nil {
		return err
	}
	err = global.Redis.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func ResetPassword(r *forms.ResetPassword) error {
	key := "cache:password:" + r.Email
	val, err := global.Redis.Get(key).Result()
	if err != nil {
		return err
	}
	if val != r.Code {
		return errors.New("验证码不一致")
	}
	u := new(model.User)
	err = u.Find("email", r.Email)
	if err != nil {
		return err
	}
	if u.Status == 2 {
		return errors.New("查询不到该用户")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = u.UpdateOneField("password", hash, "email", r.Email)
	if err != nil {
		return err
	}
	err = global.Redis.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}

func Login(l *forms.Login, ip string) (string, error) {
	u := new(model.User)

	field := "name"
	reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if reg.MatchString(l.UserName) {
		field = "email"
	}

	if err := u.Find(field, l.UserName); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
	if err != nil {
		return "", errors.New("密码输入有误")
	}
	if u.Status == 0 {
		return "", errors.New("用户尚未激活")
	}
	if u.Status == 2 {
		return "", errors.New("查询不到该用户")
	}

	o := new(model.Oauth)
	count, err := o.TotalByIp(ip, u.Id)
	if err != nil {
		return "", err
	}
	if count == 3 {
		return "", errors.New("该账号已经在3个设备登录过")
	}

	now := time.Now()
	expiresAt := now.Add(time.Hour * 48).Unix()
	createdAt := now.Unix()

	claims := &internal.JwtCustomClaims{
		u.Id,
		u.Name,
		u.Email,
		jwt.StandardClaims{ExpiresAt: expiresAt},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(global.Setting.SecretKey))
	if err != nil {
		return "", err
	}

	err = o.Find(ip, u.Id)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if o.Id > 0 {
		if err := o.UpdateTokenAndExpiresAt(map[string]interface{}{
			"token":      t,
			"expires_at": expiresAt,
			"updated_at": createdAt,
			"user_id":    u.Id,
			"client_ip":  ip,
		}); err != nil {
			return "", err
		}
	} else {
		if err := o.Create(map[string]interface{}{
			"user_id":    u.Id,
			"client_ip":  ip,
			"token":      t,
			"expires_at": expiresAt,
			"created_at": createdAt,
			"updated_at": createdAt,
		}); err != nil {
			return "", err
		}
	}

	return t, nil
}

func Logout(id int64) error {
	u := new(model.User)
	if err := u.Find("id", id); err != nil {
		return err
	}
	if u.Status == 2 {
		return errors.New("查询不到该用户")
	}
	_, err := u.UpdateOneField("status", 2, "id", id)
	if err != nil {
		return err
	}
	return nil
}

func Me(id int64) (*echo.Map, error) {
	u := new(model.User)
	if err := u.Find("id", id); err != nil {
		return nil, err
	}
	if u.Status == 2 {
		return nil, errors.New("已经注销过了")
	}
	r := &echo.Map{
		"real_name": u.RealName,
		"phone":     u.Phone,
		"email":     u.Email,
	}
	return r, nil
}

func UpdateMe(id int64, m map[string]interface{}) error {
	u := new(model.User)
	if err := u.Find("id", id); err != nil {
		return err
	}
	if u.Status == 2 {
		return errors.New("已经注销过了")
	}
	_, err := u.UpdateMultField(id, m)
	if err != nil {
		return err
	}
	return nil
}
