package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"io"
	"net/http"
	"qnc/biz/dal/db"
	mimg "qnc/biz/model/image"
	"qnc/biz/model/order"
	"qnc/biz/model/user"
	"qnc/biz/mw/redis"
	service "qnc/biz/service/user"
	"qnc/pkg/errno"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func (s *ImageService) ProcessImageUd(req *mimg.ImageUdRequest) (resp *mimg.ImageUdResponse, err error) {
	cuid, exists := s.c.Get("current_user_id")
	if !exists {
		return nil, errno.AuthorizationFailedErr
	}
	req.UserId = cuid.(int64)

	//prepare param
	base64Content, width, height, coordinates, err := prepareParam(req)
	if err != nil {
		return nil, err
	}

	//TODO get price config
	point := 2.0
	pname := "UD image"
	var pid int64 = 1

	ts := time.Now().Unix()
	// add order record
	orderId, err := db.CreateOrder(&db.Order{
		UserId:     req.UserId,
		ProdName:   pname,
		ProdId:     pid,
		RealCost:   point,
		BaseCost:   point,
		Status:     order.STATUS_INIT,
		Ip:         s.c.ClientIP(),
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		return nil, err
	}

	// decrease account
	var freq user.FundsRequest
	freq.UserId = req.UserId
	freq.Amount = point
	freq.EventType = user.TYPE_PAYMENT
	freq.OrderId = orderId
	balance, err := service.NewAccountService(s.ctx, s.c).Decrease(&freq)
	if err != nil {
		return nil, err
	}

	prompt := "dress"

	//enqueue
	queue := new(redis.Queue)
	err = queue.Enqueue("processImage", redis.ImageProcessRequestData{
		UserId:      req.UserId,
		OrderId:     orderId,
		InputImgStr: base64Content,
		W:           width,
		H:           height,
		Cords:       coordinates,
		Prompt:      prompt,
		Type:        "ud",
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	// _, err = queue.Dequeue("processImage")
	// if err != nil {
	// 	return nil, err
	// }

	ts = time.Now().Unix()
	db.UpdateOrder(&db.Order{
		ID:     orderId,
		Status: order.STATUS_PAYED,
		// Remark:     "",
		UpdateTime: ts,
	})

	// hlog.Debug("queue data", qdata)
	resp = new(mimg.ImageUdResponse)
	// resp.ProcessedImage = processedImg
	resp.OrderId = orderId
	resp.Balance = balance

	return resp, nil
}

func PostProcessImageUd(base64Content string, width, height int, coordinates []mimg.Coordinate, prompt string, orderId int64) (resp *mimg.ImageUdResponse, err error) {
	// // process image
	msg, processedImg, seed, err := processImage(base64Content, width, height, coordinates, prompt)
	if err != nil {
		hlog.Error("process image err:", err)
		return nil, err //errors.New("process image failed")
	}

	// hlog.Debug((res))
	ts := time.Now().Unix()
	if processedImg != "" {
		db.UpdateOrder(&db.Order{
			ID:         orderId,
			Status:     order.STATUS_SUCCESS,
			Remark:     strconv.FormatInt(seed, 10),
			UpdateTime: ts,
		})

		var resp = new(mimg.ImageUdResponse)
		resp.ProcessedImage = processedImg
		resp.OrderId = orderId
		return resp, nil
	} else {
		ts := time.Now().Unix()
		db.UpdateOrder(&db.Order{
			ID:         orderId,
			Status:     order.STATUS_FALID,
			Remark:     msg,
			UpdateTime: ts,
		})
		//TODO refund

		return nil, errors.New(msg)
	}
}

func prepareParam(req *mimg.ImageUdRequest) (base64Content string, width, height int, coordinates []mimg.Coordinate, err error) {
	f, err := req.FileHeader.Open()
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	defer f.Close()

	// 读取文件到字节数组
	fileRaw, err := io.ReadAll(f)
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	contentType := http.DetectContentType(fileRaw)

	base64Content = base64.StdEncoding.EncodeToString(fileRaw)
	base64Content = "data:" + contentType + ";base64," + base64Content

	f, err = req.FileHeader.Open()
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		hlog.Errorf("err = ", err)
		return "", 0, 0, nil, err
	}

	b := img.Bounds()
	width = b.Max.X
	height = b.Max.Y

	// var coordinates []Coordinate
	err = json.Unmarshal([]byte(req.Pos), &coordinates)
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	return
}
