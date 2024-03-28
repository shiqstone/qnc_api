package currency

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"qnc/biz/dal/db"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	aws "qnc/biz/service/aws"
)

func Tracking() {
	defer func() {
		if err := recover(); err != nil {
			// Log the error
			hlog.Errorf("Recovered from panic in Tracking function: %v", err)
		}
	}()

	ret, err := getLatestXt()
	if err != nil {
		hlog.Error(err)
		return
	}
	udate, ok := ret["date"].(string)
	if !ok {
		hlog.Error("Invalid date format")
		return
	}
	xts, ok := ret["cny"].(map[string]interface{})
	if !ok {
		hlog.Error("Invalid currency format")
		return
	}

	curRec, err := db.QueryByLatestList()
	if err != nil {
		hlog.Error("query db record err, ", err)
		return
	}

	changeMsg := ""
	ts := time.Now().Unix()
	for _, rec := range *curRec {
		cur := strings.ToLower(rec.Currency)
		if val, ok := xts[cur]; ok {
			changeMsg, err = updateRecord(val, rec, udate, ts, changeMsg)
			if err != nil {
				hlog.Error(err)
				return
			}
		}
	}

	if changeMsg != "" {
		recipient := "shiqstone@proton.me"
		subject := "Currency rate changed over limit"
		html := "<p>" + strings.Replace(changeMsg, "\r\n", "</p><p>", -1) + "</p>"
		err = aws.SendEmail(recipient, subject, changeMsg, html)
		if err != nil {
			hlog.Error(err)
			return
		}
	}

}

func updateRecord(val interface{}, rec db.CurrencyRecord, udate string, ts int64, changeMsg string) (string, error) {
	xt := val.(float64)
	chgRt := math.Abs(rec.BaseXt-xt) / rec.BaseXt
	if chgRt > 0.05 {
		// rate change over 5%
		_, err := db.CreateCurrency(&db.CurrencyRecord{
			Currency:   rec.Currency,
			BaseXt:     xt,
			LatestXt:   xt,
			UpdateDate: udate,
			CreateTime: ts,
			UpdateTime: ts,
		})
		if err != nil {
			return "", err
		}

		changeMsg += fmt.Sprintf("Current exchange rate CNY to %s has changed from %f to %f, change over %f%%. \r\n", rec.Currency, rec.BaseXt, rec.LatestXt, chgRt*100)
	} else if chgRt != 0 {
		rec.LatestXt = xt
		rec.UpdateDate = udate
		rec.UpdateTime = ts
		_, err := db.UpdateCurrency(&rec)
		if err != nil {
			return "", err
		}
	}
	return changeMsg, nil
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
