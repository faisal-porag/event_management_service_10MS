run:
	go run main.go

build:
	GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/event_management_service main.go

clean:
	rm -r bin

run_bin:
	./bin/event_management_service


