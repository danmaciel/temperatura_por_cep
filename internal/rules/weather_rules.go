package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danmaciel/temperatura_por_cep/internal/dto"
	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	"github.com/danmaciel/temperatura_por_cep/internal/util"
)

type Weather struct {
	http *http.Client
}

func NewWeatherUseCase(h *http.Client) *Weather {
	return &Weather{
		http: h,
	}
}

func (w *Weather) Exec(key string, city string, dto *dto.OutDto) error {
	weatherUrl := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%v&q=%q", key, util.StringPrepare(city))

	resWeather, errWeather := w.http.Get(weatherUrl)

	if errWeather != nil {
		return errWeather
	}

	var weatherData entity.WeatherData
	json.NewDecoder(resWeather.Body).Decode(&weatherData)

	dto.Celsius = weatherData.Current.TempC
	dto.Fahrenheit = dto.Celsius*1.8 + 32
	dto.Kelvin = dto.Celsius + 273

	return nil

}
