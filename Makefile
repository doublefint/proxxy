APP=proxxy
BIN_FOLDER=bin
IGNORED_FOLDER=.ignore
COVERAGE_FILE=$(IGNORED_FOLDER)/coverage.out


.PHONY: all
all: tools install lint swag build ## Start all: tools install lint swag build


##
## Local stack development.
## Do not run following commands in CI
##

.PHONY: up
up: ## Up docker containers
	@docker-compose up --remove-orphans --build -d
	@docker-compose logs -f ${APP}

.PHONY: down
down: ## Down docker containers
	@docker-compose down

##
## Building
##

.PHONY: install
install: ## Download and install go mod
	go mod download


.PHONY: build
build: ## Build App
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ${BIN_FOLDER}/app $(shell go list -m)/cmd/${APP}


##
## Quality Code
##

.PHONY: lint
lint: ## Lint
	golangci-lint run

.PHONY: test
test: ## Test
	@mkdir -p ${IGNORED_FOLDER}
	@go test -tags=integration -count=1 -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...


##
## Tools
##

.PHONY: tools-lint
tools-lint: ## Install go lint dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: tools-swag
tools-swag: ## Install go docs dependencies
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: tools
tools: tools-lint tools-swag ## Install lint dependencies

##
## Docs
##

.PHONY: swag
swag: ## Generate swagger files
	@swag init --parseDependency --parseDepth=3 -g ./internal/app/app.go -o ./internal/swagger



##
## Files management
##

.PHONY: env
env: ## Create .env file from .env.example
	cp env.example .env

.PHONY: git-set-hooks
git-set-hooks: ## Setup git hooks
	cp -v scripts/git-hooks/* .git/hooks/

.PHONY: clean
clean: ## Clean
	@rm -rf ${BIN_FOLDER}
	@rm -rf ${IGNORED_FOLDER}
	@rm -rf ${COVERAGE_FILE}


##
## Help
##

.PHONY: help
help: ## Help
	@grep -E '(^[a-zA-Z_-]+:.*?#.*$$)|(^#)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m#/[33m/'
