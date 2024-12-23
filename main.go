// main.go
package main

import (
	"my1/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo Gin router
	r := gin.Default()

	// Đăng ký các route cho API
	router.ApiRoutes(r)

	// Đăng ký các route cho web
	// router.WebRoutes(r)

	// Khởi chạy server trên cổng 8080
	r.Run(":8080")
}
