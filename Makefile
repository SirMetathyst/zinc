bench:
	go test -bench . -benchtime 15s -cpu 1,4 -count 2
	
build: 
	go build -o ./atom ./cmd/atom/main.go

install: 
	go install ./cmd/atom