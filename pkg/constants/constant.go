package constants

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
