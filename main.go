package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main(){
    router  := gin.Default()
    router.GET("/GetLocation", GetLocation)
    router.Run(":8089") // listen and server on 0.0.0.0:8080
    fmt.Println("end")
}
func GetLocation(c *gin.Context){
    carNum:=c.Query("vehicle")
    fmt.Println(carNum)
    c.JSON(200, gin.H{
        "carNum":carNum,
        "lon": "113.1524",
        "lat": "23.0822",
        "gpstime":"2016.4.21 19:32",
"speed":"16",
    })
}