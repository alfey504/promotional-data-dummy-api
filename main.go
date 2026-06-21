package main

import (
	"fmt"

	"dummyretaildata.com/dummydataserver/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	setServer()
}

func setServer() {
	router := gin.Default()

	// err := router.SetTrustedProxies([]string{"192.168.220.33"})
	// if err != nil {
	// 	panic(err)
	// }

	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", controllers.GetProductController)
		v1.GET("/skus", controllers.GetSKUController)
		v1.GET("/customers", controllers.GetCustomersController)
		v1.GET("/bundles", controllers.GetBundlesController)
		v1.GET("/promotions", controllers.GetPromotionControllers)
	}

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("error running the server Error: %s", err.Error())
		panic(err)
	}
}
