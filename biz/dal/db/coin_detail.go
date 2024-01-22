package db

import (
	"qnc/pkg/constants"
)

type CoinDetail struct {
	ID         int64   `json:"id"`
	UserId     int64   `json:"uid" gorm:"column:uid"`
	Type       int16   `json:"type"`
	EventId    int64   `json:"evnet_id"`
	Incoming   float64 `json:"incoming"`
	Expend     float64 `json:"expend"`
	Balance    float64 `json:"balance"`
	Remark     string  `json:"remark"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}

func (CoinDetail) TableName() string {
	return constants.CoinDetailTableName
}

// CreateUser create user info
func CreateCoinDetail(detail *CoinDetail) (int64, error) {
	err := DB.Create(detail).Error
	if err != nil {
		return 0, err
	}
	return detail.ID, err
}
