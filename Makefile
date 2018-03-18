PROJECT?=tutor
APP?=tutor
PORT?=8000

GOOS?=linux
GOARCH?=amd64

RELEASE?=0.2.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOENV?=docker run -ti -e PORT -v ~/kitchen/k8s-tutorial/src:/go/src --workdir /go/src/tutor golang:1.9
# add it to Docker -> Preferences -> Insecure registries
REGISTRY?=macos.local:5000

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	$(GOENV) go build \
	  -ldflags "-s -w \
		-X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(REGISTRY)/$(APP):$(RELEASE) .

push: container
	docker push $(REGISTRY)/$(APP):$(RELEASE)

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
	    -e "PORT=${PORT}" \
	    $(REGISTRY)/$(APP):$(RELEASE)

minikube: push
	for t in $(shell find ./kubernetes -type f -name "*.yaml"); do \
	    cat $$t | \
	        sed "s/{{\s*\.Release\s*}}/$(RELEASE)/g" | \
	        sed "s/{{\s*\.ServiceName}}/$(APP)/g"; \
	    echo ---; \
	done > k8s.yaml
	kubectl apply -f k8s.yaml

test:
	$(GOENV) test -v ./...
