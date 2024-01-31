package utils

import (
	"qnc/pkg/constants"
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	uuidFunc := uuid.New()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

func GenClientId() string {
	raw := []byte(constants.GlobalSetting.LocalHost + ":" + constants.CommonSetting.RPCPort)
	str, err := Encrypt(raw, []byte(constants.CommonSetting.CryptoKey))
	if err != nil {
		panic(err)
	}

	return str
}

func IsAddrLocal(host string, port string) bool {
	return host == constants.GlobalSetting.LocalHost && port == constants.CommonSetting.RPCPort
}

func IsCluster() bool {
	return constants.CommonSetting.Cluster
}

func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	addr, err = Decrypt(clientId, []byte(constants.CommonSetting.CryptoKey))
	if err != nil {
		return
	}

	// host, port, err = ParseRedisAddrValue(addr)
	// if err != nil {
	// 	return
	// }

	isLocal = IsAddrLocal(host, port)
	return
}

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}
