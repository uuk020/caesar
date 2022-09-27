package model

import (
	"caesar/controller/forms"
	"caesar/global"
	caesarInternal "caesar/internal"

	_ "github.com/go-sql-driver/mysql"
)

const (
	LogSatasCreate = iota
	LogSatasView
	LogSatasEdit
	LogSatasShare
)

var StatusText = [4]string{
	LogSatasCreate: "创建",
	LogSatasView:   "查看",
	LogSatasEdit:   "编辑",
	LogSatasShare:  "分享",
}

type Account struct {
	Id        int64  `json:"id" db:"id"`
	UserId    int64  `json:"user_id" db:"user_id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Platform  string `json:"platform" db:"platform"`
	Url       string `json:"url" db:"url"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

// CreateAccount 创建平台账户
func (a *Account) CreateAccount(f *forms.Account, userId int64) error {
	tx := global.DB.MustBegin()

	nowUnix := caesarInternal.GetNowTimestamp()

	// 加密
	encryptName, err := caesarInternal.AesEncrypt(f.Name, f.MainPassword, 16)
	if err != nil {
		return err
	}
	encryptEmail, err := caesarInternal.AesEncrypt(f.Email, f.MainPassword, 16)
	if err != nil {
		return err
	}
	encryptPassword, err := caesarInternal.AesEncrypt(f.Password, f.MainPassword, 16)
	if err != nil {
		return err
	}

	accounSql := "INSERT INTO `account`(`user_id`, `name`, `email`, `password`, `platform`, url, `created_at`, `updated_at`) VALUES(:user_id, :name, :email, :password, :platform, :url, :created_at, :updated_at)"
	ret, err := tx.NamedExec(accounSql, map[string]interface{}{
		"user_id":    userId,
		"name":       encryptName,
		"email":      encryptEmail,
		"password":   encryptPassword,
		"platform":   f.Platform,
		"url":        f.Url,
		"created_at": nowUnix,
		"updated_at": nowUnix,
	})
	if err != nil {
		return err
	}
	id, err := ret.LastInsertId()

	logSql := "INSERT INTO `account_log`(`account_id`, `type`, `created_at`, `updated_at`) VALUES (:account_id, :type, :created_at, :updated_at)"
	_, err = tx.NamedExec(logSql, map[string]interface{}{
		"account_id": id,
		"type":       LogSatasCreate,
		"created_at": nowUnix,
		"updated_at": nowUnix,
	})
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	a.Id = id
	a.UserId = userId
	a.Name = encryptName
	a.Email = encryptEmail
	a.Password = encryptPassword
	a.CreatedAt = nowUnix
	a.UpdatedAt = nowUnix
	return nil
}

// FindById 根据平台 ID 查询平台信息
func (a *Account) FindById(id int64) error {
	s := "SELECT * FROM `account` WHERE id = ? ORDER BY `id` DESC LIMIT 1"
	err := global.DB.Get(a, s, id)
	if err != nil {
		return err
	}

	nowUnix := caesarInternal.GetNowTimestamp()

	logSql := "INSERT INTO `account_log`(`account_id`, `type`, `created_at`, `updated_at`) VALUES (:account_id, :type, :created_at, :updated_at)"
	_, err = global.DB.NamedExec(logSql, map[string]interface{}{
		"account_id": id,
		"type":       LogSatasView,
		"created_at": nowUnix,
		"updated_at": nowUnix,
	})
	if err != nil {
		return err
	}

	return nil
}
