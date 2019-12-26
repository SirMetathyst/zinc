bench:
	go test -bench . -benchtime 15s -cpu 1,4 -count 2
	
build: 
	go build -o ./atom ./cmd/atom/main.go

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

testreport: test
	go tool cover -html=coverage.txt -o ./coverage.html

install: 
	go install ./cmd/atom