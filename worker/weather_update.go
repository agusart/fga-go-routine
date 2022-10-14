package worker

import (
	"fmt"
	"math/rand"
	"time"
	"weather/helper"
)

const autoUpdateSecondDuration = 15

func AutoUpdateData(endProcess <-chan int, filePath string) {
	for {
		select {
		case <-endProcess:
			fmt.Println("worker stopped")
			return
		default:
			updateWeatherData(filePath)
			time.Sleep(autoUpdateSecondDuration * time.Second)
		}
	}
}

func updateWeatherData(filePath string) {

	data, err := helper.GetWeatherData(filePath)
	if err != nil {
		fmt.Sprintln(err)
	}

	data.Status.Wind = getRandData()
	data.Status.Water = getRandData()

	fmt.Printf("wind: %d, water:%d \n", data.Status.Wind, data.Status.Water)
	err = helper.SaveWeatherData(data, filePath)
	if err != nil {
		fmt.Println(err)
	}
}

func getRandData() int {
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
