package App
import (
	"fmt"
	log"github.com/YoungPioneers/blog4go"
	"os"
//	"time"
)


func init()  {
	hook := new(MyHook)

	err := log.NewWriterFromConfigAsFile("config.xml")
	//	err := log.NewFileWriter("./", true)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// optionally define your logging hook
	log.SetHook(hook)
	log.SetHookLevel(log.INFO)
	// optionally set output colored
	log.SetColored(true)
	//defer log.Close()
	//log.Info("blog4go init")

		log.Debug("Debug")
		log.Trace("Trace")
		log.Info("Info")
		log.Warn("Warn")
		log.Error("Error")
		log.Critical("Critical")
	log.Debug("blog4go init")
	//fmt.Println("blog4go init")
}

// Debug static function for Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}
// Trace static function for Trace
func Trace(args ...interface{}) {
	log.Trace(args...)
}
// Info static function for Info
func Info(args ...interface{}) {
	log.Info(args...)
}
// Warn static function for Warn
func Warn(args ...interface{}) {
	log.Warn(args...)
}
// Error static function for Error
func Error(args ...interface{}) {
	log.Error(args...)
}
// Critical static function for Critical
func Critical(args ...interface{}) {
	log.Critical(args...)
}






type MyHook struct {
	something string
}

func (self *MyHook) Fire(level log.Level, message ...interface{}) {
	fmt.Println(message)
}
