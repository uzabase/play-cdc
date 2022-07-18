build:
	mkdir -p out
	go build -o out/play-cdc

install: build
	chmod +x out/play-cdc
	mkdir -p ~/.gauge/plugins/play-cdc/0.1.0
	mv out/play-cdc ~/.gauge/plugins/play-cdc/0.1.0
	cp plugin.json ~/.gauge/plugins/play-cdc/0.1.0

clean:
	rm -rf out

fmt:
	go fmt ./...

test:
	go test ./...
