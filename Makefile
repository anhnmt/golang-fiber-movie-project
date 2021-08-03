APP_NAME=golang-fiber-base-project
APP_VERSION=1.0.0
BUILD_DIR=./build
DOCKER_LOCAL=ghcr.io/xdorro

config:
	cp config.example.yml config.yml

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	swag init

build: swag clean
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

sonar:
	sonar-scanner.bat -Dproject.settings=./sonar-project.properties

docker.build:
	docker build -t $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) .

docker.push:
	docker push $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION)

docker.run:
	docker run --name $(APP_NAME) -d -p 8000:8000 $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) -e SERVER_PORT=8000

docker.deploy: docker.build docker.run

docker.local: docker.build docker.push