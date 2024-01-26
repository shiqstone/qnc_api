package service

import (
	"context"
	"encoding/json"
	"qnc/biz/dal/db"
	"qnc/biz/model/kv"

	"github.com/cloudwego/hertz/pkg/app"
)

type KvService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewKvService create Kv service
func NewKvService(ctx context.Context, c *app.RequestContext) *KvService {
	return &KvService{ctx: ctx, c: c}
}

func (s *KvService) GetDepositConf() (resp *kv.DepositConfResponse, err error) {
	kvs, err := db.QueryByName("deposit")
	if err != nil {
		return nil, err
	}
	resp = new(kv.DepositConfResponse)

	var products [](map[string]interface{})
	for _, v := range kvs {
		prod := make(map[string]interface{})
		err = json.Unmarshal([]byte(v.Value), &prod)
		if err != nil {
			continue
		}
		products = append(products, prod)
	}
	resp.Products = products

	payways, err := db.QueryByName("payways")
	if err != nil {
		return nil, err
	}
	var ways [](map[string]interface{})
	for _, v := range payways {
		pw := make(map[string]interface{})
		err = json.Unmarshal([]byte(v.Value), &pw)
		if err != nil {
			continue
		}
		ways = append(ways, pw)
	}
	resp.Payways = ways

	tips, err := db.QueryByName("tips")
	if err != nil {
		return nil, err
	}
	resp.Tips = tips[0].Value

	return resp, nil
}
