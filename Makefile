setup:
	echo "downloading zipkin docker-compose file"
	./starter.sh
	echo "Done! Remember to modify docker-compose to include the testserver"

start: build up

build:
	echo "building testserver"
	@env GOOS=linux GOARCH=amd64 go build -o testserver cmd/testServer/main.go
	@chmod +x testserver
	docker-compose build

up:
	echo "starting up testserver, zipkin, and mysql"
	docker-compose up

clean:
	echo "cleaning up testserver binary"
	rm testserver
	echo "cleaning up docker-compose"
	rm docker-compose.yml
	echo "All done!"

.PHONY: setup start build up clean
