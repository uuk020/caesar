package service

import (
	"caesar/controller/forms"
	"caesar/global"
	"caesar/internal"
	"caesar/model"
	"strconv"
)

// CreateAccount 创建平台账号
func CreateAccount(a *forms.AccountCreate, id int64) (int64, error) {
	m := new(model.Account)
	if err := IsMainPassword(id, a.MainPassword); err != nil {
		return 0, err
	}
	err := m.CreateAccount(a, id)
	if err != nil {
		return 0, err
	}
	return m.Id, err
}

// UpdateAccount 更新平台信息
func UpdateAccount(a *forms.AccountUpdate, userId int64) (map[string]interface{}, error) {
	m := new(model.Account)
	if err := IsMainPassword(userId, a.MainPassword); err != nil {
		return nil, err
	}

	accountId, err := strconv.Atoi(a.ID)
	if err != nil {
		return nil, err
	}

	f := model.UpdateAccountField{
		ID:           int64(accountId),
		Name:         a.Name,
		Email:        a.Email,
		Password:     a.Password,
		MainPassword: a.MainPassword,
	}
	err = m.UpdateAccount(f)
	if err != nil {
		return nil, err
	}
	r := make(map[string]interface{}, 5)
	r["id"] = m.Id
	r["name"] = m.Name
	r["email"] = m.Email
	r["password"] = m.Password
	r["updated_at"] = m.UpdatedAt
	return r, nil
}

// ReadAccount 读取平台账号信息
func ReadAccount(a *forms.AccountRead, id int64) (map[string]interface{}, error) {
	m := new(model.Account)
	if err := IsMainPassword(id, a.MainPassword); err != nil {
		return nil, err
	}
	err := m.FindById(id)
	if err != nil {
		return nil, err
	}
	// 解密
	name, err := internal.AesDecrypt(m.Name, a.MainPassword, global.Setting.AesKeyLength)
	if err != nil {
		return nil, err
	}
	email, err := internal.AesDecrypt(m.Email, a.MainPassword, global.Setting.AesKeyLength)
	if err != nil {
		return nil, err
	}
	password, err := internal.AesDecrypt(m.Password, a.MainPassword, global.Setting.AesKeyLength)
	if err != nil {
		return nil, err
	}

	r := map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": password,
	}
	return r, nil
}

// DeleteAccount 删除平台账号
func DeleteAccount(mainPassword string, accountId, userId int) error {
	m := new(model.Account)
	if err := IsMainPassword(int64(userId), mainPassword); err != nil {
		return err
	}
	err := m.FindById(int64(accountId))
	if err != nil {
		return err
	}
	if err := m.DeleteAccount(int64(accountId)); err != nil {
		return err
	}
	return nil
}

// GetLog 获取日志
func GetLog(accountId, page, pageSize int) ([]interface{}, int) {
	d, c := model.LogSelect(int64(accountId), page, pageSize)
	var r []interface{}
	for _, v := range d {
		r = append(r, v)
	}
	return r, c
}
