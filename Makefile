build: 
	go build -o ./zinc ./cmd/zinc/main.go

clean:
	rm -f ./coverage.html
	rm -f ./coverage.txt

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

testreport: test
	go tool cover -html=coverage.txt -o ./coverage.html

install: 
	go install ./cmd/zinc