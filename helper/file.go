package helper

import (
	"encoding/json"
	"os"
	"weather/model"
)

func GetWeatherData(fileName string) (*model.WeatherData, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var weatherData *model.WeatherData
	err = json.Unmarshal(content, &weatherData)
	if err != nil {
		return nil, err
	}

	return weatherData, nil
}

const readWritePermission = 0644

func SaveWeatherData(date *model.WeatherData, fileName string) error {
	content, err := json.Marshal(date)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, content, readWritePermission)
	if err != nil {
		return err
	}

	return nil
}
