all: *.go
	go build .

run: garfish
	./garfish

test: *.go
	go test -v
