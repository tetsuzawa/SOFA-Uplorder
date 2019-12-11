.PHONY: build
build:
	go build -o build/$(BIN_NAME) .

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race

.PHONY: proto
proto:
	protoc -I ./ proto/sofa/sofa.proto --go_out=plugins=grpc:./
