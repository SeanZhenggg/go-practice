package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func main() {
	c := cron.New()

	_, err := c.AddFunc("@every 1s", func() {
		f, err := os.Create("test.log")
		if err != nil {
			return
		}
		defer f.Close()

		log.New(f, "", log.LstdFlags)
	})
	if err != nil {
		return
	}

	c.Start()
	select {}
}
