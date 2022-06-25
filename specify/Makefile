build:
	go build -o out/specify

run:
	go run main.go

install: build
	chmod +x out/specify
	mkdir -p ~/.gauge/plugins/specify/0.1.0
	mv out/specify ~/.gauge/plugins/specify/0.1.0
	cp plugin.json ~/.gauge/plugins/specify/0.1.0

clean:
	rm -rf out

fmt:
	go fmt ./...

test:
	go test ./...
