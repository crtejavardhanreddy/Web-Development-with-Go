package routes

import "github.com/gin-gonic/gin"

func Route() {
	r := gin.Default()
	r.GET("/restaurant", getAllRestaurnats)
}
