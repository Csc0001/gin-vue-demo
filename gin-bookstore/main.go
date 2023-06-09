package main

import (
	"fmt"
	"gin-vue-bookStore/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)


func main() {
	InitConfig()
	r := gin.Default()
	db := common.InitDB()
	defer db.Close()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	fmt.Println(port)
	if port != ""{
		panic(r.Run(":"+port))
	}
	panic(r.Run())
}

func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}





