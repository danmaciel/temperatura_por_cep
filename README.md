# API em Golang que retorna a temperatura.

Sistema que recebe um CEP e retorna a temperatura em Celsius, Fahrenheit e Kelvin, por meio de um CEP de entrada.

No link abaixo, a aplicação está online e pronta para receber dados, onde é necessária a troca de SEU_CEP pelo CEP de pesquisa.

https://temperatura-por-cep-zufalhzjhq-uw.a.run.app/temperature/SEU_CEP

Por exemplo,

https://temperatura-por-cep-zufalhzjhq-uw.a.run.app/temperature/86010180

Onde o tipo de retornmo e no formato JSON e na seguinte estrutura:
```
{"temp_c":24,"temp_f":75.2,"temp_k":297}
```

Onde:

```
temp_c = Celsius
temp_f = Farenheit
temp_k = Kelvin
```

## O que é necessário para rodar

```
Ter instalado o Go na sua máquina.
Ter instalado o Docker na sua máquina.
Cadastro em https://www.weatherapi.com para ter a API key para utilização na aplicação.
```


## Como rodar

Para rodar localmente, pode-se fazer de três formas:

### Gerando uma imagem docker

```
docker build -t temperature-app .
```

e para rodar, é necessário passar o environment(-e) e porta(-p).

```
docker run -e "WEATHER_API_KEY=SUA_KEY" -p 8080:8080 temperature-app
```

Ou então através do docker-compose, onde é necessário setar em environment a WEATHER_API_KEY com o uso da API key.

### Através do docker-compose

```
docker-compose up
```

### Executando o projeto diretamente.

```
go run main.go
```

Em todos os casos, a porta utilizada é a 8080.

## Testes

Os testes unitários ficam no diretorio tests na raiz da aplicação sendo necessario alterar a linha os.Setenv("WEATHER_API_KEY_TEST", ""), com a usa API key e depois executá-los:

```
go test ./tests/
```