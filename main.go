package main

import "booked/config"

func main() {
	manager := &config.Manager{
		Debug: true,
		Host:  "",
		Port:  "8100",
		ENV:   "dev",
	}
	manager.InitDB()
	manager.Run()
}
