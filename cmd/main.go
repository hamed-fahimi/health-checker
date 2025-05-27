package main

import (
	"time"

	"github.com/hamed-fahimi/health-checker/internal/api"
	"github.com/hamed-fahimi/health-checker/internal/checker"
	"github.com/hamed-fahimi/health-checker/internal/db"
)

func main() {
	db.Init()
	checker.StartPeriodicCheck(5 * time.Minute) // اجرا هر ۵ دقیقه

	r := api.SetupRouter()
	r.Run(":8080")
}
