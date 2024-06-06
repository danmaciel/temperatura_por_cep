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

	var cepModel entity.ViaCep
	err := cepRules.Exec(cep, &cepModel)

	if err != nil {
		WriteResponse(w, err.Code, err.Message)
		return
	}

	weatherUseCase := rules.NewWeatherUseCase(clientHttp)
	key := os.Getenv("WEATHER_API_KEY")

	var dto dto.OutDto
	errWeather := weatherUseCase.Exec(key, cepModel.City, &dto)

	if errWeather != nil {
		WriteResponse(w, errWeather.Code, errWeather.Message)
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
