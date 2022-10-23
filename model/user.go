package model

import (
	"caesar/controller/forms"
	"caesar/global"
	"database/sql"
	"errors"
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/duke-git/lancet/v2/datetime"
)

type User struct {
	Id           int64  `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Password     string `db:"password"`
	MainPassword string `db:"main_password"`
	Email        string `json:"email" db:"email"`
	RealName     string `db:"real_name"`
	Phone        string `db:"phone"`
	Status       uint8  `db:"status"`
	CreatedAt    int64  `json:"created_at" db:"created_at"`
	UpdatedAt    int64  `json:"updated_at" db:"updated_at"`
}

func (u *User) Find(f string, v interface{}) error {
	m := map[string]bool{
		"id":    true,
		"name":  true,
		"email": true,
		"phone": true,
	}
	if _, ok := m[f]; ok {
		s := "SELECT * FROM `user` WHERE " + f + " = ? ORDER BY `id` DESC LIMIT 1"
		err := global.DB.Get(u, s, v)
		if err != nil {
			return err
		}
	} else {
		return errors.New("查询字段不存在, 字段: " + f)
	}

	return nil
}

func (u *User) Create(r *forms.Register) (int64, error) {
	// 判断是否已有用户
	type re struct {
		Id     int64 `db:"id"`
		Status uint8 `db:"status"`
	}
	queryRe := new(re)
	err := global.DB.Get(queryRe, "SELECT `id`, `status` FROM `user` WHERE "+
		"`name` = ? OR `email` = ? OR `phone` = ?",
		r.UserName, r.Email, r.Phone)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	nowUnix := datetime.NewUnixNow().ToUnix()

	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	mHash, err := bcrypt.GenerateFromPassword([]byte(r.MainPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	data := map[string]interface{}{
		"name":          r.UserName,
		"password":      hash,
		"main_password": mHash,
		"email":         r.Email,
		"real_name":     r.RealName,
		"phone":         r.Phone,
		"status":        0,
		"created_at":    nowUnix,
		"updated_at":    nowUnix,
	}

	if queryRe.Id > 0 {
		if queryRe.Status == 2 {
			s := "UPDATE `user` SET `name` = :name, `password` = :password, `main_password` = :main_password, `email` = :email, " +
				"`real_name` = :real_name, `phone` = :phone, " + "`status` = :status, `created_at` = :created_at, `updated_at` = :updated_at WHERE id = :id"
			data["id"] = queryRe.Id
			if _, err := global.DB.NamedExec(s, data); err != nil {
				return 0, err
			}
			return queryRe.Id, nil
		}
		return 0, errors.New("已经创建过用户了")
	}
	s := "INSERT INTO `user`(name, password, main_password, email, real_name, phone, status, created_at, updated_at) VALUES (:name, " +
		":password, :main_password, :email, :real_name, :phone, :status, :created_at, :updated_at)"
	ret, err := global.DB.NamedExec(s, data)
	if err != nil {
		return 0, err
	}
	lastID, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (u *User) UpdateOneField(uF string, uV interface{}, wF string, wV interface{}) (int64, error) {
	nowUnix := datetime.NewUnixNow().ToUnix()

	s := "UPDATE `user` SET " + uF + " = ?, updated_at = ? WHERE " + wF + " = ?"
	result, err := global.DB.Exec(s, uV, nowUnix, wV)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (u *User) UpdateMultField(id int64, m map[string]interface{}) (int64, error) {
	var fieldBuilder, sqlBuilder strings.Builder
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	d := make(map[string]interface{})
	for i, k := range keys {
		if i != len(keys)-1 {
			fieldBuilder.WriteString(k + " = :" + k + ",")
		} else {
			fieldBuilder.WriteString(k + " = :" + k)
		}
		d[k] = m[k]
	}
	sqlBuilder.WriteString("UPDATE `user` SET ")
	sqlBuilder.WriteString(fieldBuilder.String())
	sqlBuilder.WriteString(" WHERE `id` = :id")
	d["id"] = id
	result, err := global.DB.NamedExec(sqlBuilder.String(), d)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (u *User) Logout(id int64) error {
	tx := global.DB.MustBegin()

	nowUnix := datetime.NewUnixNow().ToUnix()
	userSql := "UPDATE `user` SET status = 2, updated_at = ? WHERE id = ?"
	_, err := tx.Exec(userSql, nowUnix, id)
	if err != nil {
		return err
	}
	accountSql := "DELETE FROM account WHERE user_id = ?"
	_, err = tx.Exec(accountSql, id)
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
