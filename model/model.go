package model
import (
    "time"
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/lunny/godbc"

    "log"
    "encoding/binary"
    "cmd/go/testdata/shadow/root1/src/math"
    "bytes"
)

var engine *xorm.Engine
func init() {
    var err error
    engine, err = xorm.NewEngine("sqlite3", "./test.db")
    //engine.ShowSQL = true
    engine.SetMapper(core.SameMapper{})
    if err != nil {
        log.Println(err)
    }

}

type Gps struct {
    CarNum string
    Lon, Lat float64
    Gpstime string//time.Time
    Speed int32
}

func GetGps(vehicle string)*Gps{
    sql := "exec GetGps '"+vehicle+"'"//"select * from userinfo"
    results, err := engine.Query(sql)
    if err != nil {
        log.Println(err)
    }
    if len(results)>0{
        one:=results[0]
        gps:=&Gps{}
        gps.CarNum=string(one["CarNum"])
        gps.Lon=Float64frombytes(one["Lon"])
        gps.Lat=Float64frombytes(one["Lat"])

        gps.Gpstime=string(one["Gpstime"])
        gps.Speed=read_int32(one["Speed"])
    }
}
func Float64frombytes(bytes []byte) float64 {
    bits := binary.LittleEndian.Uint64(bytes)
    float := math.Float64frombits(bits)
    return float
}
func Float64bytes(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, bits)
    return bytes
}
func read_int32(data []byte) (ret int32) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}
func Int64ToBytes(i int64) []byte {
    var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}

func BytesToInt64(buf []byte) int64 {
    return int64(binary.BigEndian.Uint64(buf))
}


