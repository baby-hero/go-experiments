GOTOOLS = golang.org/x/lint/golint \
	github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/goimports \
	mvdan.cc/gofumpt \
	github.com/bufbuild/buf/cmd/buf\
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/mwitkow/go-proto-validators/protoc-gen-govalidators


install-go-tools:
	go get $(GOTOOLS)

generate-proto:
	buf generate

create-docker-compose:
	docker-compose up -d  # Use -d to run in background, up to create and start containers
run-docker-compose:
	docker-compose start
stop-docker-compose:
	docker-compose stop
remove-docker-compose:
	docker-compose down  # Stop and remove all containers
