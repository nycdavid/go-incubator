csv:
	go run main.csv.go
echo-session:
	go run main.echo-session.go
writer:
	go run main.writer.go
limitreader:
	go test ./limitreader

.PHONY: limitreader
