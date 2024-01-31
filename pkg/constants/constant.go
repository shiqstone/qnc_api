package constants

import (
	"net"
	"sync"
)

// connection information
const (
	MySQLDefaultDSN = "root:root@tcp(127.0.0.1:3306)/qnc_db?charset=utf8&parseTime=True&loc=Local"

	MinioEndPoint        = "127.0.0.1:3306"
	MinioAccessKeyID     = "root"
	MinioSecretAccessKey = ""
	MiniouseSSL          = false

	RedisAddr     = "localhost:6379"
	RedisPassword = ""
)

// constants in the project
const (
	UserTableName       = "qnc_user"
	CoinDetailTableName = "qnc_coin_detail"
	DepositTableName    = "qnc_deposit"
	OrderTableName      = "qnc_order"
	KvTableName         = "qnc_kv"
	// FavoritesTableName = "likes"
	// CommentTableName   = "comments"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	// MinioVideoBucketName = "videobucket"
	// MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)

type commonConf struct {
	HttpPort  string
	RPCPort   string
	Cluster   bool
	CryptoKey string
}

type global struct {
	LocalHost      string //本机内网IP
	ServerList     map[string]string
	ServerListLock sync.RWMutex
}

var CommonSetting = &commonConf{
	HttpPort:  "8899",
	RPCPort:   "7000",
	Cluster:   false,
	CryptoKey: "Adba723b7fe06819",
}

var GlobalSetting = &global{
	LocalHost:  getIntranetIp(),
	ServerList: make(map[string]string),
}

func getIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
