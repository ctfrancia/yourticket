BINARY_NAME=yourticket
 
build:
	go build -o ./bin/${BINARY_NAME} ./cmd/api

deploy-dev:
	docker build -t yourticket .
	docker run -d -p 8080:8080 yourticket
 
run:
	go build -o ./bin/${BINARY_NAME} ./cmd/api
	./bin/${BINARY_NAME}
 
clean:
	go clean
	rm ./bin/${BINARY_NAME}

