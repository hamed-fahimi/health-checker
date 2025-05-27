package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hamed-fahimi/health-checker/internal/checker"
)

// Handler تابعی برای هندل کردن درخواست GET /status
func GetStatusHandler(c *gin.Context) {
	statuses := checker.GetStatuses()
	c.JSON(http.StatusOK, statuses)
}
