package currency

import (
	"encoding/json"
	"math"
	"net/http"
	"qnc/biz/dal/db"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Tracking() {
	ret, err := getLatestXt()
	if err != nil {
		hlog.Error(err)
	}
	udate := ret["date"].(string)
	xts := ret["cny"].(map[string]interface{})

	curRec, err := db.QueryByLatestList()
	if err != nil {
		hlog.Error("query db record err, ", err)
	}

	ts := time.Now().Unix()
	for _, rec := range *curRec {
		cur := strings.ToLower(rec.Currency)
		if val, ok := xts[cur]; ok {
			xt := val.(float64)
			chgRt := math.Abs(rec.BaseXt-xt) / rec.BaseXt
			if chgRt > 0.05 {
				// rate change over 5%
				_, err = db.CreateCurrency(&db.CurrencyRecord{
					Currency:   rec.Currency,
					BaseXt:     xt,
					LatestXt:   xt,
					UpdateDate: udate,
					CreateTime: ts,
					UpdateTime: ts,
				})
				if err != nil {
					hlog.Error(err)
				}
			} else if chgRt != 0 {
				rec.LatestXt = xt
				rec.UpdateDate = udate
				rec.UpdateTime = ts
				_, err = db.UpdateCurrency(&rec)
				if err != nil {
					hlog.Error(err)
				}
			}
		}
	}

}

func getLatestXt() (map[string]interface{}, error) {
	currencyApi := "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/cny.json"
	response, err := http.Get(currencyApi)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		hlog.Error("decode response err:", err)
		return nil, err
	}

	return reply, nil
}
