package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"weather/helper"
)

func Home(filePath string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weatherData, err := helper.GetWeatherData(filePath)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.HTML(http.StatusOK, "homepage.tmpl", gin.H{
			"water":       weatherData.Status.Water,
			"wind":        weatherData.Status.Wind,
			"windStatus":  weatherData.GetWindStatus(),
			"waterStatus": weatherData.GetWaterStatus(),
			"lastUpdate":  time.Now().Format(time.RFC850),
		})
	}
}
