# CEP GCR API

## Visão Geral
Serviço HTTP em Go que recebe um CEP via rota `/cep/{cep}`, valida o formato, busca endereço na ViaCEP e consulta a WeatherAPI pela localidade retornada. Responde JSON com temperaturas em Celsius, Fahrenheit e Kelvin, permitindo consultas simples de clima por CEP brasileiro.

## Pré-requisitos
- Go 1.24+ instalado e disponível em `PATH`.
- Docker (opcional) caso deseje buildar e executar o container publicado em `Dockerfile`.
- Chaves de API: o código inclui uma chave de demonstração da WeatherAPI; substitua `WeatherApiPath` em `infra/service/WeatherApi.go` por sua chave real em produção.

## Instalação
```bash
git clone git@github.com:brunobdl97/cepGCR.git
cd cepGCR
go mod download
```
Executar os testes garante que tudo foi instalado corretamente:
```bash
go test ./...
```

## Inicialização do Servidor
### Go local
```bash
go run ./cmd
```
O serviço escuta em `:8080`. Exemplo de chamada: `curl http://localhost:8080/cep/01001000`.

### Docker
```bash
docker build -t cepgcr .
docker run -p 8080:8080 --env WEATHER_API_KEY=<sua-chave> cepgcr
```
Se alterar a chave via variável de ambiente, adapte o código para ler esse valor antes do build.

### Através do Google Cloud Run
```bash
curl https://cepgcr-834760883240.europe-west1.run.app/cep/30290140
```
Alterar o CEP conforme desejar.

## Observações Importantes
- Fluxo depende de conectividade externa com ViaCEP e WeatherAPI; trate limites de rate ao implantar.
- Inputs inválidos retornam `422`, CEPs inexistentes retornam `404` com mensagem `cannot find zipcode`.
- Utilize `go fmt ./...` e `go vet ./...` antes de abrir PRs; testes vivem em `internal/*.go`.
- A API_KEY foi implementada de forma hardcoded no código apenas para fins didáticos, o ideal é utilizar variáveis de ambientes ou até mesmo um secret manager em cloud.

## Estrutura do Projeto
```
cmd/            # main e wiring HTTP
internal/       # regras de domínio (CEP, Weather, handler GetWeather)
internal/dto/   # contratos externos (ViaCepResponse, WeatherApiResponse)
infra/service/  # integrações HTTP com APIs
Dockerfile      # build multi-stage para container minimalista
```
