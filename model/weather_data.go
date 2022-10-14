package model

type WeatherData struct {
	Status WeatherStatus `json:"status"`
}

type WeatherStatus struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

const (
	statusBahaya = "bahaya"
	statusSiaga  = "siaga"
	statusAman   = "aman"
)

func (weatherData *WeatherData) GetWaterStatus() string {
	waterScale := weatherData.Status.Water
	if waterScale > 8 {
		return statusBahaya
	}

	if waterScale <= 8 && waterScale >= 6 {
		return statusSiaga
	}

	return statusAman
}

func (weatherData *WeatherData) GetWindStatus() string {
	windScale := weatherData.Status.Water

	if windScale > 15 {
		return statusBahaya
	}

	if windScale <= 15 && windScale >= 7 {
		return statusSiaga
	}

	return statusAman
}
