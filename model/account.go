package model

import (
	"caesar/controller/forms"
	"caesar/global"
	caesarInternal "caesar/internal"
	"strings"

	"github.com/duke-git/lancet/v2/datetime"
)

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
func (a *Account) CreateAccount(f *forms.AccountCreate, userId int64) error {
	tx := global.DB.MustBegin()

	nowUnix := datetime.NewUnixNow().ToUnix()

	// 加密
	encryptName, err := caesarInternal.AesEncrypt(f.Name, f.MainPassword)
	if err != nil {
		return err
	}
	encryptEmail, err := caesarInternal.AesEncrypt(f.Email, f.MainPassword)
	if err != nil {
		return err
	}
	encryptPassword, err := caesarInternal.AesEncrypt(f.Password, f.MainPassword)
	if err != nil {
		return err
	}

	accounSql := "INSERT INTO `account`(`user_id`, `name`, `email`, `password`, `platform`, url, `created_at`, `updated_at`) VALUES (:user_id, :name, :email, :password, :platform, :url, :created_at, :updated_at)"
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
	if err != nil {
		return err
	}

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

type UpdateAccountField struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	MainPassword string
}

func (a *Account) UpdateAccount(u UpdateAccountField, userId int64) error {
	tx := global.DB.MustBegin()

	nowUnix := datetime.NewUnixNow().ToUnix()

	err := tx.Get(a, "SELECT * FROM account WHERE `id` = ? AND `user_id` = ?", u.ID, userId)
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

	// 加密
	encryptName, err := caesarInternal.AesEncrypt(u.Name, u.MainPassword)
	if err != nil {
		return err
	}
	encryptEmail, err := caesarInternal.AesEncrypt(u.Email, u.MainPassword)
	if err != nil {
		return err
	}
	encryptPassword, err := caesarInternal.AesEncrypt(u.Password, u.MainPassword)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"name":       encryptName,
		"email":      encryptEmail,
		"password":   encryptPassword,
		"updated_at": nowUnix,
		"id":         u.ID,
	}
	accounSql := "UPDATE `account` SET `name`=:name, `email`=:email, `password`=:password, `updated_at`=:updated_at WHERE id = :id"
	_, err = tx.NamedExec(accounSql, data)
	if err != nil {
		return err
	}

	logSql := "INSERT INTO `account_log`(`account_id`, `type`, `created_at`, `updated_at`) VALUES (:account_id, :type, :created_at, :updated_at)"
	_, err = tx.NamedExec(logSql, map[string]interface{}{
		"account_id": u.ID,
		"type":       LogSatasEdit,
		"created_at": nowUnix,
		"updated_at": nowUnix,
	})
	if err != nil {
		return err
	}
	return nil
}

// FindById 根据平台 ID 查询平台信息
func (a *Account) FindById(id, userId int64) error {
	s := "SELECT * FROM `account` WHERE `id` = ? AND `user_id` = ? ORDER BY `id` DESC LIMIT 1"
	err := global.DB.Get(a, s, id, userId)
	if err != nil {
		return err
	}

	nowUnix := datetime.NewUnixNow().ToUnix()

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

// DeleteAccount 删除平台账号
func (a *Account) DeleteAccount(id, userId int64) error {
	tx := global.DB.MustBegin()
	s := "DELETE FROM `account` WHERE `id` = ? AND `user_id` = ?"
	_, err := tx.Exec(s, id, userId)
	if err != nil {
		return err
	}

	s1 := "DELETE FROM `account_log` WHERE `account_id` = ?"
	_, err = tx.Exec(s1, id)
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
	return nil
}

// getListWhere 构建列表 where 条件
func getListWhere(a *forms.AccountList) (string, map[string]interface{}) {
	var s strings.Builder
	m := make(map[string]any, 3)
	if a.Platform != "" {
		s.WriteString(" AND `platform` LIKE :platform")
		m["platform"] = "%" + a.Platform + "%"
	}
	if a.DateStart != "" && a.DateEnd != "" {
		s.WriteString(" AND `created_at` BETWEEN :date_start AND :date_end")
		sT, _ := datetime.FormatStrToTime(a.DateStart+"00:00:00", "yyyy-mm-dd hh:mm:ss")
		eT, _ := datetime.FormatStrToTime(a.DateStart+"23:59:59", "yyyy-mm-dd hh:mm:ss")
		m["date_start"] = sT.Unix()
		m["date_end"] = eT.Unix()
	} else if a.DateStart != "" {
		s.WriteString(" AND `created_at` >= :date_start")
		sT, _ := datetime.FormatStrToTime(a.DateStart+"00:00:00", "yyyy-mm-dd hh:mm:ss")
		m["date_start"] = sT.Unix()
	} else if a.DateEnd != "" {
		s.WriteString(" AND `created_at` <= :date_end")
		eT, _ := datetime.FormatStrToTime(a.DateStart+"23:59:59", "yyyy-mm-dd hh:mm:ss")
		m["date_end"] = eT.Unix()
	}

	return s.String(), m
}

// AccountList 列表
func AccountList(a *forms.AccountList, userId int64) ([]Account, int) {
	var (
		count int
		s, s1 strings.Builder
		r     []Account
	)

	offset := a.PageSize * (a.Page - 1)
	m := map[string]interface{}{
		"user_id":   userId,
		"offset":    offset,
		"page_size": a.PageSize,
	}
	m1 := map[string]interface{}{
		"user_id": userId,
	}

	s.WriteString("SELECT * FROM `account` WHERE `user_id` = :user_id")
	s1.WriteString("SELECT count(*) FROM `account` WHERE `user_id` = :user_id")
	whereSql, where := getListWhere(a)
	if whereSql != "" && len(where) != 0 {
		s.WriteString(whereSql)
		s1.WriteString(whereSql)
		for k, v := range where {
			m1[k] = v
			m[k] = v
		}
	}

	s.WriteString(" ORDER BY `id` DESC LIMIT :page_size OFFSET :offset")
	s1.WriteString(" LIMIT 1")

	rows, err := global.DB.NamedQuery(s.String(), m)
	if err != nil {
		return r, 0
	}

	rows1, err := global.DB.NamedQuery(s1.String(), m1)
	if err != nil {
		return r, 0
	}
	for rows1.Next() {
		if err := rows1.Scan(&count); err != nil {
			return r, 0
		}
	}

	for rows.Next() {
		d := Account{}
		err = rows.StructScan(&d)
		if err != nil {
			return nil, 0
		}
		r = append(r, d)
	}
	return r, count
}
