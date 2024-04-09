SHELL := /bin/bash
# $(shell git rev-parse --short HEAD)
VERSION := 1.0
ver ?= v1
run: tidy
	export WORKER_NS=ns-personal-shibmish; go run cmd/k8s-object-churner/main.go

build: tidy
	go build -o bin/koc cmd/k8s-object-churner/main.go

docker:
	docker build -t koc .

docker-run:
	 docker run -p 8080:8080 --rm -e WORKER_NS=ns-personal-shibmish -e  KUBECONFIG -e AWS_ACCESS_KEY_ID -e AWS_SESSION_TOKEN -e AWS_SECURITY_TOKEN -e AWS_SECRET_ACCESS_KEY --name my-running-app koc

docker-release:
	docker build -t koc-release -f Dockerfile-release .
	docker tag koc-release registry.gitlab.com/shib1000/koc:$(ver)

docker-push:
	docker push registry.gitlab.com/shib1000/koc:$(ver)

down:
	KILL -INT $(shell ps | grep go-build | grep -v grep | awk '{ print $$1 }')

fmt:
	gofmt -s -w .

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1

# ==============================================================================
# Install on k8s cluster
k8s-install:
	kubectl apply -f deployments/k8s.yaml

k8s-uninstall:
	kubectl delete -f deployments/k8s.yaml