package db

import (
	"qnc/pkg/constants"
)

type DepositConfig struct {
	ID          int64   `json:"id"`
	CountryCode string  `json:"country_code"`
	Coins       float64 `json:"coins"`
	Currency    string  `json:"currency"`
	Sign        string  `json:"sign"`
	Price       float64 `json:"price"`
	ActualPrice float64 `json:"actual_price"`
	Remark      string  `json:"remark"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

func (DepositConfig) TableName() string {
	return constants.DepositConfigTableName
}

func CreateDepositConfig(detail *DepositConfig) (int64, error) {
	err := DB.Create(detail).Error
	if err != nil {
		return 0, err
	}
	return detail.ID, err
}

func QueryDepositConfigByCountryCode(countryCode string) (*[]DepositConfig, error) {
	var recs []DepositConfig
	if err := DB.Where("country_code = ?", countryCode).Find(&recs).Error; err != nil {
		return nil, err
	}
	return &recs, nil
}

func UpdateDepositConfig(rec *DepositConfig) (*DepositConfig, error) {
	err := DB.Updates(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, err
}
