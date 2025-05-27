package checker

import (
	"net/http"
	"time"

	"github.com/hamed-fahimi/health-checker/internal/db"
)

func CheckURL(name, url string) db.CheckResult {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	result := db.CheckResult{
		ServiceName:  name,
		URL:          url,
		CheckedAt:    time.Now(),
		ResponseTime: duration.String(),
	}

	if err != nil {
		result.Status = "down"
		result.StatusCode = 0
	} else {
		defer resp.Body.Close()
		result.StatusCode = resp.StatusCode
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			result.Status = "up"
		} else {
			result.Status = "down"
		}
	}

	// ذخیره در دیتابیس
	db.DB.Create(&result)

	return result
}
