
APP_NAME := erpnet
BIN_DIR  := bin
MAIN_PATH := ./cmd

build: ## Compila o binário
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_PATH)

run: ## Roda a aplicação localmente
	go run ./...

test:
	go test ./... -v

generate:
	go generate ./...

test-race: ## Roda testes com detector de data race
	go test -race -count=1 ./...

vet: ## Roda go vet
	go vet ./...

secrets: ## Escaneia segredos vazados (gitleaks)
	gitleaks detect --source . -v

check: vet test-race 