package setting
import (
    "encoding/json"
    "os"
    "log"
)
var AppConfig Config//系统配置
type Config struct {
    User string
    Pwd string
}
func LoadProfile(){//从配置文件读取配置
    file:="conf.json"
    r, err := os.Open(file)
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    var c Config
    err = decoder.Decode(&c)
    if err != nil {
        log.Fatalln(err)
    }
    AppConfig=c
    log.Println("load setting from conf.json.")
}