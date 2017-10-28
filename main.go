package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.AddFunc("@every 1s", func() { fmt.Println("Every second") })
	c.AddFunc("@every 3s", func() { fmt.Println("Every 3 second") })
	c.AddFunc("@every 1s", DefaultDeliverJob)
	c.Start()
	// Funcs are invoked in their own goroutine, asynchronously.
	djob := NewDeliverJob("bj", "sz", "niurou", "100.00")

	c.AddJob("@every 1s", djob)
	// c.AddJob("@every 1s", cron.FuncJob(DefaultDeliverJob))
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	// Inspect the cron job entries' next and previous run times.
	// inspect(c.Entries())

	defer c.Stop()
	wg.Wait()

}
