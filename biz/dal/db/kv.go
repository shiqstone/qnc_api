package db

import (
	"qnc/pkg/constants"
)

type Kv struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (Kv) TableName() string {
	return constants.KvTableName
}

func CreateKv(detail *Kv) (int64, error) {
	err := DB.Create(detail).Error
	if err != nil {
		return 0, err
	}
	return detail.ID, err
}

func QueryByName(name string) (kvs []Kv, err error) {
	if err = DB.Where("name like ? and status = ?", name+"%", 1).Find(&kvs).Error; err != nil {
		return nil, err
	}
	return
}

func QueryAllByName(name string) (kvs []Kv, err error) {
	if err = DB.Where("name like ?", name+"%").Find(&kvs).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateKv(rec *Kv) (*Kv, error) {
	err := DB.Updates(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, err
}
