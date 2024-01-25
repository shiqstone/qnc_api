package service

import (
	"context"
	"fmt"
	"math/rand"
	"qnc/biz/dal/db"
	"qnc/biz/model/user"
	"qnc/pkg/errno"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type AccountService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewAccountService create user service
func NewAccountService(ctx context.Context, c *app.RequestContext) *AccountService {
	return &AccountService{ctx: ctx, c: c}
}

// account increase return user balance.
func (s *AccountService) Increase(req *user.FundsRequest) (balance float64, err error) {
	//query account
	rec, err := db.QueryBalanceById(req.UserId)
	if err != nil {
		return 0, err
	}

	ts := time.Now().Unix()
	//incre balance
	rec.ID = req.UserId
	balance = rec.Coin + req.Amount
	rec.Coin = balance
	rec.UpdateTime = ts
	rec, err = db.UpdateUser(rec)
	if err != nil {
		return 0, err
	}

	//add account log
	_, err = db.CreateCoinDetail(&db.CoinDetail{
		UserId:     req.UserId,
		Type:       req.EventType,
		EventId:    req.OrderId,
		Incoming:   req.Amount,
		Expend:     0.0,
		Balance:    balance,
		Remark:     req.Remark,
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		return 0, err
	}
	return rec.Coin, nil
}

// account decrease return user balance.
func (s *AccountService) Decrease(req *user.FundsRequest) (balance float64, err error) {
	//query account
	rec, err := db.QueryBalanceById(req.UserId)
	if err != nil {
		return 0, err
	}

	if rec.Coin < req.Amount {
		return rec.Coin, errno.RecordNotExistErr
	}

	ts := time.Now().Unix()
	//incre balance
	rec.ID = req.UserId
	balance = rec.Coin - req.Amount
	rec.Coin = balance
	rec.UpdateTime = ts
	rec, err = db.UpdateUser(rec)
	if err != nil {
		return 0, err
	}

	//add account log
	_, err = db.CreateCoinDetail(&db.CoinDetail{
		UserId:     req.UserId,
		Type:       req.EventType,
		EventId:    req.OrderId,
		Incoming:   0.0,
		Expend:     req.Amount,
		Balance:    balance,
		Remark:     req.Remark,
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		return 0, err
	}
	return rec.Coin, nil
}

func (s *AccountService) CreateAccountTopup(req *user.AccountTopupRequest) (depositId string, err error) {
	depositId = genId()

	if req.Currency == "" {
		req.Currency = "USD"
	}

	ts := time.Now().Unix()
	// add deposit record
	_, err = db.CreateDeposit(&db.Deposit{
		UserId:     req.UserId,
		DepositId:  depositId,
		Amount:     req.Amount,
		Status:     user.PAY_STATUS_INIT,
		Currency:   req.Currency,
		Ip:         req.Ip,
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		hlog.Errorf("err = ", err)
		return "", err
	}
	return depositId, nil
}

func genId() string {
	// YYYYMMDDHHmmSSmi
	ts := time.Now().Format("20060102150405,000")
	ts = strings.Replace(ts, ",", "", -1)
	max := 9999
	min := 1000
	rt := rand.Intn(max-min) + min
	return fmt.Sprintf("%s%d", ts, rt)
}

func (s *AccountService) GetPayStatus(req *user.PayStatusRequest) (resp *user.PayStatusResponse, err error) {
	deposit, err := db.QueryByDepositId(req.DepositId)
	if err != nil {
		return nil, err
	}

	resp.DepositId = deposit.DepositId
	resp.Status = deposit.Status

	return resp, nil
}

func (s *AccountService) UpdatePayStatus(req *user.PayStatusRequest) (err error) {
	deposit, err := db.QueryByDepositId(req.DepositId)
	if err != nil {
		return err
	}

	if req.Amount != deposit.Amount {
		return errno.AmountNotMatchErr
	}
	if deposit.Status == user.PAY_STATUS_SUCCESS || deposit.Status == user.PAY_STATUS_FALID {
		return errno.TransactionDoneErr
	}

	ts := time.Now().Unix()
	// update deposit record
	if req.Result == "success" {
		deposit.Status = user.PAY_STATUS_SUCCESS
	} else {
		deposit.Status = user.PAY_STATUS_FALID
	}
	deposit.FinishTime = ts
	deposit.UpdateTime = ts
	// deposit.PayChannel = req.PayChannel
	deposit.BankTradeNo = req.PayId
	deposit.Ext = req.Remark
	_, err = db.UpdateDeposit(deposit)
	if err != nil {
		return err
	}

	if req.Result == "success" {
		//increase account
		_, err = s.Increase(&user.FundsRequest{
			UserId:    deposit.UserId,
			Amount:    req.Amount,
			OrderId:   deposit.ID,
			EventType: user.TYPE_RECHARGE,
		})
		if err != nil {
			return err
		}
	}

	return nil
}