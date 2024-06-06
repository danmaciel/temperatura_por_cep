package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/danmaciel/temperatura_por_cep/internal/dto"
	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestWeatherApi(t *testing.T) {

	//os.Setenv("WEATHER_API_KEY", "SUA_API_KEY")
	os.Setenv("WEATHER_API_KEY_TEST", "")

	clientHttp := httpClient.NewHttpClient()

	cep := rules.NewCepRules(clientHttp)

	var cepModel entity.ViaCep
	err := cep.Exec("03928-040", &cepModel)

	assert.Nil(t, err)
	assert.NotNil(t, cepModel)
	assert.Equal(t, cepModel.City, "São Paulo")

	weather := rules.NewWeatherUseCase(clientHttp)

	key := os.Getenv("WEATHER_API_KEY_TEST")

	if key == "" {
		t.Error(fmt.Errorf("WEATHER_API_KEY is empty"))
		return
	}

	var dto dto.OutDto
	errWeather := weather.Exec(key, cepModel.City, &dto)

	fmt.Printf("Celsius  %v", dto.Celsius)

	assert.Nil(t, errWeather)
	assert.NotNil(t, dto)
	assert.Greater(t, dto.Celsius, float64(0))
	assert.Greater(t, dto.Fahrenheit, 0.0)
	assert.Greater(t, dto.Kelvin, 0.0)

	os.Unsetenv("WEATHER_API_KEY_TEST")
}
