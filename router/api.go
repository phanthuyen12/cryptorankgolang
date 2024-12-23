// /router/api.go
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	// Định nghĩa các route cho API
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "thanhthuyen frodm thuyen!",
		})
	})

	// Các API khác có thể được thêm vào đây
}
