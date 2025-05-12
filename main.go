package main

import (
	"coupon-system/db"
	cache "coupon-system/redis"
	"coupon-system/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("DB Initialization Begin() ..........")
	godotenv.Load()
	db.Connect()
	cache.Connect()
	fmt.Println("DB Initialization Done() ..........")
}

func main() {
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(fmt.Sprintf("%v", os.Getenv("GO_PORT")))
}
