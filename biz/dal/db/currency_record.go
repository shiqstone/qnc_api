package db

import (
	"qnc/pkg/constants"
)

type CurrencyRecord struct {
	ID         int64   `json:"id"`
	Currency   string  `json:"currency"`
	BaseXt     float64 `json:"base_xt"`
	LatestXt   float64 `json:"latest_xt"`
	UpdateDate string  `json:"update_date"`
	Remark     string  `json:"remark"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}

func (CurrencyRecord) TableName() string {
	return constants.CurrencyTableName
}

func CreateCurrency(rec *CurrencyRecord) (int64, error) {
	err := DB.Create(rec).Error
	if err != nil {
		return 0, err
	}
	return rec.ID, err
}

func QueryByLatestList() (*[]CurrencyRecord, error) {
	var recs []CurrencyRecord
	if err := DB.Where("id IN (SELECT MAX(id) FROM qnc_currency_record GROUP BY currency)").Find(&recs).Error; err != nil {
		return nil, err
	}

	return &recs, nil
}

func UpdateCurrency(rec *CurrencyRecord) (*CurrencyRecord, error) {
	err := DB.Updates(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, err
}
