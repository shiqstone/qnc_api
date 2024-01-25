package db

import (
	"qnc/pkg/constants"
	"qnc/pkg/errno"
)

type Order struct {
	ID         int64   `json:"id"`
	UserId     int64   `json:"uid" gorm:"column:uid"`
	ProdName   string  `json:"prod_name"`
	ProdId     int64   `json:"prod_id"`
	RealCost   float64 `json:"real_cost"`
	BaseCost   float64 `json:"base_cost"`
	Status     int16   `json:"status"`
	Ip         string  `json:"ip"`
	Remark     string  `json:"remark"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}

func (Order) TableName() string {
	return constants.OrderTableName
}

func CreateOrder(order *Order) (int64, error) {
	err := DB.Create(order).Error
	if err != nil {
		return 0, err
	}
	return order.ID, err
}

func QueryById(id string) (*Order, error) {
	var rec Order
	if err := DB.Where("id = ?", id).Find(&rec).Error; err != nil {
		return nil, err
	}
	if rec == (Order{}) {
		err := errno.RecordNotExistErr
		return nil, err
	}
	return &rec, nil
}

func UpdateOrder(rec *Order) (*Order, error) {
	err := DB.Updates(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, err
}
