package viper

type Config struct {
	App       *App       `yaml:"App"`
	DB        *DB        `yaml:"DB"`
	Redis     *Redis     `yaml:"Redis"`
	SdService *SdService `yaml:"SdService"`
	Log       *Log       `yaml:"Log"`
	Aws       *Aws       `yaml:"Aws"`
}

type App struct {
	HostPorts          string `yaml:"HostPorts"`          // The address and port the service listens on
	Domain             string `yaml:"Domain"`             // The domain of service
	BaseUrl            string `yaml:"BaseUrl"`            // The base url
	AllowOrigins       string `yaml:"AllowOrigins"`       // Custom domain detection with lower priority than AllowOrigins
	MaxRequestBodySize int    `yaml:"MaxRequestBodySize"` // Maximum request body size
}

type DB struct {
	MysqlDSN string `yaml:"MysqlDSN"` // Mysql Dsn
}

type Redis struct {
	Addr               string `yaml:"Addr"`               // Redis service address and port
	Password           string `yaml:"Password"`           // Passord
	Db                 int    `yaml:"Db"`                 // Database no
	EncodeLockSecond   int    `yaml:"EncodeLockSecond"`   // Encryption lock current limit interval
	DecodeLockSecond   int    `yaml:"DecodeLockSecond"`   // Decryption lock current limit interval
	DriftLockSecond    int    `yaml:"DriftLockSecond"`    // Drift lock current limit interval
	CompressLockSecond int    `yaml:"CompressLockSecond"` // Compress lock current limit interval
	DriftLimit         int    `yaml:"DriftLimit"`         // Limit on the number of drifting letter caches
}

type SdService struct {
	BaseUrl string `yaml:"BaseUrl"` // Stable diffusion Service base url
}

type Log struct {
	LogPath  string `yaml:"LogPath"`
	LogLevel int    `yaml:"LogLevel"`
}

type Aws struct {
	AccessKey     string `yaml:"AccessKey"`
	SecretKey     string `yaml:"SecretKey"`
	Region        string `yaml:"Region"`
	IdleTime      int64  `yaml:"IdleTime"`
	AutoStartStop bool   `yaml:"AutoStartTop"`
	SesSender     string `yaml:"SesSender"`
}
