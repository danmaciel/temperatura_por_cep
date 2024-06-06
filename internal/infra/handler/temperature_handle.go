package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/danmaciel/temperatura_por_cep/internal/dto"
	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
)

func TemperatureHandle(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	clientHttp := httpClient.NewHttpClient()

	cepRules := rules.NewCepRules(clientHttp)

	if !cepRules.IsCepValid(cep) {
		WriteResponse(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	var cepModel entity.ViaCep
	err := cepRules.Exec(cep, &cepModel)

	if err != nil {
		WriteResponse(w, http.StatusNotFound, "can not find zipcode")
		return
	}

	key := os.Getenv("WEATHER_API_KEY")

	if key == "" {
		WriteResponse(w, http.StatusInternalServerError, "weather api key not found")
		return
	}

	weatherUseCase := rules.NewWeatherUseCase(clientHttp)

	var dto dto.OutDto
	errWeather := weatherUseCase.Exec(key, cepModel.City, &dto)

	if errWeather != nil {
		WriteResponse(w, http.StatusNotFound, "weather data not found")
		return
	}

	result, errOnJson := json.Marshal(dto)

	if errOnJson != nil {
		WriteResponse(w, http.StatusInternalServerError, "error on generate json from data")
		return
	}

	GetResponseHeader(w)
	w.Write(result)
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func GetResponseHeader(w http.ResponseWriter) {
	w.Header().Add("status-code", "200")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("charset", "utf-8")
}
