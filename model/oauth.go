package model

import (
	"caesar/global"

	_ "github.com/go-sql-driver/mysql"
)

type Oauth struct {
	Id        int64  `json:"id" db:"id"`
	UserId    int64  `json:"user_id" db:"user_id"`
	ClientIp  string `json:"client_ip" db:"client_ip"`
	Token     string `json:"token" db:"token"`
	Revoked   uint8  `db:"revoked"`
	ExpiresAt int64  `json:"expires_at" db:"expires_at"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

func (o *Oauth) TotalByIp(ip string, userId int64) (int, error) {
	s := "SELECT COUNT(`id`) FROM `oauth_access_tokens` WHERE `user_id` = ? AND `client_ip` = ?"
	var count int
	err := global.DB.Get(&count, s, userId, ip)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (o *Oauth) Find(ip string, userId int64) error {
	s := "SELECT * FROM `oauth_access_tokens` WHERE `user_id` = ? AND `client_ip` = ?"
	err := global.DB.Get(o, s, userId, ip)
	if err != nil {
		return err
	}
	return nil
}

func (o *Oauth) Create(v map[string]interface{}) error {
	s := "INSERT INTO `oauth_access_tokens`(`user_id`, `client_ip`, `token`, `expires_at`, `created_at`, " +
		"`updated_at`) VALUES(:user_id, :client_ip, :token, :expires_at, :created_at, :updated_at)"
	_, err := global.DB.NamedExec(s, v)
	if err != nil {
		return err
	}
	return nil
}

func (o *Oauth) UpdateTokenAndExpiresAt(v map[string]interface{}) error {
	s := "UPDATE `oauth_access_tokens` SET `token`= :token, `expires_at` = :expires_at, " +
		"`updated_at` = :updated_at WHERE `user_id` = :user_id AND `client_ip` = :client_ip"
	_, err := global.DB.NamedExec(s, v)
	if err != nil {
		return err
	}
	return nil
}
