default: out/example

clean:
	rm -rf out

test: *.go
	go test ./...

out/example: implementation.go main.go
	mkdir -p out
	go build -o out/example
