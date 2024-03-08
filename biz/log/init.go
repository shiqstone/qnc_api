package log

import (
	"io"
	"log"
	"os"
	"path"
	"qnc/biz/mw/viper"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Init() {
	defer initLog().Close()
	initLog()
}

func initLog() *os.File {
	config := viper.Conf.Log
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(config.LogPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil
		}
	}

	hlog.SetLevel((hlog.Level)(config.LogLevel))
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetOutput(fileWriter)
	return f
}
