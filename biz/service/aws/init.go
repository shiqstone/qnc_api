package service

import "qnc/biz/mw/viper"

func Init() {
	config := viper.Conf.Aws
	accessKey = config.AccessKey
	secretKey = config.SecretKey

	region = config.Region
	instanceId = "i-0dbf23e4a46fa2119"

	sender = config.SesSender
}
