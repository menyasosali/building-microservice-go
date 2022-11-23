.PHONY: protos

protos:
	protoc -I product-api/protos/ --go_out=./product-api/protos/currency --go-grpc_out=. product-api/protos/currency.proto

swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger && swagger generate spec -o ./swagger.yaml --scan-models

