NAME := stats
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'
SRCS := $(shell find . -type f -name '*.go')

bin/$(NAME): $(SRCS)
	go build -ldflags "$(LDFLAGS)" -o $@

test:
	go test ./...

tools:
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		golang.org/x/tools/cmd/goimports \
		github.com/alecthomas/gometalinter \
		github.com/Songmu/goxz/cmd/goxz \
		github.com/tcnksm/ghr

lint: tools
	gometalinter --disable-all --enable=vet --enable=golint --enable=goimports ./...

fmt: tools
	goimports -w -d -local $$(go list) $(SRCS)

install:
	go install -ldflags "$(LDFLAGS)"

clean:
	rm -rf bin/

cross-build: tools
	goxz -pv $(VERSION) -build-ldflags "$(LDFLAGS)" -arch 386,amd64 -d dist/$(VERSION)
	cd dist/$(VERSION) && shasum -a 256 * > ./$(VERSION)_SHASUMS

upload: tools
	ghr $(VERSION) dist/$(VERSION)

.PHONY: test tools lint fmt insatll clean cross-build upload
