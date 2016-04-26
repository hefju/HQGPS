package blog
import (
	"fmt"
	log "github.com/YoungPioneers/blog4go"
	"os"
//	"time"
)
func init()  {
	hook := new(MyHook)

	err := log.NewWriterFromConfigAsFile("config.xml")
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// optionally define your logging hook
	log.SetHook(hook)
	log.SetHookLevel(log.INFO)
	// optionally set output colored
	log.SetColored(true)
	defer log.Close()
	//log.Info("blog4go init")

//	log.Debug("Debug")
//	log.Trace("Trace")
//	log.Info("Info")
//	log.Warn("Warn")
//	log.Error("Error")
//	log.Critical("Critical")

	fmt.Println("blog4go init")
}
type MyHook struct {
	something string
}

func (self *MyHook) Fire(level log.Level, message ...interface{}) {
	fmt.Println(message)
}


