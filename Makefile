csv:
	go run main.csv.go
echo-session:
	go run main.echo-session.go
writer:
	go run main.writer.go
limitreader:
	go test -v ./limitreader

run:
	docker run \
	-it \
	-p 8000:8000 \
	-v $(shell pwd)/:/go/src/github.com/nycdavid/go-incubator \
	--rm \
	golang:1.10-alpine \
	go run src/github.com/nycdavid/go-incubator/${PACKAGE}/main.go ${ARG}

test:
	docker run \
	-it \
	-v $(shell pwd)/${PACKAGE}/:/go/src/app \
	--rm \
	golang:1.10-alpine \
	go test -v ./src/app/...

.PHONY: limitreader
