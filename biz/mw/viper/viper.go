package viper

import (
	"flag"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

var Conf *Config

// init viper
func Init() {
	var configFile = flag.String("f", "config.yaml", "set config file which viper will loading.")
	flag.Parse()

	// hlog.Debug(*configFile)
	viper.AddConfigPath(".")
	viper.SetConfigFile(*configFile)

	if err := viper.ReadInConfig(); err != nil {
		hlog.Errorf("[Viper] ReadInConfig failed, err: %v", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		hlog.Errorf("[Viper] Unmarshal failed, err: %v", err)
	}

	hlog.Infof("[Viper] Conf.App: %#v", Conf.App)
	// hlog.Infof("[Viper] Conf.DB: %#v", Conf.DB)
	// hlog.Infof("[Viper] Conf.Redis: %#v", Conf.Redis)
}
