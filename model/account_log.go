package model

import "caesar/global"

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

type AccountLog struct {
	Id        int64 `json:"id" db:"id"`
	AccountId int64 `json:"account_id" db:"account_id"`
	Status    uint8 `json:"type" db:"type"`
	CreatedAt int64 `json:"created_at" db:"created_at"`
	UpdatedAt int64 `json:"updated_at" db:"updated_at"`
}

func (l *AccountLog) Select(accountId int64) (map[string]interface{}, error) {
	var count int
	err := global.DB.Get(&count, "SELECT count(*) FROM `account_log` WHERE `account_id` = ?", accountId)
	if err != nil {
		return nil, err
	}
	m := make([]map[string]interface{}, count)
	rows, err := global.DB.Queryx("SELECT * FROM `account_log` WHERE `account_id` = ?", accountId)
	if err != nil {
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.StructScan(&l)
		if err != nil {
			return nil, err
		}
		m[i] = map[string]interface{}{
			"id":         l.Id,
			"account_id": l.AccountId,
			"type":       l.Status,
			"type_text":  StatusText[l.Status],
			"created_at": l.CreatedAt,
			"updated_at": l.UpdatedAt,
		}
		i++
	}
	r := make(map[string]interface{}, 2)
	r["data"] = m
	r["count"] = count
	return r, nil
}
