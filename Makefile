all: *.go
	go build .

run: garfish
	./garfish
