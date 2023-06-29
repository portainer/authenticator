PLATFORM="$(shell go env GOOS)"
ARCH="$(shell go env GOARCH)"
TAG="latest"

.PHONY: pre build release image clean

dist := dist
bin := $(shell basename $(CURDIR))
image := portainer/authenticator:$(TAG)

pre:
	mkdir -pv $(dist) 

build: pre
	GOOS=$(PLATFORM) GOARCH=$(ARCH) CGO_ENABLED=0 go build --installsuffix cgo --ldflags '-s' -o $(bin) cmd/authenticator/main.go
	mv $(bin) $(dist)/

release: pre
	GOOS=$(PLATFORM) GOARCH=$(ARCH) CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags '-s' -o $(bin) cmd/authenticator/main.go
	mv $(bin) $(dist)/

image: release
	docker buildx build --push --platform=$(PLATFORM)/$(ARCH) -t $(image) .

clean:
	rm -rf $(dist)/*