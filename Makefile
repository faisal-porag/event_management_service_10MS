run:
	go run main.go

build:
	GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/event_management_service main.go

clean:
	rm -r bin

run_bin:
	./bin/event_management_service

docker_build:
	docker build -t event_management_service_10ms .

docker_run:
	#docker run -p 8090:8090 -t --name event_management_service_10MS_container event_management_service_10ms
	docker run -p 8090:8090 -d -t --name event_management_service_10MS_container event_management_service_10ms
