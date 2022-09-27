package service

import (
	"caesar/controller/forms"
	"caesar/internal"
	"caesar/model"
)

// CreateAccount - 创建平台账号
func CreateAccount(r *forms.Account, id int64) (int64, error) {
	a := new(model.Account)
	if err := IsMainPassword(id, r.MainPassword); err != nil {
		return 0, err
	}
	err := a.CreateAccount(r, id)
	if err != nil {
		return 0, err
	}
	return a.Id, err
}

// ReadAccount - 读取平台账号信息
func ReadAccount(a *forms.AccountParams, id int64) (map[string]interface{}, error) {
	m := new(model.Account)
	if err := IsMainPassword(id, a.MainPassword); err != nil {
		return nil, err
	}
	err := m.FindById(id)
	if err != nil {
		return nil, err
	}
	// 解密
	name, err := internal.AesDecrypt(m.Name, a.MainPassword, 16)
	if err != nil {
		return nil, err
	}
	email, err := internal.AesDecrypt(m.Email, a.MainPassword, 16)
	if err != nil {
		return nil, err
	}
	password, err := internal.AesDecrypt(m.Password, a.MainPassword, 16)
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
