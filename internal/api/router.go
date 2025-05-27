package api

import "github.com/gin-gonic/gin"

// فرض بر اینه که GetStatusHandler تو همین پکیج api تعریف شده
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// مسیر وضعیت سرویس‌ها
	router.GET("/status", GetStatusHandler)

	return router
}
