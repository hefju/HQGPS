package App
import (
//	"fmt"
	//log"github.com/YoungPioneers/blog4go"
//	"os"
//	"time"
	//github.com/iamthemuffinman/logsip
	//github.com/donnie4w/go-logger
//	"github.com/donnie4w/go-logger/logger"
//	"github.com/hefju/BackupCopyFiles/tools/setting"
	"github.com/donnie4w/go-logger/logger"
	"github.com/hefju/HQGPS/tools/setting"

)

//var JLog logger
func AppInit()  {

	setting.LoadProfile()// loadConfig()//读取配置文件
	LogInit()
//	logger.Debug("app init")
}
func LogInit(){
	logger.SetConsole(false)
	path:=setting.AppConfig.LogPath//="D:/Programs/HSGPS/log"//
	//log.Println("setting:",setting.AppConfig.LogPath)
	logger.SetRollingDaily(path,"test.log")
	logger.SetLevel(logger.DEBUG)
//	log.Println("app init")
}




