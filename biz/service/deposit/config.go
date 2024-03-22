package deposit

import (
	"context"
	"encoding/json"
	"fmt"
	"qnc/biz/dal/db"
	"qnc/biz/model/kv"

	"github.com/cloudwego/hertz/pkg/app"
)

type DepositService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewDepositService(ctx context.Context, c *app.RequestContext) *DepositService {
	return &DepositService{ctx: ctx, c: c}
}

func (s *DepositService) GetDepositProds(countryCode string) (resp *kv.DepositConfResponse, err error) {
	confs, err := db.QueryDepositConfigByCountryCode(countryCode)
	if err != nil {
		return nil, err
	}
	resp = new(kv.DepositConfResponse)

	var products [](map[string]interface{})
	for _, v := range *confs {
		prod := make(map[string]interface{})
		prod["coin"] = v.Coins
		prod["name"] = fmt.Sprintf("%.0f%s", v.Coins, " Coins")
		prod["price"] = v.ActualPrice
		prod["sign"] = v.Sign
		products = append(products, prod)
	}
	resp.Products = products

	payways, err := db.QueryByName("payway")
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
