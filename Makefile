.PHONY: tests
tests:
	go test -v -cover ./internal/loan_calc/controller/...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: docker_build
docker_build: lint tests
	docker build -t loan-calc .

.PHONY: docker_up
docker_up: docker_build
	docker run -d -p 8080:8080 --name loan-calc loan-calc

.PHONY: docker_down
docker_down:
	docker stop loan-calc
	docker rm loan-calc
