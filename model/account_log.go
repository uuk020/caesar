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

func LogSelect(accountId int64, page, pageSize int) ([]AccountLog, int) {
	var count int
	err := global.DB.Get(&count, "SELECT count(*) FROM `account_log` WHERE `account_id` = ?", accountId)
	if err != nil {
		return nil, count
	}

	offset := pageSize * (page - 1)
	rows, err := global.DB.Queryx("SELECT * FROM `account_log` WHERE `account_id` = ? ORDER BY `id` DESC LIMIT ? OFFSET ?", accountId, pageSize, offset)
	if err != nil {
		return nil, count
	}

	var m []AccountLog
	for rows.Next() {
		a := AccountLog{}
		err := rows.StructScan(&a)
		if err != nil {
			return nil, 0
		}
		m = append(m, a)
	}
	return m, count
}
