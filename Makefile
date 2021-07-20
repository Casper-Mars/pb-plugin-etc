.PHONY: init
init:
	mkdir out

.PHONY: run
run:
	go install . && protoc --go_out=./out -I=. -I=./third_party --myplugin_out=./out --go-grpc_out=./out api/*.proto
