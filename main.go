package main

import (
	"fmt"

	"dummyretaildata.com/dummydataserver/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", controllers.GetProductController)
		v1.GET("/sku", controllers.GetSKUController)
	}

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("error running the server Error: %s", err.Error())
		panic(err)
	}
}
