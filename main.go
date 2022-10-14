package main

import (
	"fmt"
	"os"
	"weather/router"
	"weather/worker"
)

const weatherFileName = "weather.json"

func main() {
	workPath, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", workPath, weatherFileName)

	var endSignal = make(chan int, 1)
	go worker.AutoUpdateData(endSignal, filePath)
	r := router.StartServer(filePath)
	r.Run(":8080")

	endSignal <- 1
}
