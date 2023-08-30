VERSION := 0.1.6

clean:
	rm -rf out

distclean:
	rm -rf dist

fmt:
	go fmt ./...

test:
	go test ./...

build:
	mkdir -p out
	go build -o out/play-cdc
	chmod +x out/play-cdc

build-darwin-amd64:
	mkdir -p out/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o out/darwin-amd64/play-cdc
	chmod +x out/darwin-amd64/play-cdc

build-darwin-arm64:
	mkdir -p out/darwin-arm64
	GOOS=darwin GOARCH=arm64 go build -o out/darwin-arm64/play-cdc
	chmod +x out/darwin-arm64/play-cdc

build-linux-amd64:
	mkdir -p out/linux-amd64
	GOOS=linux GOARCH=amd64 go build -o out/linux-amd64/play-cdc
	chmod +x out/linux-amd64/play-cdc

build-linux-arm64:
	mkdir -p out/linux-arm64
	GOOS=linux GOARCH=arm64 go build -o out/linux-arm64/play-cdc
	chmod +x out/linux-arm64/play-cdc

install: build
	mkdir -p ~/.gauge/plugins/play-cdc/$(VERSION)
	mv out/play-cdc ~/.gauge/plugins/play-cdc/$(VERSION)
	cp plugin.json ~/.gauge/plugins/play-cdc/$(VERSION)

package: distclean build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64
	mkdir -p dist/linux-amd64
	cp out/linux-amd64/play-cdc dist/linux-amd64
	cp plugin.json dist/linux-amd64
	cd dist/linux-amd64 && zip -r play-cdc-$(VERSION)-linux.x86_64.zip .

	mkdir -p dist/linux-arm64
	cp out/linux-arm64/play-cdc dist/linux-arm64
	cp plugin.json dist/linux-arm64
	cd dist/linux-arm64 && zip -r play-cdc-$(VERSION)-linux.arm64.zip .

	mkdir -p dist/darwin-amd64
	cp out/darwin-amd64/play-cdc dist/darwin-amd64
	cp plugin.json dist/darwin-amd64
	cd dist/darwin-amd64 && zip -r play-cdc-$(VERSION)-darwin.x86_64.zip .

	mkdir -p dist/darwin-arm64
	cp out/darwin-arm64/play-cdc dist/darwin-arm64
	cp plugin.json dist/darwin-arm64
	cd dist/darwin-arm64 && zip -r play-cdc-$(VERSION)-darwin.arm64.zip .
