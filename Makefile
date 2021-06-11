
APP_NAME = golang-fiber-base-project
BUILD_DIR = ./build

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

docker.mariadb:
	docker run --rm -d \
		--name demo_mariadb \
		-e MYSQL_ROOT_PASSWORD=123456aA@ \
		-v data_mariadb:/var/lib/mysql \
		-p 3306:3306 \
		-d mariadb:10 \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci