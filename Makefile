appname := aeroflot-api
sources := vendor $(wildcard *.go)
GIT_SHA_FETCH := $(shell git rev-parse HEAD | cut -c 1-8)

build = GO111MODULE=on GOOS=$(1) GOARCH=$(2) go build -ldflags "-X main.Commit=$(GIT_SHA_FETCH)" -mod=vendor -o build/$(appname)-$(1)-$(2)$(3)
static-build = GO111MODULE=on CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) GOARM=$(4) go build -ldflags "-s -w -X main.Commit=$(GIT_SHA_FETCH)" -mod=vendor -a -installsuffix cgo -o build/$(appname)-$(1)-$(2)$(4)-static$(3) .
debug-build = GO111MODULE=on CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) GOARM=$(4) go build -ldflags "-s -w -X main.DEBUG=true -X main.Commit=$(GIT_SHA_FETCH)" -mod=vendor -a -installsuffix cgo -o build/$(appname)-$(1)-$(2)$(4)-debug$(3) .
tar = cd build && tar -cvzf $(appname)-$(1)-$(2).tar.gz $(appname)-$(1)-$(2)$(3) && rm $(appname)-$(1)-$(2)$(3)
zip = cd build && zip $(appname)-$(1)-$(2).zip $(appname)-$(1)-$(2)$(3) && rm $(appname)-$(1)-$(2)$(3)

.PHONY: all windows darwin linux clean windows-static darwin-static linux-static windows-compressed linux-compressed darwin-compressed test

all: linux linux-static-amd64 linux-static-arm6 test

all-static: linux-static test

test:
	go test -v

release-files: windows-compressed linux-compressed darwin-compressed

clean:
	rm -rf build/

fmt:
	@gofmt -l -w $(sources)


##### LINUX #####

linux: build/$(appname)-linux-amd64

linux-compressed: build/$(appname)-linux-amd64.tar.gz

linux-static-amd64: build/$(appname)-linux-amd64-static

linux-debug-amd64: build/$(appname)-linux-amd64-debug

build/$(appname)-linux-amd64: $(sources)
	$(call build,linux,amd64,)

build/$(appname)-linux-amd64.tar.gz: build/$(appname)-linux-amd64
	$(call tar,linux,amd64)

build/$(appname)-linux-amd64-static: $(sources)
	$(call static-build,linux,amd64,)

build/$(appname)-linux-amd64-debug: $(sources)
	$(call debug-build,linux,amd64,)


##### RaspberryPi #####

linux-static-arm5: build/$(appname)-linux-arm5-static

linux-static-arm6: build/$(appname)-linux-arm6-static

linux-debug-arm6: build/$(appname)-linux-arm6-debug

linux-static-arm7: build/$(appname)-linux-arm7-static

build/$(appname)-linux-arm5-static: $(sources)
	$(call static-build,linux,arm,,5)

build/$(appname)-linux-arm6-static: $(sources)
	$(call static-build,linux,arm,,6)

build/$(appname)-linux-arm6-debug: $(sources)
	$(call debug-build,linux,arm,,6)

build/$(appname)-linux-arm7-static: $(sources)
	$(call static-build,linux,arm,,7)


##### DARWIN (MAC) #####

darwin: build/$(appname)-darwin-amd64

darwin-compressed: build/$(appname)-darwin-amd64.tar.gz

darwin-static: build/$(appname)-darwin-amd64-static

build/$(appname)-darwin-amd64: $(sources)
	$(call build,darwin,amd64,)

build/$(appname)-darwin-amd64.tar.gz: build/$(appname)-darwin-amd64
	$(call tar,darwin,amd64)


build/$(appname)-darwin-amd64-static: $(sources)
	$(call static-build,darwin,amd64,)


##### WINDOWS #####

windows: build/$(appname)-windows-amd64.exe

windows-compressed: build/$(appname)-windows-amd64.zip

windows-static: build/$(appname)-windows-amd64-static

build/$(appname)-windows-amd64.exe: $(sources)
	$(call build,windows,amd64,.exe)

build/$(appname)-windows-amd64.zip: build/$(appname)-windows-amd64
	$(call zip,windows,amd64,.exe)

build/$(appname)-windows-amd64-static: $(sources)
	$(call static-build,windows,amd64,)

build/$(appname)-windows-amd64: $(sources)
	$(call build,windows,amd64,.exe)


##### Vendoring #####

go.mod:
	go mod init

vendor-update: go.mod
	GO111MODULE=on go get -u=patch

vendor: go.mod
	GO111MODULE=on go mod vendor