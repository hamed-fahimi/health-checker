package checker

import (
	"log"
	"time"

	"github.com/hamed-fahimi/health-checker/internal/db"
)

// گرفتن همه سرویس‌هایی که باید چک بشن
func GetMonitoredServices() []db.MonitoredService {
	var services []db.MonitoredService
	db.DB.Find(&services)
	return services
}

// گرفتن وضعیت همه سرویس‌ها با چک جدید (اگر می‌خوای همیشه آخرین وضعیت ذخیره‌شده رو بدی)
// یا اگه می‌خوای چک جدید بزنی:
func GetStatuses() []db.CheckResult {
	var results []db.CheckResult
	db.DB.Order("checked_at desc").Find(&results)
	return results
}

// شروع چک دوره‌ای با استفاده از goroutine و time.Ticker
func StartPeriodicCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			log.Println("Starting periodic health check...")
			services := GetMonitoredServices()
			for _, svc := range services {
				result := CheckURL(svc.Name, svc.URL)
				log.Printf("Checked %s: status %s", svc.Name, result.Status)
			}
		}
	}()
}
