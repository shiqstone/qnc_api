package redis

import (
	"encoding/json"
	"qnc/biz/model/image"
)

const (
	queueSuffix = ":queue"
)

type (
	Queue struct{}
)

type ImageProcessRequestData struct {
	UserId      int64
	OrderId     int64
	InputImgStr string
	W           int
	H           int
	Cords       []image.Coordinate
	Prompt      string
	Type        string
}

func (q Queue) Enqueue(key string, data ImageProcessRequestData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rdbQueue.LPush(key+queueSuffix, jsonData).Result()
	return err
}

func (q Queue) Dequeue(key string) (ImageProcessRequestData, error) {
	result, err := rdbQueue.BRPop(0, key+queueSuffix).Result()
	if err != nil {
		return ImageProcessRequestData{}, err
	}

	var data ImageProcessRequestData
	err = json.Unmarshal([]byte(result[1]), &data)
	return data, err
}

func (q Queue) ReEnqueue(key string, data ImageProcessRequestData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rdbQueue.RPush(key+queueSuffix, jsonData).Result()
	return err
}
