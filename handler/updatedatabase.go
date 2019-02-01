package handler

import (
	"os"
	"time"

	"github.com/golang/glog"
)

// UpdateDatabase will update the latest pipelines detail and status
func UpdateDatabase() {
	// Read token environment variable
	token, ok := os.LookupEnv(token)
	if !ok {
		glog.Fatalf("TOKEN environment variable required")
	}
	// loop will iterate at every 6000 seconds
	tick := time.Tick(20000 * time.Millisecond)
	for range tick {
		BuildData(token)
		// Trigger GkeData function for update GKE related data to database
		go GkeData(token)
		// 	// Trigger GkeData function for update EKS related data to database
		// 	go EksData(token)
		// 	// Trigger GkeData function for update AKS related data to database
		// 	go AksData(token)
	}
}
