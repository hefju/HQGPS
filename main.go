package main

import (
    "github.com/hefju/HQGPS/tools/setting"
	"github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "encoding/base64"
    "bytes"
    "time"
    "github.com/hefju/HQGPS/model"
    "github.com/hefju/HQGPS/App"

	"github.com/donnie4w/go-logger/logger"
    "fmt"
)

func main() {
	App.AppInit()
    model.Init2()

    gin.SetMode(gin.ReleaseMode)
	router := gin.New()
    router.GET("/Chinese",Chinese)
    router.GET("/UpdateConfig",ReloadSetting)

    router.GET("/GetLocationTest",GetLocation)

    authorized := router.Group("/")
    authorized.Use(AuthRequired())
    {
        authorized.POST("/GetLocation", GetLocation)
        authorized.POST("/GetLocationCN", GetLocationCN)
}
  //  App.Debug("lisent at 8089")
	logger.Debug("lisent at 8089")
	logger.Fatal(router.Run(":8089"))
}

func ReloadSetting(c *gin.Context){
   // log.Println("ReloadSetting")
    setting.LoadProfile()
    c.JSON(200,time.Now().Format("2006-01-02 15:04:05")+ " ReloadSetting")
}
func AuthRequired()gin.HandlerFunc{
    return func(c *gin.Context) {
        basicAuthPrefix := "Basic "
        // 获取 request header
        auth :=  c.Request.Header.Get("Authorization")//auth: Basic Zm9vOmJhcjU=
        // 如果是 http basic auth
        if strings.HasPrefix(auth, basicAuthPrefix) {
            // 解码认证信息
            payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])//payload: foo:bar5
            if err == nil {
                pair := bytes.SplitN(payload, []byte(":"), 2)
                user:= []byte(setting.AppConfig.User)
                passwd := []byte(setting.AppConfig.Pwd)
                if len(pair) == 2 && bytes.Equal(pair[0], user) && bytes.Equal(pair[1], passwd) {
                    c.Set("auth", "abcd") //c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
                    return
                }
            }
        }
        // 认证失败，提示 401 Unauthorized
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        c.Abort()
    }
}
func GetLocation(c *gin.Context) {
	//carNum := c.Query("vehicle")
	carNum :=c.PostForm("vehicle")  //c.Param("vehicle")
    fmt.Println("GetLocation Param [vehicle] ",carNum)
    gps:=  model.GetGps(carNum)
    c.JSON(200, gps)
}
func GetLocationCN(c *gin.Context) {
    carNum := c.Query("vehicle")
    fmt.Println(carNum)
    gps:=  model.GetGpsCN(carNum)
    c.JSON(200, gps)
}
//http://www.ruanyifeng.com/blog/2010/02/url_encoding.html
func Chinese(c *gin.Context){
    ch:=c.Query("ch")
    fmt.Println(ch+"end")
    c.JSON(200,ch)
}
