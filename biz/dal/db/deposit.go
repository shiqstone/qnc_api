package db

import (
	"qnc/pkg/constants"
	"qnc/pkg/errno"
)

type Deposit struct {
	ID          int64   `json:"id"`
	UserId      int64   `json:"uid" gorm:"column:uid"`
	DepositId   string  `json:"deposit_id"`
	Amount      float64 `json:"amount"`
	Status      int16   `json:"status"`
	Currency    string  `json:"currency"`
	PayMode     int16   `json:"pay_mode"`
	PayBank     string  `json:"pay_bank"`
	PayChannel  int16   `json:"pay_channel"`
	BankTradeNo string  `json:"bank_trade_no"`
	TotalRefund float64 `json:"total_refund"`
	Ip          string  `json:"ip"`
	Ext         string  `json:"ext"`
	FinishTime  int64   `json:"finish_time"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

func (Deposit) TableName() string {
	return constants.DepositTableName
}

func CreateDeposit(detail *Deposit) (int64, error) {
	err := DB.Create(detail).Error
	if err != nil {
		return 0, err
	}
	return detail.ID, err
}

func QueryByDepositId(depositId string) (*Deposit, error) {
	var rec Deposit
	if err := DB.Where("deposit_id = ?", depositId).Find(&rec).Error; err != nil {
		return nil, err
	}
	if rec.DepositId != depositId {
		err := errno.RecordNotExistErr
		return nil, err
	}
	return &rec, nil
}

func UpdateDeposit(rec *Deposit) (*Deposit, error) {
	err := DB.Updates(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, err
}
